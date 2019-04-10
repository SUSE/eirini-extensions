// Package testing contains methods to create test data. It's a seaparate
// package to avoid import cycles. Helper functions can be found in the package
// `testhelper`.
package testing

import (
	"context"

	operator_testing "code.cloudfoundry.org/cf-operator/testing"
	corev1 "k8s.io/api/core/v1"
)

// NewCatalog returns a Catalog, our helper for test cases
func NewCatalog() Catalog {
	return Catalog{Catalog: &operator_testing.Catalog{}}
}

// NewContext returns a non-nil empty context, for usage when it is unclear
// which context to use.  Mostly used in tests.
func NewContext() context.Context {
	return operator_testing.NewContext()
}

// Catalog provides several instances for test, based on the cf-operator's catalog
type Catalog struct{ *operator_testing.Catalog }

// PodWithVcapServices generates a labeled pod with VCAP_SERVICES environment variable set
func (c *Catalog) PodWithVcapServices(name string, labels map[string]string, vcapServices string) corev1.Pod {

	pod := c.Catalog.LabeledPod(name, labels)
	pod.Spec.Containers[0].Env = []corev1.EnvVar{
		corev1.EnvVar{
			Name:  "VCAP_SERVICES",
			Value: vcapServices,
		}}

	return pod
}

// DefaultEiriniAppPod generates an Eirini Application pod with VCAP_SERVICES environment variable set
func (c *Catalog) DefaultEiriniAppPod(name string, vcapServices string) corev1.Pod {
	return c.PodWithVcapServices(name, map[string]string{"source_type": "APP"}, vcapServices)
}

// SimplePersiApp generates an Eirini Application pod which requires persistent volume (1 volume)
func (c *Catalog) SimplePersiApp(name string) corev1.Pod {
	return c.DefaultEiriniAppPod(name, `{"eirini-persi": [	  {
		"credentials": {},
		"label": "eirini-persi",
		"name": "my-instance",
		"plan": "hostpath",
		"tags": [
			"erini",
			"kubernetes",
			"storage"
		],
		"volume_mounts": [
		  {
			"container_dir": "/var/vcap/data/de847d34-bdcc-4c5d-92b1-cf2158a15b47",
			"device_type": "shared",
			"mode": "rw",
			"device": {
				"volume_id": "the-volume-id"
			}
		  }
		]
	  }
	]
}`)
}

// MultipleVolumePersiApp generates an Eirini Application pod which requires persistent volume (3 volumes)
func (c *Catalog) MultipleVolumePersiApp(name string) corev1.Pod {
	return c.DefaultEiriniAppPod(name, `{"eirini-persi": [	  {
		"credentials": {},
		"label": "eirini-persi",
		"name": "my-instance",
		"plan": "hostpath",
		"tags": [
			"erini",
			"kubernetes",
			"storage"
		],
		"volume_mounts": [
			{
				"container_dir": "/var/vcap/data/de847d34-bdcc-4c5d-92b1-cf2158a15b47",
				"device_type": "shared",
				"mode": "rw",
				"device": {
					"volume_id": "the-volume-id1"
				}
			},
			{
				"container_dir": "/var/vcap/data/de847d34-bdcc-4c5d-92b1-cf2158a15b47",
				"device_type": "shared",
				"mode": "rw",
				"device": {
					"volume_id": "the-volume-id2"
				}
			},
			{
				"container_dir": "/var/vcap/data/de847d34-bdcc-4c5d-92b1-cf2158a15b47",
				"device_type": "shared",
				"mode": "rw",
				"device": {
					"volume_id": "the-volume-id3"
				}
			}
		]
	  }
	]
}`)
}
