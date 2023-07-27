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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	_ "embed"
)

// SemVer is a semantic version. See https://go.dev/doc/modules/version-numbers
// for details.
type SemVer struct {
	Major int64 `json:"major,omitempty"`
	Minor int64 `json:"minor,omitempty"`
	Patch int64 `json:"patch,omitempty"`
}

type Versions struct {
	Module   SemVer `json:"module,omitempty"`
	Deployer SemVer `json:"deployer,omitempty"`
	Codegen  SemVer `json:"codegen,omitempty"`
}

var (
	//go:embed files/versions.json
	VersionsJson []byte
)

var (
	v Versions

	ModuleVersion = v.Module
	ModuleMajor   = v.Module.Major
	ModuleMinor   = v.Module.Minor
	ModulePatch   = v.Module.Patch

	DeployerVersion = v.Deployer
	DeployerMajor   = v.Deployer.Major
	DeployerMinor   = v.Deployer.Minor

	CodegenVersion = v.Codegen
	CodegenMajor   = v.Codegen.Minor
	CodegenMinor   = v.Codegen.Minor

	baseURL = "https://raw.githubusercontent.com/renanbastos93/weaver/chore/validate-versions/runtime/version/files/"
)

func init() {
	_ = json.Unmarshal(VersionsJson, &v)
	ModuleVersion = v.Module
	ModuleMajor = v.Module.Major
	ModuleMinor = v.Module.Minor
	ModulePatch = v.Module.Patch
	DeployerVersion = v.Deployer
	DeployerMajor = v.Deployer.Major
	DeployerMinor = v.Deployer.Minor
	CodegenVersion = v.Codegen
	CodegenMajor = v.Codegen.Minor
	CodegenMinor = v.Codegen.Minor
}

func (s SemVer) GetLastVersion() (v string) {
	return s.getVersionFileOnGit("versions.json")
}

func (s SemVer) String() string {
	return fmt.Sprintf("v%d.%d.%d", s.Major, s.Minor, s.Patch)
}

func (s SemVer) getVersionFileOnGit(filename string) (v string) {
	var (
		client   *http.Client
		response *http.Response
		body     []byte
		err      error
	)

	client = &http.Client{
		Timeout: time.Second,
	}

	response, err = client.Get(baseURL + filename)
	if err != nil {
		return
	}

	if response.StatusCode != http.StatusOK {
		return
	}

	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	return string(body)
}
