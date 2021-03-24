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

// HostWatcherSpec defines the desired state of HostWatcher
type HostWatcherSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// ConfigMapName is an example field of HostWatcher. Edit HostWatcher_types.go to remove/update
	ConfigMapName string `json:"configMapName,omitempty"`
}

// HostWatcherStatus defines the observed state of HostWatcher
type HostWatcherStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Hosts []Host `json:"hosts,omitempty"`
}

type Host struct {
	URL        string `json:"url,omitempty"`
	Path       string `json:"path,omitempty"`
	StatusCode string `json:"statuscode,omitempty"`
}

// +kubebuilder:object:root=true

// HostWatcher is the Schema for the hostwatchers API
type HostWatcher struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HostWatcherSpec   `json:"spec,omitempty"`
	Status HostWatcherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostWatcherList contains a list of HostWatcher
type HostWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostWatcher `json:"items"`
}

func init() {
	SchemeBuilder.Register(&HostWatcher{}, &HostWatcherList{})
}
