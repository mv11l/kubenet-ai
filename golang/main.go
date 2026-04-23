package main

import (
	"fmt"

	pods "github.com/mv11l/kubenet-ai/internal/k8s"
)

func main() {
	pods, err := pods.GetPodsFromNs("kube-system")
	if err != nil {
		fmt.Println(err)
	}
	for _, pod := range pods.Items {
		name := pod.Name
		namespace := pod.Namespace
		fmt.Printf("- [%s, %s] \n", namespace, name)
	}
}
