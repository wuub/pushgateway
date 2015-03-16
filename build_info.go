// Copyright 2014 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"text/template"
)

// Build information. Populated by Makefile.
var (
	buildVersion string
	buildRev     string
	buildBranch  string
	buildUser    string
	buildDate    string
)

// BuildInfo encapsulates compile-time metadata about Prometheus made available
// via go tool ld such that this can be reported on-demand.
var BuildInfo = map[string]string{
	"version": buildVersion,
	"commit":  buildRev,
	"branch":  buildBranch,
	"user":    buildUser,
	"date":    buildDate,
}

var versionInfoTmpl = template.Must(template.New("version").Parse(
	`pushgateway, version {{.version}} ({{.branch}}, commit {{.commit}})
  build user:       {{.user}}
  build date:       {{.date}}
`))
