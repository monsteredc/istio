// Copyright Istio Authors
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

package diag

import (
	"istio.io/istio/pkg/config/resource"
)

var _ resource.Origin = &testOrigin{}
var _ resource.Reference = &testReference{}

type testOrigin struct {
	name string
	ref  resource.Reference
}

func (o testOrigin) FriendlyName() string {
	return o.name
}

func (o testOrigin) Namespace() resource.Namespace {
	return ""
}

func (o testOrigin) Reference() resource.Reference {
	return o.ref
}

type testReference struct {
	name string
}

func (r testReference) String(bool) string {
	return r.name
}

func (r testReference) ProcessMap([]byte) {
}

func (r testReference) YamlMap() map[string]interface{} {
	yamlMap := make(map[string]interface{})
	return yamlMap
}

func (r testReference) FindErrors(map[string]map[string]string) ([]string, []int) {
	return []string{}, []int{}
}