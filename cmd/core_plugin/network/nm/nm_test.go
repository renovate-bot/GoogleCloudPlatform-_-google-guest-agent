// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nm

import (
	"testing"
)

func TestDefaultModule(t *testing.T) {
	mod := defaultModule()

	if mod.networkScriptsDir != defaultNetworkScriptsDir {
		t.Errorf("defaultModule() returned diff networkScriptsDir: got %q, want %q.", mod.networkScriptsDir, defaultNetworkScriptsDir)
	}

	if mod.configDir != defaultNetworkManagerConfigDir {
		t.Errorf("defaultModule() returned diff configDir: got %q, want %q.", mod.configDir, defaultNetworkManagerConfigDir)
	}
}
