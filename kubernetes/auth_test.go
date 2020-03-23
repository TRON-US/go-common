package kubernetes

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func newFakeClient() *Config {
	k := Config{}
	k.clientset = fake.NewSimpleClientset()
	return &k
}

func TestAuthenticationDefaultError(t *testing.T) {
	t.Parallel()
	k := newFakeClient()
	err := k.Authenticate()
	assert.NotEqual(t, nil, err, "Authentication against fake client should error out")
}

func TestAuthenticationKubeFileError(t *testing.T) {
	t.Parallel()
	k := newFakeClient()
	kubeconfig := "/thefilewhichdoesnotExists"
	os.Setenv("KUBECONFIG", kubeconfig)
	err := k.Authenticate()
	if !fileExists(kubeconfig) {
		assert.NotEqual(t, nil, err, "Kubeconfig env var set but file not present should error out")
	}
}

func TestAuthenticationKubeConfigError(t *testing.T) {
	t.Parallel()
	k := newFakeClient()
	kubeconfig := "./testdata/corruptedkubeconfig"
	os.Setenv("KUBECONFIG", kubeconfig)
	err := k.Authenticate()
	assert.NotEqual(t, nil, err, "Corrupted Kubeconfig should throw an error")
}

func TestAuthenticationKubeError(t *testing.T) {
	t.Parallel()
	k := newFakeClient()
	kubeconfig := "./testdata/samplekubeconfig"
	os.Setenv("KUBECONFIG", kubeconfig)
	err := k.Authenticate()
	assert.Equal(t, nil, err, "Authentication error")
	_, err = k.clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
	assert.NotEqual(t, nil, err, "Fake kubernetes should throw an error")
}
