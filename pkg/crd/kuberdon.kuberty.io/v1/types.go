package v1

type Registry struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	metav1.ObjectMeta `json:"metadata"`

	Spec CakeSpec `json:"spec"`
	Status CakeStatus `json:"status"`
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
