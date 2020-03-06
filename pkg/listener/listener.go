package listener

import (
	"github.com/kuberty/kuberdon/pkg/client/informers/externalversions"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
)

type Listener interface {
	/**
	Messages the returned channel whenever an api update happens to the following resources: [ Kubernetes namespaces, Kuberdon repository, docker registry secret ]

	More formal specification: All though it may update more or less than this, it should at least push an update whenever this controller has to sync.
	*/

	GetEventStream(kuberdonFactory externalversions.SharedInformerFactory, kubeFactory informers.SharedInformerFactory) chan struct{}
}

/**
Testing documentation: To test this class, we need a kubernetes cluster. This will thus be tested in CI once Kuberty comes out.
*/

func GetEventStream(kuberdonFactory externalversions.SharedInformerFactory, kubeFactory informers.SharedInformerFactory) chan struct{} {
	c := make(chan struct{})

	kuberdonFactory.Kuberdon().V1beta1().Registries().Informer().AddEventHandler(getResourceEventHandlerFuncs(c, nil))
	kubeFactory.Core().V1().Namespaces().Informer().AddEventHandler(getResourceEventHandlerFuncs(c, nil))
	kubeFactory.Core().V1().Secrets().Informer().AddEventHandler(getResourceEventHandlerFuncs(c, nil)) //todo: update the filter to specify which secrets we should fetch (eg. only the ones that resemble a Registry)

	return c
}

func getResourceEventHandlerFuncs(c chan struct{}, filter func(obj interface{}) bool) cache.ResourceEventHandlerFuncs {
	if filter == nil {
		filter = func(obj interface{}) bool { return true }
	}

	return cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			if filter(obj) {
				c <- struct{}{}
			}
		},
		UpdateFunc: func(old, new interface{}) {
			if filter(new) || filter(old){
				c <- struct{}{}
			}
		},
		DeleteFunc: func(obj interface{}) {
			if filter(obj) {
				c <- struct{}{}
			}
		},
	}
}
