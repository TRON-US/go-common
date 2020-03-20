package kubernetes

import (
	"k8s.io/client-go/kubernetes/fake"
)

func newFakeClient() *Config {
	k := Config{}
	k.clientset = fake.NewSimpleClientset()
	return &k
}
