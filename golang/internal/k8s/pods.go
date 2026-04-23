// Package pods implements utility to get and list pods
package pods

import (
	"context"
	"os"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPodsFromNs(ns string) []unstructured.Unstructured {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}

	podList, err := dynClient.Resource(gvr).
		Namespace(ns).
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	return podList.Items
}

func homeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic("Could not determine home directory")
	}
	return home
}
