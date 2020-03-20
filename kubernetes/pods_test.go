package kubernetes

import (
	"os"
	"testing"

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

func TestGetActivePodsErrorWithoutNamespace(t *testing.T) {
	k := newFakeClient()
	pods, err := k.GetActivePods("", "")
	if !fileExists("/var/run/secrets/kubernetes.io/serviceaccount/namespace") && err == nil && len(pods) > 0 {
		t.Error("Get Active Pods should return an error on non-kube environments when namespace is not provided")
	}

	if fileExists("/var/run/secrets/kubernetes.io/serviceaccount/namespace") && err != nil {
		t.Error(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func TestGetActivePods(t *testing.T) {
	k := Config{}
	t.Parallel()
	for _, tt := range testdata {
		tt := tt
		t.Run(tt.title, func(tt struct {
			title     string
			namespace string
			labels    string
			podcount  int
		}) func(t *testing.T) {
			return func(t *testing.T) {
				k.clientset = fakeclientset
				pods, err := k.GetActivePods(tt.namespace, tt.labels)
				if err != nil {
					t.Error(err)
				}

				got := len(pods)
				want := tt.podcount
				if got != want {
					t.Errorf("[No. of pods] got '%v' want '%v'", got, want)
				}
			}
		}(tt))
	}
}
