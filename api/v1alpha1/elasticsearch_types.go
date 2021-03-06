/*
Copyright 2020 Opstree Solutions.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ElasticsearchSpec defines the desired state of Elasticsearch
type ElasticsearchSpec struct {
	ClusterName     string            `json:"clusterName"`
	Image           string            `json:"image"`
	ImagePullPolicy corev1.PullPolicy `json:"imagePullPolicy,omitempty"`
	Security        Security          `json:"security,omitempty"`
	Plugins         []*string         `json:"plugins,omitempty"`
	Master          NodeSpec          `json:"master,omitempty"`
	Data            NodeSpec          `json:"data,omitempty"`
	Ingestion       NodeSpec          `json:"ingestion,omitempty"`
	Client          NodeSpec          `json:"client,omitempty"`
}

// ElasticsearchStatus defines the observed state of Elasticsearch
type ElasticsearchStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ClusterName  string `json:"clusterName"`
	Master       *int32 `json:"master"`
	Data         *int32 `json:"data"`
	Ingestion    *int32 `json:"ingestion"`
	Client       *int32 `json:"client"`
	ClusterState string `json:"clusterState"`
}

// Security defines the security of elasticsearch
type Security struct {
	TLSEnabled *bool  `json:"tlsEnabled,omitempty"`
	Password   string `json:"password,omitempty"`
}

// NodeSpec define the state of elasticsearch nodes
type NodeSpec struct {
	Enabled           bool               `json:"enabled,omitempty"`
	Count             *int32             `json:"count,omitempty"`
	Resources         *Resources         `json:"resources,omitempty"`
	ExtraEnvVariables *map[string]string `json:"extraEnvVariables,omitempty"`
	Storage           *Storage           `json:"storage,omitempty"`
	JVMOptions        JVMOptions         `json:"jvmOptions,omitempty"`
	Affinity          *corev1.Affinity   `json:"affinity,omitempty"`
}

// JVMOptions define the JVM size for elasticsearch nodes
type JVMOptions struct {
	Max string `json:"Xmx,omitempty"`
	Min string `json:"Xms,omitempty"`
}

// Resources describes requests and limits for the cluster resouces.
type Resources struct {
	ResourceRequests ResourceDescription `json:"requests,omitempty"`
	ResourceLimits   ResourceDescription `json:"limits,omitempty"`
}

// Storage is the inteface to add pvc and pv support in redis
type Storage struct {
	VolumeClaimTemplate corev1.PersistentVolumeClaim `json:"volumeClaimTemplate,omitempty"`
}

// ResourceDescription describes CPU and memory resources defined for a cluster.
type ResourceDescription struct {
	CPU    string `json:"cpu"`
	Memory string `json:"memory"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Elasticsearch is the Schema for the elasticsearches API
type Elasticsearch struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticsearchSpec   `json:"spec,omitempty"`
	Status ElasticsearchStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ElasticsearchList contains a list of Elasticsearch
type ElasticsearchList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Elasticsearch `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Elasticsearch{}, &ElasticsearchList{})
}
