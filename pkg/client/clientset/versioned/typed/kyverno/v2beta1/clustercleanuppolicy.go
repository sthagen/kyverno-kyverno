/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v2beta1

import (
	"context"
	"time"

	v2beta1 "github.com/kyverno/kyverno/api/kyverno/v2beta1"
	scheme "github.com/kyverno/kyverno/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterCleanupPoliciesGetter has a method to return a ClusterCleanupPolicyInterface.
// A group's client should implement this interface.
type ClusterCleanupPoliciesGetter interface {
	ClusterCleanupPolicies() ClusterCleanupPolicyInterface
}

// ClusterCleanupPolicyInterface has methods to work with ClusterCleanupPolicy resources.
type ClusterCleanupPolicyInterface interface {
	Create(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.CreateOptions) (*v2beta1.ClusterCleanupPolicy, error)
	Update(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.UpdateOptions) (*v2beta1.ClusterCleanupPolicy, error)
	UpdateStatus(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.UpdateOptions) (*v2beta1.ClusterCleanupPolicy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2beta1.ClusterCleanupPolicy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2beta1.ClusterCleanupPolicyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterCleanupPolicy, err error)
	ClusterCleanupPolicyExpansion
}

// clusterCleanupPolicies implements ClusterCleanupPolicyInterface
type clusterCleanupPolicies struct {
	client rest.Interface
}

// newClusterCleanupPolicies returns a ClusterCleanupPolicies
func newClusterCleanupPolicies(c *KyvernoV2beta1Client) *clusterCleanupPolicies {
	return &clusterCleanupPolicies{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterCleanupPolicy, and returns the corresponding clusterCleanupPolicy object, and an error if there is any.
func (c *clusterCleanupPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2beta1.ClusterCleanupPolicy, err error) {
	result = &v2beta1.ClusterCleanupPolicy{}
	err = c.client.Get().
		Resource("clustercleanuppolicies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterCleanupPolicies that match those selectors.
func (c *clusterCleanupPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2beta1.ClusterCleanupPolicyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2beta1.ClusterCleanupPolicyList{}
	err = c.client.Get().
		Resource("clustercleanuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterCleanupPolicies.
func (c *clusterCleanupPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("clustercleanuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterCleanupPolicy and creates it.  Returns the server's representation of the clusterCleanupPolicy, and an error, if there is any.
func (c *clusterCleanupPolicies) Create(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.CreateOptions) (result *v2beta1.ClusterCleanupPolicy, err error) {
	result = &v2beta1.ClusterCleanupPolicy{}
	err = c.client.Post().
		Resource("clustercleanuppolicies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterCleanupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterCleanupPolicy and updates it. Returns the server's representation of the clusterCleanupPolicy, and an error, if there is any.
func (c *clusterCleanupPolicies) Update(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.UpdateOptions) (result *v2beta1.ClusterCleanupPolicy, err error) {
	result = &v2beta1.ClusterCleanupPolicy{}
	err = c.client.Put().
		Resource("clustercleanuppolicies").
		Name(clusterCleanupPolicy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterCleanupPolicy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *clusterCleanupPolicies) UpdateStatus(ctx context.Context, clusterCleanupPolicy *v2beta1.ClusterCleanupPolicy, opts v1.UpdateOptions) (result *v2beta1.ClusterCleanupPolicy, err error) {
	result = &v2beta1.ClusterCleanupPolicy{}
	err = c.client.Put().
		Resource("clustercleanuppolicies").
		Name(clusterCleanupPolicy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterCleanupPolicy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterCleanupPolicy and deletes it. Returns an error if one occurs.
func (c *clusterCleanupPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("clustercleanuppolicies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterCleanupPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("clustercleanuppolicies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterCleanupPolicy.
func (c *clusterCleanupPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2beta1.ClusterCleanupPolicy, err error) {
	result = &v2beta1.ClusterCleanupPolicy{}
	err = c.client.Patch(pt).
		Resource("clustercleanuppolicies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
