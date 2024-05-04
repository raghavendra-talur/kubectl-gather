package gather

import "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

type Options struct {
	Kubeconfig string
	Context    string
	Namespace  string
	Verbose    bool
}

type Addon interface {
	Gather(*unstructured.Unstructured) error
}
