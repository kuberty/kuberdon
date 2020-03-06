module kuberdon

go 1.13

require (
	github.com/kuberty/kuberdon v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
	k8s.io/code-generator v0.0.0
)

replace github.com/kuberty/kuberdon v0.0.0 => ./

replace k8s.io/api v0.0.0 => ../../../k8s.io/api/

replace k8s.io/apimachinery v0.0.0 => ../../../k8s.io/apimachinery/

replace k8s.io/client-go v0.0.0 => ../../../k8s.io/client-go/

replace k8s.io/code-generator v0.0.0 => ../../../k8s.io/code-generator/
