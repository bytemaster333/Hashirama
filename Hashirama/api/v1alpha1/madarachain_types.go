/*
Copyright 2025.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MadaraChainSpec defines the desired state of MadaraChain
type MadaraChainSpec struct {
	// ChainID is the custom name of the L3 chain.
	ChainID string `json:"chainID"`

	// Replicas is the number of nodes (default to 1).
	// +kubebuilder:default=1
	Replicas int32 `json:"replicas,omitempty"`

	// Port is the RPC port to expose (default to 9944).
	// +kubebuilder:default=9944
	Port int32 `json:"port,omitempty"`

	// Image is the Madara docker image.
	// +kubebuilder:default="ghcr.io/madara-alliance/madara:latest"
	Image string `json:"image,omitempty"`

	// Network is the Starknet network to connect to (mainnet, sepolia, devnet).
	// +kubebuilder:default="sepolia"
	Network string `json:"network,omitempty"`
}

// MadaraChainStatus defines the observed state of MadaraChain.
type MadaraChainStatus struct {
	// NodesRunning is the number of running nodes.
	NodesRunning int32 `json:"nodesRunning,omitempty"`

	// conditions represent the current state of the MadaraChain resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	// +listType=map
	// +listMapKey=type
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MadaraChain is the Schema for the madarachains API
type MadaraChain struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitzero"`

	// spec defines the desired state of MadaraChain
	// +required
	Spec MadaraChainSpec `json:"spec"`

	// status defines the observed state of MadaraChain
	// +optional
	Status MadaraChainStatus `json:"status,omitzero"`
}

// +kubebuilder:object:root=true

// MadaraChainList contains a list of MadaraChain
type MadaraChainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitzero"`
	Items           []MadaraChain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MadaraChain{}, &MadaraChainList{})
}
