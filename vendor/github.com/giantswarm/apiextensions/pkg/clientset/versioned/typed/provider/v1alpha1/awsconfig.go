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
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AWSConfigsGetter has a method to return a AWSConfigInterface.
// A group's client should implement this interface.
type AWSConfigsGetter interface {
	AWSConfigs(namespace string) AWSConfigInterface
}

// AWSConfigInterface has methods to work with AWSConfig resources.
type AWSConfigInterface interface {
	Create(*v1alpha1.AWSConfig) (*v1alpha1.AWSConfig, error)
	Update(*v1alpha1.AWSConfig) (*v1alpha1.AWSConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.AWSConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.AWSConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AWSConfig, err error)
	AWSConfigExpansion
}

// aWSConfigs implements AWSConfigInterface
type aWSConfigs struct {
	client rest.Interface
	ns     string
}

// newAWSConfigs returns a AWSConfigs
func newAWSConfigs(c *ProviderV1alpha1Client, namespace string) *aWSConfigs {
	return &aWSConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSConfig, and returns the corresponding aWSConfig object, and an error if there is any.
func (c *aWSConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.AWSConfig, err error) {
	result = &v1alpha1.AWSConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSConfigs that match those selectors.
func (c *aWSConfigs) List(opts v1.ListOptions) (result *v1alpha1.AWSConfigList, err error) {
	result = &v1alpha1.AWSConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSConfigs.
func (c *aWSConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a aWSConfig and creates it.  Returns the server's representation of the aWSConfig, and an error, if there is any.
func (c *aWSConfigs) Create(aWSConfig *v1alpha1.AWSConfig) (result *v1alpha1.AWSConfig, err error) {
	result = &v1alpha1.AWSConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsconfigs").
		Body(aWSConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aWSConfig and updates it. Returns the server's representation of the aWSConfig, and an error, if there is any.
func (c *aWSConfigs) Update(aWSConfig *v1alpha1.AWSConfig) (result *v1alpha1.AWSConfig, err error) {
	result = &v1alpha1.AWSConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsconfigs").
		Name(aWSConfig.Name).
		Body(aWSConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the aWSConfig and deletes it. Returns an error if one occurs.
func (c *aWSConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aWSConfig.
func (c *aWSConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.AWSConfig, err error) {
	result = &v1alpha1.AWSConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
