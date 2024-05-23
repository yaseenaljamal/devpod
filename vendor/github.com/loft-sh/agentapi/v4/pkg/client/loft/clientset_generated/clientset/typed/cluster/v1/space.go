// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/agentapi/v4/pkg/apis/loft/cluster/v1"
	scheme "github.com/loft-sh/agentapi/v4/pkg/client/loft/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SpacesGetter has a method to return a SpaceInterface.
// A group's client should implement this interface.
type SpacesGetter interface {
	Spaces() SpaceInterface
}

// SpaceInterface has methods to work with Space resources.
type SpaceInterface interface {
	Create(ctx context.Context, space *v1.Space, opts metav1.CreateOptions) (*v1.Space, error)
	Update(ctx context.Context, space *v1.Space, opts metav1.UpdateOptions) (*v1.Space, error)
	UpdateStatus(ctx context.Context, space *v1.Space, opts metav1.UpdateOptions) (*v1.Space, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Space, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SpaceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Space, err error)
	SpaceExpansion
}

// spaces implements SpaceInterface
type spaces struct {
	client rest.Interface
}

// newSpaces returns a Spaces
func newSpaces(c *ClusterV1Client) *spaces {
	return &spaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the space, and returns the corresponding space object, and an error if there is any.
func (c *spaces) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Space, err error) {
	result = &v1.Space{}
	err = c.client.Get().
		Resource("spaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Spaces that match those selectors.
func (c *spaces) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SpaceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SpaceList{}
	err = c.client.Get().
		Resource("spaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested spaces.
func (c *spaces) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("spaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a space and creates it.  Returns the server's representation of the space, and an error, if there is any.
func (c *spaces) Create(ctx context.Context, space *v1.Space, opts metav1.CreateOptions) (result *v1.Space, err error) {
	result = &v1.Space{}
	err = c.client.Post().
		Resource("spaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(space).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a space and updates it. Returns the server's representation of the space, and an error, if there is any.
func (c *spaces) Update(ctx context.Context, space *v1.Space, opts metav1.UpdateOptions) (result *v1.Space, err error) {
	result = &v1.Space{}
	err = c.client.Put().
		Resource("spaces").
		Name(space.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(space).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *spaces) UpdateStatus(ctx context.Context, space *v1.Space, opts metav1.UpdateOptions) (result *v1.Space, err error) {
	result = &v1.Space{}
	err = c.client.Put().
		Resource("spaces").
		Name(space.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(space).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the space and deletes it. Returns an error if one occurs.
func (c *spaces) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("spaces").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *spaces) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("spaces").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched space.
func (c *spaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Space, err error) {
	result = &v1.Space{}
	err = c.client.Patch(pt).
		Resource("spaces").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
