// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SelfSubjectAccessReviewsGetter has a method to return a SelfSubjectAccessReviewInterface.
// A group's client should implement this interface.
type SelfSubjectAccessReviewsGetter interface {
	SelfSubjectAccessReviews() SelfSubjectAccessReviewInterface
}

// SelfSubjectAccessReviewInterface has methods to work with SelfSubjectAccessReview resources.
type SelfSubjectAccessReviewInterface interface {
	Create(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.CreateOptions) (*v1.SelfSubjectAccessReview, error)
	Update(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (*v1.SelfSubjectAccessReview, error)
	UpdateStatus(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (*v1.SelfSubjectAccessReview, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SelfSubjectAccessReview, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SelfSubjectAccessReviewList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SelfSubjectAccessReview, err error)
	SelfSubjectAccessReviewExpansion
}

// selfSubjectAccessReviews implements SelfSubjectAccessReviewInterface
type selfSubjectAccessReviews struct {
	client rest.Interface
}

// newSelfSubjectAccessReviews returns a SelfSubjectAccessReviews
func newSelfSubjectAccessReviews(c *ManagementV1Client) *selfSubjectAccessReviews {
	return &selfSubjectAccessReviews{
		client: c.RESTClient(),
	}
}

// Get takes name of the selfSubjectAccessReview, and returns the corresponding selfSubjectAccessReview object, and an error if there is any.
func (c *selfSubjectAccessReviews) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Get().
		Resource("selfsubjectaccessreviews").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SelfSubjectAccessReviews that match those selectors.
func (c *selfSubjectAccessReviews) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SelfSubjectAccessReviewList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SelfSubjectAccessReviewList{}
	err = c.client.Get().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested selfSubjectAccessReviews.
func (c *selfSubjectAccessReviews) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a selfSubjectAccessReview and creates it.  Returns the server's representation of the selfSubjectAccessReview, and an error, if there is any.
func (c *selfSubjectAccessReviews) Create(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.CreateOptions) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Post().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a selfSubjectAccessReview and updates it. Returns the server's representation of the selfSubjectAccessReview, and an error, if there is any.
func (c *selfSubjectAccessReviews) Update(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Put().
		Resource("selfsubjectaccessreviews").
		Name(selfSubjectAccessReview.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *selfSubjectAccessReviews) UpdateStatus(ctx context.Context, selfSubjectAccessReview *v1.SelfSubjectAccessReview, opts metav1.UpdateOptions) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Put().
		Resource("selfsubjectaccessreviews").
		Name(selfSubjectAccessReview.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(selfSubjectAccessReview).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the selfSubjectAccessReview and deletes it. Returns an error if one occurs.
func (c *selfSubjectAccessReviews) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("selfsubjectaccessreviews").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *selfSubjectAccessReviews) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("selfsubjectaccessreviews").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched selfSubjectAccessReview.
func (c *selfSubjectAccessReviews) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SelfSubjectAccessReview, err error) {
	result = &v1.SelfSubjectAccessReview{}
	err = c.client.Patch(pt).
		Resource("selfsubjectaccessreviews").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
