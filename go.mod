module github.com/project-koku/koku-metrics-operator

go 1.13

require (
	github.com/go-logr/logr v0.1.0
	github.com/google/uuid v1.1.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/onsi/ginkgo v1.11.0
	github.com/onsi/gomega v1.8.1
	github.com/openshift/api v0.0.0-20200117162508-e7ccdda6ba67
	github.com/operator-framework/api v0.2.0
	github.com/prometheus/client_golang v1.7.1
	github.com/prometheus/common v0.14.0
	k8s.io/api v0.18.2
	k8s.io/apimachinery v0.19.2
	k8s.io/client-go v0.18.2
	sigs.k8s.io/controller-runtime v0.6.0
)

replace (
	github.com/openshift/api => github.com/openshift/api v0.0.0-20200117162508-e7ccdda6ba67
	github.com/openshift/client-go => github.com/openshift/client-go v0.0.0-20200116152001-92a2713fa240
	k8s.io/api => k8s.io/api v0.18.2
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.2
	k8s.io/apiserver => k8s.io/apiserver v0.18.2
	k8s.io/client-go => k8s.io/client-go v0.18.2
)
