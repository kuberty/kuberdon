package main
// Attribution to martin-helmich: https://github.com/martin-helmich/kubernetes-crd-example

import (
	"flag"

	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	clientV1beta1 "github.com/kuberty/kuberdon/pkg/client/clientset/versioned"

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

	clientSet, err := clientV1beta1.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	registries, err := clientSet.Registries("").List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("registries found: %+v\n", registries)

	store := WatchResources(clientSet)

	for {
		registriesFromStore := store.List()
		fmt.Printf("Registries in store: %+v\n", registriesFromStore))

		time.Sleep(2 * time.Second)
	}


}
