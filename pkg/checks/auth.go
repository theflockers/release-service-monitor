/*
Copyright 2024 Red Hat Inc

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
package checks

type QuayAuth struct {
	username string
	password string
}

// NewQuayAuth returns a new instance of QuayAuth
func NewQuayAuth(username string, password string) *QuayAuth {
	auth := &QuayAuth{
		username: username,
		password: password,
	}

	return auth
}

// getUsername return a QuayAuth username
func (a *QuayAuth) getUsername() string {
	return a.username
}

// getPassword return a QuayAuth password
func (a *QuayAuth) getPassword() string {
	return a.password
}
