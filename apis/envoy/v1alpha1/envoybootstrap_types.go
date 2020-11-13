/*


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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EnvoyBootstrapSpec defines the desired state of EnvoyBootstrap
type EnvoyBootstrapSpec struct {
	// DiscoveryService is the name of the DiscoveryService resource the envoy will be a client of
	DiscoveryService string `json:"discoveryService"`
	// ClientCertificate is a struct containing options for the certificate used to authenticate with the
	// discovery service
	// +optional
	ClientCertificate *ClientCertificate `json:"clientCertificate,omitempty"`
	// EnvoyStaticConfig is a struct that controls options for the envoy's static config file
	// +optional
	EnvoyStaticConfig *EnvoyStaticConfig `json:"envoyStaticConfig,omitempty"`
}

// EnvoyStaticConfig allows specifying envoy static config
// options
type EnvoyStaticConfig struct {
	// The ConfigMap where the envoy client static config will be stored
	ConfigMapName string `json:"configMap"`
	// The xDS transport protocol version to use when talking to the discovery service
	// +kubebuilder:validation:Enum=v2;v3
	XdsProtocolVersion string `json:"xdsProtocolVersion"`
	// ConfigFile is the path of envoy's bootstrap config file
	ConfigFile *string `json:"configFile,omitempty"`
	// ResourcesDir is the path where resource files are loaded from. It is used to
	// load discovery messages directly from the filesystem, for example in order to be able
	// to bootstrap certificates and support rotation when they are modified.
	ResourcesDir *string `json:"resourceDir,omitempty"`
}

// ClientCertificate allows specifying options for the
// client certificate used to authenticate with the discovery
// service
type ClientCertificate struct {
	// The Secret where the certificate will be stored
	SecretName string `json:"secretName"`
	// The requested ‘duration’ (i.e. lifetime) of the Certificate
	Duration metav1.Duration `json:"duration"`
}

// EnvoyBootstrapStatus defines the observed state of EnvoyBootstrap
type EnvoyBootstrapStatus struct{}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// EnvoyBootstrap is the Schema for the envoybootstraps API
type EnvoyBootstrap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EnvoyBootstrapSpec   `json:"spec,omitempty"`
	Status EnvoyBootstrapStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvoyBootstrapList contains a list of EnvoyBootstrap
type EnvoyBootstrapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvoyBootstrap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&EnvoyBootstrap{}, &EnvoyBootstrapList{})
}
