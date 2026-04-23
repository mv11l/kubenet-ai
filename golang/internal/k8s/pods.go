// Package pods implements utility to get and list pods
package pods

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPodsFromNs(ns string) (*v1.PodList, error) {
	homeDir, err := homeDir()
	if err != nil {
		return nil, fmt.Errorf("failed get homeDir: %w", err)
	}
	kubeconfig := filepath.Join(homeDir, ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return nil, fmt.Errorf("failed to build kubeconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to set client: %w", err)
	}

	pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	if errors.IsNotFound(err) {
		return nil, fmt.Errorf("Pod not found: %w", err)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		return nil, fmt.Errorf("Error getting pod %v\n", statusError.ErrStatus.Message)
	} else if err != nil {
		return nil, fmt.Errorf("Error getting pod %v\n", err.Error())
	}

	return pods, nil
}

func homeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home, nil
}
