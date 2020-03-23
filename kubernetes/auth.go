package kubernetes

import (
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Config contains the kubernetes clientset and configs.
type Config struct {
	clientset kubernetes.Interface
}

// authenticate authenticates and authorizes the client.
func (k *Config) Authenticate() (err error) {
	var config *rest.Config
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if len(kubeconfigPath) > 0 {
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	} else {
		config, err = rest.InClusterConfig()
	}
	if err != nil {
		return err
	}
	k.clientset, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	return nil
}
