package kubernetes

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

var fakeclientset = fake.NewSimpleClientset(
	&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod-1-no-labels",
			Namespace:   "testing",
			Annotations: map[string]string{},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
		},
	},
	&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod-2",
			Namespace:   "testing",
			Annotations: map[string]string{},
			Labels: map[string]string{
				"pod2label1":  "value1",
				"pod2label2":  "value2",
				"commonlabel": "commonvalue",
			},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
		},
	},
	&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod-3",
			Namespace:   "testing",
			Annotations: map[string]string{},
			Labels: map[string]string{
				"pod3label1":  "value1",
				"pod3label2":  "value2",
				"commonlabel": "commonvalue",
			},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
		},
	},
	&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod-4",
			Namespace:   "testing",
			Annotations: map[string]string{},
			Labels: map[string]string{
				"commonlabel": "commonvalue",
			},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodRunning,
		},
	},
	&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:        "pod-5",
			Namespace:   "testing",
			Annotations: map[string]string{},
			Labels: map[string]string{
				"commonlabel": "commonvalue",
			},
		},
		Status: corev1.PodStatus{
			Phase: corev1.PodPending,
		},
	},
)

var testdata = []struct {
	title     string
	namespace string
	labels    string
	podcount  int
}{
	{
		title:     "Get All Active Pods",
		namespace: "testing",
		labels:    "",
		podcount:  4,
	},
	{
		title:     "Get Active Pods with unique labels",
		namespace: "testing",
		labels:    "pod2label2=value2",
		podcount:  1,
	},
	{
		title:     "Get Active Pods with common labels",
		namespace: "testing",
		labels:    "commonlabel=commonvalue",
		podcount:  3,
	},
}

var nstestdata = []struct {
	title         string
	namespaceFile string
	namespace     string
	err           error
}{
	{
		title:         "Namespace file exists",
		namespaceFile: "./testdata/namespace",
		namespace:     "test",
		err:           nil,
	},
	{
		title:         "Namespace file does not exists",
		namespaceFile: "./testdata/nonexistentfile",
		namespace:     "",
		err:           errors.New("Cannot get namespace from file"),
	},
	{
		title:         "Empty namespace",
		namespaceFile: "./testdata/emptynamespace",
		namespace:     "",
		err:           errors.New("Cannot get namespace from file"),
	},
}

func TestGetActivePodsErrorWithoutNamespace(t *testing.T) {
	t.Parallel()
	k := newFakeClient()
	pods, err := k.GetActivePods("", "")
	namespaceFile := "/var/run/secrets/kubernetes.io/serviceaccount/namespace"
	if fileExists(namespaceFile) {
		assert.Equal(t, nil, err, "Error getting namespace from file")
	} else {
		assert.Equal(t, 0, len(pods), "Get Active Pods should return an error on non-kube environments when namespace is not provided")
	}
}

func TestGetNamespaceFromFile(t *testing.T) {
	t.Parallel()
	for _, tt := range nstestdata {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			namespace, err := getNamespaceFromFile(tt.namespaceFile)
			assert.Equal(t, tt.namespace, namespace, "Error getting namespace from file")
			assert.Equal(t, tt.err, err, "Error getting namespace from file")
		})
	}
}

func TestGetActivePods(t *testing.T) {
	t.Parallel()
	k := Config{}
	for _, tt := range testdata {
		tt := tt
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()
			k.clientset = fakeclientset
			pods, err := k.GetActivePods(tt.namespace, tt.labels)
			assert.Equal(t, nil, err, "Error executing GetActivePods")

			got := len(pods)
			want := tt.podcount
			assert.Equal(t, want, got, "Number of pods")
		})
	}
}
