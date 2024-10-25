/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package test

import (
	"fmt"

	providerconfig "k8c.io/machine-controller/pkg/providerconfig/types"
)

// mock that raises an error when try to read secret.
func ShouldNotBeCalled(configVar *providerconfig.GlobalSecretKeySelector, key string) (string, error) {
	return "", fmt.Errorf("'GetGlobalSecretKeySelectorValue' should not be called")
}

// mock that returns default value when reading secret or value specify in generator map.
// Default value is key + "-value"
// generator is map of key (of GlobalSecretKeySelector) and value to return for this key. Value can be an error or a string.
func DefaultOrOverride(generator map[string]interface{}) func(configVar *providerconfig.GlobalSecretKeySelector, key string) (string, error) {
	return func(configVar *providerconfig.GlobalSecretKeySelector, key string) (string, error) {
		if val, ok := generator[key]; ok {
			if err, ok := val.(error); ok {
				return "", err
			}
			return val.(string), nil
		}
		return key + "-value", nil
	}
}

// return an error with message: secret "default/the-secret" has no key "<key>".
func MissingKeyErr(key string) error {
	return fmt.Errorf("secret \"default/the-secret\" has no key \"%v\"", key)
}
