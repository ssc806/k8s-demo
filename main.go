
package main

import (
	//"context"
	"flag"
	"fmt"
	"path/filepath"
	//"time"

	//"k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func GetKubeClient() (*kubernetes.Clientset, error) {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset, nil
}

func main() {
    clientset, _ := GetKubeClient()

    fmt.Printf("Create PV......\n")
    CreatePV (clientset)

    fmt.Printf("\n\nList PV......\n")
    ListPV (clientset)

    fmt.Printf("\n\nGet PV......\n")
    GetPV (clientset, "foo-pv")

    //DeletePV (clientset, "foo-pv")

    fmt.Printf("\n\nCreate PVC By Existing PV......\n")
    CreatePVCByExistingPV (clientset, "foo-pv", "default")

    fmt.Printf("\n\nList PVC......\n")
    ListPVC (clientset, "default")

    fmt.Printf("\n\nGet PVC......\n")
    GetPVC (clientset, "foo-pvc", "default")

    //DeletePVC (clientset, "nfs-pv", "default") 
}
