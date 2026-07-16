package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func buildClient(config *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	return clientset, err
}

func listNodes(clientset *kubernetes.Clientset) error {
	nodes, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{FieldSelector: "metadata.name=ai-infra-lab-worker"})
	if err != nil {
		return err
	}
	for _, node := range nodes.Items {
		// fmt.Println(node.ObjectMeta.Name)
		// // fmt.Printf("%+v\n", node.Status.Capacity)
		// cpu := node.Status.Allocatable[v1.ResourceCPU]
		// fmt.Print("CPU:", cpu)
		// fmt.Println(node.Spec.PodCIDR)

		for resourceName, quantity := range node.Status.Allocatable {
			fmt.Printf(
				"%s = %s\n",
				resourceName,
				quantity.String(),
			)

		}
	}
	return nil
}

func loadConfig() (*rest.Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	kubeconfigpath := filepath.Join(
		home,
		".kube",
		"config",
	)

	return clientcmd.BuildConfigFromFlags("", kubeconfigpath)
}

func main() {

	config, err := loadConfig()
	if err != nil {
		// TODO
	}

	clientset, err := buildClient(config)
	if err != nil {
		// TODO
	}

	err = listNodes(clientset)
	if err != nil {
		// TODO
	}
}
