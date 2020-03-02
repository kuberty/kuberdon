package main
// Attribution to martin-helmich: https://github.com/martin-helmich/kubernetes-crd-example
// Attribution to Kubernetes	: https://github.com/kubernetes/sample-controller/blob/master/main.go

import (
	"flag"
	"k8s.io/client-go/tools/cache"
	"log"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	clienset "github.com/kuberty/kuberdon/pkg/client/clientset/versioned"
	informers "github.com/kuberty/kuberdon/pkg/client/informers/externalversions"

)

var kubeconfig string


func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {
	var config *rest.Config
	var err error


	if kubeconfig == "" {
		log.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		log.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err)
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}


	exampleClient, err := clienset.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	// define controller
	mockSignalHandler := make(chan struct{})
	kubeInformerFactory.Start(mockSignalHandler)
	exampleInformerFactory.Start(mockSignalHandler)

	exampleInformerFactory.Kuberdon().V1beta1().Registries().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			log.Printf("Added: %v", obj)
		} ,
		UpdateFunc: func(old, new interface{}) {
			log.Printf("Changed from %v to %v", old, new)
		},
	})
}
