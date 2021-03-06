package service

import (
	"time"

	"github.com/giantswarm/versionbundle"
)

func newVersionBundles() []versionbundle.Bundle {
	return []versionbundle.Bundle{
		{
			Changelogs: []versionbundle.Changelog{
				{
					Component:   "calico",
					Description: "Calico version updated.",
					Kind:        "changed",
				},
				{
					Component:   "docker",
					Description: "Docker version updated.",
					Kind:        "changed",
				},
				{
					Component:   "etcd",
					Description: "Etcd version updated.",
					Kind:        "changed",
				},
				{
					Component:   "kubedns",
					Description: "KubeDNS version updated.",
					Kind:        "changed",
				},
				{
					Component:   "kubernetes",
					Description: "Kubernetes version updated.",
					Kind:        "changed",
				},
				{
					Component:   "nginx-ingress-controller",
					Description: "Nginx-ingress-controller version updated.",
					Kind:        "changed",
				},
			},
			Components: []versionbundle.Component{
				{
					Name:    "calico",
					Version: "2.6.2",
				},
				{
					Name:    "docker",
					Version: "1.12.6",
				},
				{
					Name:    "etcd",
					Version: "3.2.7",
				},
				{
					Name:    "kubedns",
					Version: "1.14.5",
				},
				{
					Name:    "kubernetes",
					Version: "1.8.1",
				},
				{
					Name:    "nginx-ingress-controller",
					Version: "0.9.0",
				},
			},
			Dependencies: []versionbundle.Dependency{},
			Deprecated:   true,
			Name:         "aws-operator",
			Time:         time.Date(2017, time.November, 29, 16, 16, 0, 0, time.UTC),
			Version:      "0.1.0",
			WIP:          false,
		},
		{
			Changelogs: []versionbundle.Changelog{
				{
					Component:   "cloudformation",
					Description: "First version of Cloud Formation resources.",
					Kind:        "added",
				},
			},
			Components: []versionbundle.Component{
				{
					Name:    "calico",
					Version: "2.6.2",
				},
				{
					Name:    "docker",
					Version: "1.12.6",
				},
				{
					Name:    "etcd",
					Version: "3.2.7",
				},
				{
					Name:    "kubedns",
					Version: "1.14.5",
				},
				{
					Name:    "kubernetes",
					Version: "1.8.1",
				},
				{
					Name:    "nginx-ingress-controller",
					Version: "0.9.0",
				},
			},
			Dependencies: []versionbundle.Dependency{},
			Deprecated:   true,
			Name:         "aws-operator",
			Time:         time.Date(2017, time.November, 29, 16, 16, 0, 0, time.UTC),
			Version:      "0.2.0",
			WIP:          true,
		},
		{
			Changelogs: []versionbundle.Changelog{
				{
					Component:   "kubernetes",
					Description: "Updated to kubernetes 1.8.4. Fixes a goroutine leak in the k8s api.",
					Kind:        "changed",
				},
			},
			Components: []versionbundle.Component{
				{
					Name:    "calico",
					Version: "2.6.2",
				},
				{
					Name:    "docker",
					Version: "1.12.6",
				},
				{
					Name:    "etcd",
					Version: "3.2.7",
				},
				{
					Name:    "kubedns",
					Version: "1.14.5",
				},
				{
					Name:    "kubernetes",
					Version: "1.8.4",
				},
				{
					Name:    "nginx-ingress-controller",
					Version: "0.9.0",
				},
			},
			Dependencies: []versionbundle.Dependency{},
			Deprecated:   false,
			Name:         "aws-operator",
			Time:         time.Date(2017, time.December, 5, 13, 00, 0, 0, time.UTC),
			Version:      "1.0.0",
			WIP:          false,
		},
		{
			Changelogs: []versionbundle.Changelog{
				{
					Component:   "CloudFormation",
					Description: "All AWS resources use CloudFormation except KMS and S3.",
					Kind:        versionbundle.KindAdded,
				},
				{
					Component:   "Kubernetes",
					Description: "Updated to Kubernetes 1.9.2.",
					Kind:        versionbundle.KindChanged,
				},
				{
					Component:   "Kubernetes",
					Description: "Switched to vanilla (previously CoreOS) hyperkube image.",
					Kind:        versionbundle.KindChanged,
				},
				{
					Component:   "Docker",
					Description: "Updated to 17.09.0-ce.",
					Kind:        versionbundle.KindChanged,
				},
				{
					Component:   "Calico",
					Description: "Updated to 3.0.1.",
					Kind:        versionbundle.KindChanged,
				},
				{
					Component:   "CoreDNS",
					Description: "Version 1.0.5 replaces kube-dns.",
					Kind:        versionbundle.KindAdded,
				},
				{
					Component:   "Nginx Ingress Controller",
					Description: "Updated to 0.10.2.",
					Kind:        versionbundle.KindChanged,
				},
			},
			Components: []versionbundle.Component{
				{
					Name:    "calico",
					Version: "3.0.1",
				},
				{
					Name:    "docker",
					Version: "17.09.0",
				},
				{
					Name:    "etcd",
					Version: "3.2.7",
				},
				{
					Name:    "coredns",
					Version: "1.0.5",
				},
				{
					Name:    "kubernetes",
					Version: "1.9.2",
				},
				{
					Name:    "nginx-ingress-controller",
					Version: "0.10.2",
				},
			},
			Dependencies: []versionbundle.Dependency{},
			Deprecated:   false,
			Name:         "aws-operator",
			Time:         time.Date(2018, time.January, 22, 16, 00, 0, 0, time.UTC),
			Version:      "2.0.0",
			WIP:          false,
		},
	}
}
