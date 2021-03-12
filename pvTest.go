
package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/api/resource"
)

func ListPV(clientset *kubernetes.Clientset) {
	pvs, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pv := range pvs.Items {
	    fmt.Printf("%+v\n", pv)		
	}
}

func ListAvailablePV(){
	
}

func CreatePV(clientset *kubernetes.Clientset) {
    fs := corev1.PersistentVolumeFilesystem

	pv := &corev1.PersistentVolume {
		ObjectMeta: metav1.ObjectMeta {
			Name: "foo-pv",
			Labels: map[string]string{
				"type": "demo",
			},			
		},

		Spec: corev1.PersistentVolumeSpec {
			Capacity: corev1.ResourceList{
				corev1.ResourceName(corev1.ResourceStorage): resource.MustParse("1Gi"),
			},

			PersistentVolumeSource: corev1.PersistentVolumeSource {
				NFS: &corev1.NFSVolumeSource{
					Server: "172.20.45.125",
					Path: "/data/nfs",
					ReadOnly: false,			
				},
			},

            // ReadWriteOnce / ReadOnlyMany / ReadWriteMany
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
			PersistentVolumeReclaimPolicy: "Retain",
			//StorageClassName: scName,
			VolumeMode: &fs,
            MountOptions: []string{"hard", "nfsvers=4.1"},
		},
	}

	_, err := clientset.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

//Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PersistentVolume, error)
func GetPV(clientset *kubernetes.Clientset, name string){
	pv, err := clientset.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}


	fmt.Printf("%+v\n", pv)		
}


func GetCapacityOfPV(){

}



//Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
func DeletePV(clientset *kubernetes.Clientset, name string) {
	err := clientset.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	} else {
	    fmt.Printf("Delete PV successfully.")
	}
     
}





