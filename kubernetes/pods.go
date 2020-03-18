package kubernetes

import (
	"io/ioutil"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Config) GetActivePods(namespace, labels string) (pods []string, err error) {
	// Authenticate with kubernetes cluster
	err = k.authenticate()
	if err != nil {
		return
	}

	// Get the current namespace, if no namespace is provided
	if len(namespace) == 0 {
		nsBytes, nsErr := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if nsErr != nil {
			return nil, nsErr
		}
		namespace = strings.TrimSpace(string(nsBytes))
	}

	// Options greps the output with the conditions provided.
	// List of pods is filtered by labels and status=Running.
	options := metav1.ListOptions{LabelSelector: labels, FieldSelector: "status.phase=Running"}

	// Query Kubernetes API to get the list
	podObjects, err := k.clientset.CoreV1().Pods(namespace).List(options)
	if err != nil {
		return
	}

	// Get the name of pods from podObjects.
	// Returns emtpy slice if no pods are found.
	for _, podObject := range podObjects.Items {
		pods = append(pods, podObject.GetName())
	}

	return
}
