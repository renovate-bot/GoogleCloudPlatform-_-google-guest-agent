//  Copyright 2023 Google LLC
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//go:build unix

package osinfo

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/GoogleCloudPlatform/galog"
	"golang.org/x/sys/unix"
)

var (
	// Known default OS release file location.
	osRelease = "/etc/os-release"
	// Known default system release file location.
	systemRelease = "/etc/system-release"
)

func parseOSRelease(osRelease string) (OSInfo, error) {
	var ret OSInfo
	for _, line := range strings.Split(osRelease, "\n") {
		id := line
		if id = strings.TrimPrefix(line, "ID="); id != line {
			id = strings.Trim(id, `"`)
			ret.OS = parseID(id)
		}
		if id = strings.TrimPrefix(line, "VERSION_ID="); id != line {
			id = strings.Trim(id, `"`)
			ret.VersionID = id
			version, err := parseVersion(id)
			if err != nil {
				return ret, fmt.Errorf("couldn't parse version id: %v", err)
			}
			ret.Version = version
		}
		if name := strings.TrimPrefix(line, "PRETTY_NAME="); name != line {
			name = strings.Trim(name, `"`)
			ret.PrettyName = name
		}
	}
	return ret, nil
}

func parseSystemRelease(systemRelease string) (OSInfo, error) {
	var ret OSInfo
	key := " release "
	idx := strings.Index(systemRelease, key)
	if idx == -1 {
		return ret, errors.New("SystemRelease does not match expected format")
	}
	ret.OS = parseID(systemRelease[:idx])

	versionFromRelease := strings.Split(systemRelease[idx+len(key):], " ")[0]
	version, err := parseVersion(versionFromRelease)
	if err != nil {
		return ret, fmt.Errorf("couldn't parse version: %w", err)
	}
	ret.Version = version
	return ret, nil
}

func parseVersion(version string) (Ver, error) {
	versionparts := strings.Split(version, ".")
	ret := Ver{Length: len(versionparts)}

	// Must have at least major version.
	var err error
	ret.Major, err = strconv.Atoi(versionparts[0])
	if err != nil {
		return ret, err
	}
	if ret.Length > 1 {
		ret.Minor, err = strconv.Atoi(versionparts[1])
		if err != nil {
			return ret, err
		}
	}
	if ret.Length > 2 {
		ret.Patch, err = strconv.Atoi(versionparts[2])
		if err != nil {
			return ret, err
		}
	}
	return ret, nil
}

func parseID(id string) string {
	switch id {
	case "Red Hat Enterprise Linux Server":
		return "rhel"
	case "CentOS", "CentOS Linux":
		return "centos"
	default:
		return id
	}
}

func parseRelease() (OSInfo, error) {
	if releaseFile, err := os.ReadFile(osRelease); err == nil {
		return parseOSRelease(string(releaseFile))
	}
	if releaseFile, err := os.ReadFile(systemRelease); err == nil {
		return parseSystemRelease(string(releaseFile))
	}
	return OSInfo{}, errors.New("no known release file found")
}

// Read returns OSInfo on the running system.
func Read() OSInfo {
	osInfo, err := parseRelease()
	if err != nil {
		// This is a non critical error, we can still return a partially populated OSInfo.
		galog.Warnf("Error parsing release info: %v", err)
	}

	var uts unix.Utsname
	if err := unix.Uname(&uts); err != nil {
		galog.Warnf("unix.Uname error: %v", err)
		return osInfo
	}

	osInfo.KernelVersion = unix.ByteSliceToString(uts.Version[:])
	osInfo.KernelRelease = unix.ByteSliceToString(uts.Release[:])
	osInfo.Architecture = unix.ByteSliceToString(uts.Machine[:])

	return osInfo
}
