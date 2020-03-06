package main
// Attribution to martin-helmich: https://github.com/martin-helmich/kubernetes-crd-example
// Attribution to Kubernetes	: https://github.com/kubernetes/sample-controller/blob/master/main.go

import (
	"flag"
	"kuberdon/pkg/listener"
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

	//exampleInformerFactory.Kuberdon().V1beta1().Registries().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		registry := obj.(*v1beta1.Registry)
	//		log.Printf("Added registry: %v", registry.Name)
	//	} ,
	//	UpdateFunc: func(old, new interface{}) {
	//		registry := new.(*v1beta1.Registry).DeepCopy()
	//		namespaces, err := kubeInformerFactory.Core().V1().Namespaces().Lister().List(labels.Everything())
	//		if err != nil {
	//			log.Panic(err)
	//		}
	//		if registry.GetLabels() == nil{
	//			registry.SetLabels(map[string]string{})
	//		}
	//		labels := registry.GetLabels()
	//		labels["a"] = "b"
	//		registry.SetLabels(labels)
	//		registry.SetNamespace("test")
	//		r, err := exampleClient.KuberdonV1beta1().Registries().Update(context.TODO(), registry, metav1.UpdateOptions{})
	//		log.Printf("R: %v, E: %v", r, err)
	//		for _, n := range namespaces {
	//			log.Printf("namespace: %v", n.Name)
	//		}
	//		log.Printf("Changed registry to: %v", registry.Spec.Namespaces[0].Name)
	//	},
	//})
	//
	//kubeInformerFactory.Core().V1().Namespaces().Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
	//	AddFunc: func(obj interface{}) {
	//		namespace := obj.(*v1.Namespace)
	//		log.Printf("Added namespace: %v", namespace.Name)
	//	} ,
	//	UpdateFunc: func(old, new interface{}) {
	//		namespace := new.(*v1.Namespace)
	//		log.Printf("Changed to namespace: %v", namespace.Name)
	//	},
	//})

	c:=	listener.GetEventStream(exampleInformerFactory, kubeInformerFactory)
	kubeInformerFactory.Start(mockSignalHandler)
	exampleInformerFactory.Start(mockSignalHandler)
	for range c {
		print("Updated")
	}
	for {
		time.Sleep(10 * time.Second)
	}
}
