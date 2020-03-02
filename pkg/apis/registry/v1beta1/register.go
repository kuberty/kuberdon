package v1beta1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/kuberty/kuberdon/pkg/apis/registry"
)

// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API
var SchemeGroupVersion = schema.GroupVersion{
	Group:   registry.GroupName,
	Version: "v1beta1",
}

var AddToScheme = runtime.NewSchemeBuilder(addKnownTypes).AddToScheme

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&Registry{},
	)

	// register the type in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
