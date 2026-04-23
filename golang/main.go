package main

import (
	"fmt"

	pods "github.com/mv11l/kubenet-ai/internal/k8s"
)

func main() {
	pods := pods.GetPodsFromNs("kube-system")

	for _, pod := range pods {
		name := pod.GetName()
		test := pod.GetLabels()
		namespace := pod.GetNamespace()
		fmt.Printf("- [%s, %s] %s\n", namespace, test, name)
	}
}
