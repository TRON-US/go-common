package kubernetes

import (
	"os"
	"testing"

	"k8s.io/client-go/kubernetes/fake"
)

func newFakeClient() *Config {
	k := Config{}
	k.clientset = fake.NewSimpleClientset()
	return &k
}

func TestAuthenticationDefaultError(t *testing.T) {
	k := newFakeClient()
	err := k.Authenticate()
	if err == nil {
		t.Error(err)
	}
}

func TestAuthenticationKubeFileError(t *testing.T) {
	k := newFakeClient()
	kubeconfig := "/thefilewhichdoesnotExists"
	os.Setenv("KUBECONFIG", kubeconfig)
	err := k.Authenticate()
	if !fileExists(kubeconfig) {
		if err == nil {
			t.Error("Kubeconfig env var set but file not present should error out")
		}
	}
}

func TestAuthenticationKubeConfigError(t *testing.T) {
	k := newFakeClient()
	kubeconfig := "./testdata/corruptedkubeconfig"
	os.Setenv("KUBECONFIG", kubeconfig)
	err := k.Authenticate()
	if err == nil {
		t.Error("Corrupted Kubeconfig should throw an error")
	}
}
