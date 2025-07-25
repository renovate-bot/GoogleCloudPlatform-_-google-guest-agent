//  Copyright 2024 Google LLC
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//go:build linux

package ps

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"

	"github.com/GoogleCloudPlatform/galog"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/run"
	"github.com/GoogleCloudPlatform/google-guest-agent/internal/utils/file"
)

// linuxClient is for finding processes on linux distributions.
type linuxClient struct {
	commonClient
	// procDir is the location of proc filesystem mount point in a linux system.
	procDir string
}

const (
	// defaultLinuxProcDir is the default location of proc filesystem mount point
	// in a linux system.
	defaultLinuxProcDir = "/proc/"
	// defaultClkTime is the default value of CLK_TKC. 100 is a typical value
	// and is used in case `getconf CLK_TCK` fails.
	defaultClkTime = 100
)

var (
	// sawClkTckErr is a flag to indicate if getconf CLK_TCK failed.
	// If it fails, we use this flag to avoid repeated logging.
	sawClkTckErr error
)

// init creates the Linux process finder.
func init() {
	Client = &linuxClient{
		procDir: defaultLinuxProcDir,
	}
}

// readClkTicks reads the [CLK_TCK] value by running [getconf CLK_TCK] command.
func readClkTicks(ctx context.Context) (float64, error) {
	if _, err := exec.LookPath("getconf"); err != nil {
		return defaultClkTime, nil
	}
	opts := run.Options{
		Name:       "getconf",
		Args:       []string{"CLK_TCK"},
		OutputType: run.OutputStdout,
	}
	clkRes, err := run.WithContext(ctx, opts)
	if err != nil {
		// Avoid repeated logging if we already saw the error.
		if !errors.Is(err, sawClkTckErr) {
			galog.V(2).Warnf("getconf CLK_TCK failed: %v", err)
			sawClkTckErr = err
		}
		return defaultClkTime, nil
	}
	return strconv.ParseFloat(strings.TrimSpace(clkRes.Output), 64)
}

func (p linuxClient) FindPid(pid int) (Process, error) {
	return p.readPidDetails(pid)
}

// FindRegex finds all processes with the executable path matching the provided
// regex.
func (p linuxClient) FindRegex(exeMatch string) ([]Process, error) {
	var result []Process

	procExpression, err := regexp.Compile("^[0-9]+$")
	if err != nil {
		return nil, fmt.Errorf("failed to compile process dir expression: %w", err)
	}

	exeExpression, err := regexp.Compile(exeMatch)
	if err != nil {
		return nil, fmt.Errorf("failed to compile process exec matching expression: %w", err)
	}

	files, err := os.ReadDir(p.procDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read linux proc dir: %w", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		if !procExpression.MatchString(file.Name()) {
			continue
		}

		// Ignore the error due to the `procExpression` regex, which ensures that
		// the file name is a valid PID, and therefore we should not expect any
		// errors.
		pid, _ := strconv.Atoi(file.Name())

		process, err := p.readPidDetails(pid)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			galog.V(2).Debugf("Failed to read process(%d), while finding processes matching regex %q: [%v], skipping...", pid, exeMatch, err)
			continue
		}

		if !exeExpression.MatchString(process.Exe) {
			continue
		}

		result = append(result, process)
	}

	return result, nil
}

func (p linuxClient) readPidDetails(pid int) (Process, error) {
	var process Process

	// Find the executable path.
	processRootDir := filepath.Join(p.procDir, fmt.Sprintf("%d", pid))

	if !file.Exists(processRootDir, file.TypeDir) {
		return process, fmt.Errorf("process %d does not exist, %q not found", pid, processRootDir)
	}

	exeLinkPath := filepath.Join(processRootDir, "exe")

	exePath, err := os.Readlink(exeLinkPath)
	if err != nil {
		return process, fmt.Errorf("error reading executable path: %w", err)
	}

	// Find the command line.
	cmdlinePath := filepath.Join(processRootDir, "cmdline")
	dat, err := os.ReadFile(cmdlinePath)
	if err != nil {
		return process, fmt.Errorf("error reading cmdline file: %w", err)
	}

	var commandLine []string
	var token []byte
	for _, curr := range dat {
		if curr == 0 {
			commandLine = append(commandLine, string(token))
			token = nil
		} else {
			token = append(token, curr)
		}
	}

	return Process{
		PID:         pid,
		Exe:         exePath,
		CommandLine: commandLine,
	}, nil
}

