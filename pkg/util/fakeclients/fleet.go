package fakeclients

import (
	fleetv1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FleetBundleCache func(namespace, name string) (*fleetv1alpha1.Bundle, error)

func (c FleetBundleCache) Get(namespace, name string) (*fleetv1alpha1.Bundle, error) {
	return c(namespace, name)
}

type FleetClusterCache func(namespace, name string) (*fleetv1alpha1.Cluster, error)

func (c FleetClusterCache) Get(namespace, name string) (*fleetv1alpha1.Cluster, error) {
	return c(namespace, name)
}

type BundleDeploymentClient func(namespace, name string, options *metav1.DeleteOptions) error

func (c BundleDeploymentClient) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return c(namespace, name, options)
}

type BundleDeploymentCache func(namespace, name string) (*fleetv1alpha1.BundleDeployment, error)

func (c BundleDeploymentCache) Get(namespace, name string) (*fleetv1alpha1.BundleDeployment, error) {
	return c(namespace, name)
}
