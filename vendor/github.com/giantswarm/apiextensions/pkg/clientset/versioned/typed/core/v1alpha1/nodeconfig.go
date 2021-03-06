/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeConfigsGetter has a method to return a NodeConfigInterface.
// A group's client should implement this interface.
type NodeConfigsGetter interface {
	NodeConfigs(namespace string) NodeConfigInterface
}

// NodeConfigInterface has methods to work with NodeConfig resources.
type NodeConfigInterface interface {
	Create(*v1alpha1.NodeConfig) (*v1alpha1.NodeConfig, error)
	Update(*v1alpha1.NodeConfig) (*v1alpha1.NodeConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodeConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.NodeConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeConfig, err error)
	NodeConfigExpansion
}

// nodeConfigs implements NodeConfigInterface
type nodeConfigs struct {
	client rest.Interface
	ns     string
}

// newNodeConfigs returns a NodeConfigs
func newNodeConfigs(c *CoreV1alpha1Client, namespace string) *nodeConfigs {
	return &nodeConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeConfig, and returns the corresponding nodeConfig object, and an error if there is any.
func (c *nodeConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeConfig, err error) {
	result = &v1alpha1.NodeConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodeconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeConfigs that match those selectors.
func (c *nodeConfigs) List(opts v1.ListOptions) (result *v1alpha1.NodeConfigList, err error) {
	result = &v1alpha1.NodeConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeConfigs.
func (c *nodeConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodeconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a nodeConfig and creates it.  Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *nodeConfigs) Create(nodeConfig *v1alpha1.NodeConfig) (result *v1alpha1.NodeConfig, err error) {
	result = &v1alpha1.NodeConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodeconfigs").
		Body(nodeConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeConfig and updates it. Returns the server's representation of the nodeConfig, and an error, if there is any.
func (c *nodeConfigs) Update(nodeConfig *v1alpha1.NodeConfig) (result *v1alpha1.NodeConfig, err error) {
	result = &v1alpha1.NodeConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodeconfigs").
		Name(nodeConfig.Name).
		Body(nodeConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeConfig and deletes it. Returns an error if one occurs.
func (c *nodeConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodeconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodeconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeConfig.
func (c *nodeConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeConfig, err error) {
	result = &v1alpha1.NodeConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodeconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
