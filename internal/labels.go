// Copyright 2020 Red Hat, Inc. and/or its affiliates
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

package internal

import (
	"fmt"
	"strings"

	"github.com/kiegroup/rhpam-kogito-operator/version"
)

var meteringLabels = map[string]string{
	"com.company":   "Red_Hat",
	"rht.prod_name": "Red_Hat_Process_Automation",
	"rht.prod_ver":  trimMinorVersion(version.Version),
	"rht.comp":      "PAM",
	"rht.comp_ver":  version.PAMVersion,
	"rht.subcomp":   "rhpam-kogito-runtime",
	"rht.subcomp_t": "application",
}

// SetMeteringLabels for product operator
func SetMeteringLabels(deploymentLabels map[string]string) map[string]string {
	if deploymentLabels == nil {
		deploymentLabels = map[string]string{}
	}
	for labelKey, labelValue := range meteringLabels {
		deploymentLabels[labelKey] = labelValue
	}
	return deploymentLabels
}

func trimMinorVersion(version string) string {
	splitVersion := strings.Split(version, ".")
	return fmt.Sprintf("%s.%s", splitVersion[0], splitVersion[1])
}
