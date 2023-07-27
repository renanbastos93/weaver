// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package version contains the version of the weaver module and its
// constituent APIs (e.g., the pipe API, the codegen API).
package version

import (
	_ "embed"
	"testing"
)

func TestNewSemVer(t *testing.T) {
	// assert.EqualValues(t, "v0.18.0", NewSemVer("0.18.0").String())
	// assert.EqualValues(t, "v0.0.0", NewSemVer("v0.r.z").String())
}

func TestXxx(t *testing.T) {
	// s := NewSemVer("0.11.0")
	// lastModuleVersion := s.GetLastVersionModule()
	// assert.EqualValues(t, s.String(), NewSemVer(lastModuleVersion).String())
	// s.WeNeedUpdatedVersion(lastModuleVersion)
}
