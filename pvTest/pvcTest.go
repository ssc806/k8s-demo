
package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apimachinery/pkg/api/resource"
)

func ListPVC(clientset *kubernetes.Clientset, nsName string) {
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims(nsName).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	for _, pvc := range pvcs.Items {
	    fmt.Printf("%+v\n", pvc)		
	}
}


// Need to wait the status to be Bound
// By PV need to specify the storage limit From PV capacity
func CreatePVCByExistingPV(clientset *kubernetes.Clientset, pvName string, nsName string) {
    scName := ""

	pvc := &corev1.PersistentVolumeClaim {
		ObjectMeta: metav1.ObjectMeta {
			Name: "foo-pvc",
			Labels: map[string]string{
				"type": "demo",
			},			
		},

		Spec: corev1.PersistentVolumeClaimSpec {
			Resources: corev1.ResourceRequirements{
				Limits: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse("2Gi"),
				},
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse("2Gi"),
				},		
			},

			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
            StorageClassName: &scName,
            VolumeName: pvName,
		},
	}

	_, err := clientset.CoreV1().PersistentVolumeClaims(nsName).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}


// Need to wait the status to be Bound
func CreatePVCByStorageClass(clientset *kubernetes.Clientset, scName string, nsName string) {
	pvc := &corev1.PersistentVolumeClaim {
		ObjectMeta: metav1.ObjectMeta {
			Name: "bar",
			Labels: map[string]string{
				"type": "demo",
			},			
		},

		Spec: corev1.PersistentVolumeClaimSpec {
			Resources: corev1.ResourceRequirements{
				Limits: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse("1Gi"),
				},
				Requests: map[corev1.ResourceName]resource.Quantity{
					corev1.ResourceStorage: resource.MustParse("1Gi"),
				},		
			},

			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
			StorageClassName: &scName,
		},
	}

	_, err := clientset.CoreV1().PersistentVolumeClaims(nsName).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
}

//Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PersistentVolume, error)
func GetPVC(clientset *kubernetes.Clientset, name string, nsName string){
	pvc, err := clientset.CoreV1().PersistentVolumeClaims(nsName).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}


	fmt.Printf("%+v\n", pvc)		
}


//Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
func DeletePVC(clientset *kubernetes.Clientset, name string, nsName string) {
	err := clientset.CoreV1().PersistentVolumeClaims(nsName).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		panic(err.Error())
	} else {
	    fmt.Printf("Delete PVC successfully.")
	}
     
}