// Memory returns the memory usage in kB of the process with the provided PID.
func (p linuxClient) Memory(pid int) (int, error) {
	baseProcDir := filepath.Join(p.procDir, strconv.Itoa(pid))

	var stats []byte
	var readErrors error
	var readFile bool
	var openFile string

	// Read the smaps file. This is where the memory usage of the process is
	// stored.
	for _, fpath := range []string{"smaps", "smaps_rollup"} {
		galog.V(3).Debugf("Reading %s file for pid %d", fpath, pid)
		var err error
		openFile = filepath.Join(baseProcDir, fpath)
		stats, err = os.ReadFile(openFile) // NOLINT
		if err != nil {
			// If the error is not "not exist" means we failed to read it, in that
			// case we don't fallback to other files.
			if !os.IsNotExist(err) {
				return 0, fmt.Errorf("error reading %s file: %w", fpath, err)
			}

			// If the error is "not exist", we fallback to other files and keep track
			// of the errors.
			if readErrors == nil {
				readErrors = err
			} else {
				readErrors = fmt.Errorf("%w; %w", readErrors, err)
			}
		}

		readFile = true
		if err == nil {
			break
		}
	}

	if !readFile && readErrors != nil {
		return 0, fmt.Errorf("error reading smaps/smaps_rollup file: %w", readErrors)
	}

	statsLines := strings.Split(string(stats), "\n")
	foundRss := false
	var memUsage int

	// Now find the memory line. This line is the RSS line.
	for _, line := range statsLines {
		if strings.HasPrefix(line, "Rss") {
			galog.V(3).Debugf("Found RSS line: %s", line)
			foundRss = true
			partial, err := strconv.Atoi(strings.Fields(line)[1])
			if err != nil {
				return 0, fmt.Errorf("error parsing RSS line: %w", err)
			}
			memUsage += partial
		}
	}

	if !foundRss {
		return 0, fmt.Errorf("no RSS line found in %s file", openFile)
	}

	return memUsage, nil
}

// CPUUsage returns the % CPU usage of the process with the provided PID.
func (p linuxClient) CPUUsage(ctx context.Context, pid int) (float64, error) {
	baseProcDir := filepath.Join(p.procDir, strconv.Itoa(pid))

	// Read the stat file. This is where the usage data is kept.
	stats, err := os.ReadFile(filepath.Join(baseProcDir, "stat"))
	if err != nil {
		return 0, fmt.Errorf("error reading stat file: %w", err)
	}
	statsLines := strings.Fields(string(stats))

	// Read utime and stime values. The fields are not labelled, so get them
	// as-is.
	utime, err := strconv.ParseFloat(statsLines[13], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing utime: %w", err)
	}

	stime, err := strconv.ParseFloat(statsLines[14], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing stime: %w", err)
	}

	// Total time used is the sum of both values.
	// Since the values are in clock ticks, divide by clock tick.
	totalTimeTicks := utime + stime
	clkTime, err := readClkTicks(ctx)
	if err != nil {
		return 0, fmt.Errorf("error reading clk time: %w", err)
	}

	runTime := totalTimeTicks / clkTime

	// Get the process start time and system uptime. These make up for the total
	// elapsed time since the process started.
	startTimeTicks, err := strconv.ParseFloat(statsLines[21], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing start time: %w", err)
	}
	startTime := startTimeTicks / clkTime

	// uptime is the total running time of the system.
	uptimeContents, err := os.ReadFile(filepath.Join(p.procDir, "uptime"))
	if err != nil {
		return 0, fmt.Errorf("error reading /proc/uptime file: %w", err)
	}
	uptime, err := strconv.ParseFloat(strings.Fields(string(uptimeContents))[0], 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing /proc/uptime: %w", err)
	}

	// Time the process spent running.
	elapsedTime := uptime - startTime

	// Divide by elapsed time.
	return runTime / elapsedTime, nil
}

// IsProcessAlive returns true if the process with the provided PID is alive.
// It does not return an error in any case and caller can simply rely on the
// returned boolean. Error is part of the signature to be compatible with the
// other platforms.
func (p linuxClient) IsProcessAlive(pid int) (bool, error) {
	galog.Debugf("Checking if process %d is alive", pid)

	// Its safe to ignore the error, as per docs find process never returns error
	// on unix systems.
	proc, _ := os.FindProcess(pid)
	err := proc.Signal(syscall.Signal(0))
	if err != nil {
		galog.Debugf("Process with pid %d not running, signal 0 returned error: %v", pid, err)
		return false, nil
	}
	return true, nil
}
