package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"


// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Registry struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec RegistrySpec `json:"spec"`
	Status RegistryStatus `json:"status"`
}

type RegistrySpec struct {
	Secret string `json:"secret"`
	Namespaces []NamespaceFilter `json:"namespaces"`
}

type RegistryStatus struct {
	State   string `json:"state"`
	Message string `json:"message"`
}

type NamespaceFilter struct {
	Name string `json:"name"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type RegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Registry `json:"items"`
}
