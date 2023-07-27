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
	"fmt"
	"strconv"
	"strings"

	_ "embed"

	"golang.org/x/mod/semver"
)

// SemVer is a semantic version. See https://go.dev/doc/modules/version-numbers
// for details.
type SemVer struct {
	Major int
	Minor int
	Patch int
}

var (
	//go:embed files/module.version
	moduleVersion string

	//go:embed files/deployer.version
	deployerVersion string

	//go:embed files/codegen.version
	codegenVersion string
)

var (
	// The weaver module version.
	ModuleVersion = NewSemVer(moduleVersion)

	// The deployer API version.
	DeployerVersion = NewSemVer(deployerVersion)

	// The codegen API version.
	CodegenVersion = NewSemVer(codegenVersion)
)

func NewSemVer(v string) (version SemVer) {
	before, after, _ := strings.Cut(v, ".")
	major, _ := strconv.Atoi(before)
	after, last, _ := strings.Cut(after, ".")
	minor, _ := strconv.Atoi(after)
	patch, _ := strconv.Atoi(last)
	return SemVer{
		Major: major,
		Minor: minor,
		Patch: patch,
	}
}

func (s SemVer) GetLastVersionModule() (v string) {
	return
}

func (s SemVer) GetLastVersionDeployer() (v string) {
	return
}

func (s SemVer) GetLastVersionCodegen() (v string) {
	return
}

func (s SemVer) WeNeedUpdatedVersion(lastVersion string) bool {
	currentVersion := s.String()
	return semver.Max(currentVersion, lastVersion) != currentVersion
}

func (s SemVer) String() string {
	return fmt.Sprintf("v%d.%d.%d", s.Major, s.Minor, s.Patch)
}
