package kubernetes

import (
	"io/ioutil"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *Config) GetActivePods(namespace, labels string) (pods []string, err error) {
	// Get the current namespace, if no namespace is provided
	if len(namespace) == 0 {
		nsBytes, nsErr := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
		if nsErr != nil {
			return nil, nsErr
		}
		namespace = strings.TrimSpace(string(nsBytes))
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
