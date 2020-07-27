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

package util

import (
	"strings"

	"k8s.io/apimachinery/pkg/labels"

	"istio.io/istio/pkg/config/resource"
)

const (
	// Path templates for different fields with different paths, may edited by future developers if not covered in this list
	// Use the path template to find the exact line number for the field

	// Path for host in VirtualService.
	// Required parameters: route rule, route rule index, route index.
	DestinationHost = "{.spec.%s[%d].route[%d].destination.host}"

	// Path for mirror host in VirtualService.
	// Required parameters: http index.
	MirrorHost = "{.spec.http[%d].mirror.host}"

	// Path for VirtualService gateway.
	// Required parameters: gateway index.
	VSGateway = "{.spec.gateways[%d]}"

	// Path for regex match of uri, scheme, method and authority.
	// Required parameters: http index, match index, where to match.
	URISchemeMethodAuthorityRegexMatch = "{.spec.http[%d].match[%d].%s.regex}"

	// Path for regex match of headers and queryParams.
	// Required parameters: http index, match index, where to match, match key.
	HeaderAndQueryParamsRegexMatch = "{.spec.http[%d].match[%d].%s.%s.regex}"

	// Path for regex match of allowOrigins.
	// Required parameters: http index, allowOrigins index.
	AllowOriginsRegexMatch = "{.spec.http[%d].corsPolicy.allowOrigins[%d].regex}"

	// Path for workload selector.
	// Required parameters: selector label.
	WorkloakSelector = "{.spec.workloadSelector.labels.%s}"

	// Path for port from ports collections.
	// Required parameters: port index.
	PortInPorts = ".spec.ports[%d].port}"

	// Path for fromRegistry in the mesh networks.
	// Required parameters: network name, endPoint index.
	FromRegistry = "{.networks.%s.endpoints[%d]}"

	// Path for the image in the container.
	// Required parameters: container index.
	ImageInContainer = "{.spec.containers[%d].image}"

	// Path for namespace in metadata.
	// Required parameters: none.
	MetadataNamespace = "{.metadata.namespace}"

	// Path for name in metadata.
	// Required parameters: none.
	MetadataName = "{.metadata.name}"

	// Path for namespace in authorizationPolicy.
	// Required parameters: rule index, from index, namespace index.
	AuthorizationPolicyNameSpace = "{.spec.rules[%d].from[%d].source.namespaces[%d]}"

	// Path for annotation.
	// Required parameters: annotation name.
	Annotation = "{.metadata.annotations.%s}"

	// Path for selector in Gateway.
	// Required parameters: selector label
	GatewaySelector = "{.spec.selector.%s}"

	// Path for credentialName.
	// Required parameters: server index
	CredentialName = "{.spec.servers[%d].tls.credentialName}"
)

func ErrorLine(r *resource.Instance, path string) (line int, found bool) {
	return FindErrorLine(path, r.Origin.FieldMap())
}

// FindErrorLine returns the line number of the input key from the input map, and true if retrieving successfully,
// else return 0 and false
func FindErrorLine(key string, m map[string]int) (int, bool) {
	line, ok := m[key]
	if !ok {
		return 0, false
	}
	return line, true
}

// FindLabelForSelector returns the label for the k8s labels.Selector
func FindLabelForSelector(selector labels.Selector) string {
	s := selector.String()
	equalIndex := strings.Index(s, "=")
	return s[:equalIndex]
}
