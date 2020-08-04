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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// IstioClusterSpec defines the desired state of IstioCluster
type IstioClusterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	//version, istio version
	Version string `json:"version,omitempty"`
	//ClusterId
	ClusterId string `json:"clusterId,omitempty"`
}

// IstioClusterStatus defines the observed state of IstioCluster
type IstioClusterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Individual status of each component controlled by the operator. The map key is the name of the component.
	ComponentStatus      map[string]*InstallStatus_VersionStatus `json:"componentStatus,omitempty"`
}

// VersionStatus is the status and version of a component.
type InstallStatus_VersionStatus struct {
	Version              string               `json:"version,omitempty"`
	Status               InstallStatus_Status `json:"status,omitempty"`
	Error                string               `json:"error,omitempty"`
}

// Status describes the current state of a component.
type InstallStatus_Status string

const (
	// Component is not present.
	InstallStatus_NONE InstallStatus_Status = "NONE"
	// Component is being updated to a different version.
	InstallStatus_UPDATING InstallStatus_Status = "UPDATING"
	// Controller has started but not yet completed reconciliation loop for the component.
	InstallStatus_RECONCILING InstallStatus_Status = "RECONCILING"
	// Component is healthy.
	InstallStatus_HEALTHY InstallStatus_Status = "HEALTHY"
	// Component is in an error state.
	InstallStatus_ERROR InstallStatus_Status = "ERROR"
)

// +kubebuilder:object:root=true

// IstioCluster is the Schema for the istioclusters API
type IstioCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IstioClusterSpec   `json:"spec,omitempty"`
	Status IstioClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IstioClusterList contains a list of IstioCluster
type IstioClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IstioCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IstioCluster{}, &IstioClusterList{})
}
