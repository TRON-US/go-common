package kubernetes

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Config) GetActivePods(namespace, labels string) (pods []string, err error) {
	// Get the current namespace, if no namespace is provided
	if len(namespace) == 0 {
		namespaceFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
		namespace, err = getNamespaceFromFile(namespaceFile)
		if err != nil {
			return nil, err
		}
	}

	// Options greps the output with the conditions provided.
	options := metav1.ListOptions{LabelSelector: labels}

	// Query Kubernetes API to get the list
	podObjects, err := k.clientset.CoreV1().Pods(namespace).List(options)
	if err != nil {
		return
	}

	// Get the name of pods from podObjects.
	// Returns emtpy slice if no pods are found.
	for _, podObject := range podObjects.Items {
		// List of pods is filtered by labels and status=Running.
		if podObject.Status.Phase == corev1.PodRunning {
			pods = append(pods, podObject.GetName())
		}
	}

	return
}

func getNamespaceFromFile(namespaceFile string) (namespace string, err error) {
	if !fileExists(namespaceFile) {
		return "", errors.New("Cannot get namespace from file")
	}
	nsBytes, nsErr := ioutil.ReadFile(namespaceFile)
	if nsErr != nil {
		return "", nsErr
	}
	return strings.TrimSpace(string(nsBytes)), nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
