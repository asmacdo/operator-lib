// Copyright 2021 The Operator-SDK Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package proxy

import (
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
)

var ProxyEnvNames = []string{"HTTP_PROXY", "HTTPS_PROXY", "NO_PROXY"}

func ReadProxyVarsFromEnv() []corev1.EnvVar {
	envVars := []corev1.EnvVar{}
	for _, s := range ProxyEnvNames {
		value := os.Getenv(s)
		envVars = append(envVars, corev1.EnvVar{
			Name:  s,
			Value: value,
		})
		// Duplicate values for upper and lower case
		envVars = append(envVars, corev1.EnvVar{
			Name:  strings.ToLower(s),
			Value: value,
		})
	}
	return envVars
}
