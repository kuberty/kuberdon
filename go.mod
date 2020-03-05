module kuberdon

go 1.13

require (
	github.com/kuberty/kuberdon v0.0.0
	k8s.io/api v0.17.3
	k8s.io/apimachinery v0.17.3
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/code-generator v0.17.3
)

replace github.com/kuberty/kuberdon v0.0.0 => ./
