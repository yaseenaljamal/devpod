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

// LicensesGetter has a method to return a LicenseInterface.
// A group's client should implement this interface.
type LicensesGetter interface {
	Licenses() LicenseInterface
}

// LicenseInterface has methods to work with License resources.
type LicenseInterface interface {
	Create(ctx context.Context, license *v1.License, opts metav1.CreateOptions) (*v1.License, error)
	Update(ctx context.Context, license *v1.License, opts metav1.UpdateOptions) (*v1.License, error)
	UpdateStatus(ctx context.Context, license *v1.License, opts metav1.UpdateOptions) (*v1.License, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.License, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LicenseList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.License, err error)
	LicenseRequest(ctx context.Context, licenseName string, licenseRequest *v1.LicenseRequest, opts metav1.CreateOptions) (*v1.LicenseRequest, error)

	LicenseExpansion
}

// licenses implements LicenseInterface
type licenses struct {
	client rest.Interface
}

// newLicenses returns a Licenses
func newLicenses(c *ManagementV1Client) *licenses {
	return &licenses{
		client: c.RESTClient(),
	}
}

// Get takes name of the license, and returns the corresponding license object, and an error if there is any.
func (c *licenses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.License, err error) {
	result = &v1.License{}
	err = c.client.Get().
		Resource("licenses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Licenses that match those selectors.
func (c *licenses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LicenseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LicenseList{}
	err = c.client.Get().
		Resource("licenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested licenses.
func (c *licenses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("licenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a license and creates it.  Returns the server's representation of the license, and an error, if there is any.
func (c *licenses) Create(ctx context.Context, license *v1.License, opts metav1.CreateOptions) (result *v1.License, err error) {
	result = &v1.License{}
	err = c.client.Post().
		Resource("licenses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(license).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a license and updates it. Returns the server's representation of the license, and an error, if there is any.
func (c *licenses) Update(ctx context.Context, license *v1.License, opts metav1.UpdateOptions) (result *v1.License, err error) {
	result = &v1.License{}
	err = c.client.Put().
		Resource("licenses").
		Name(license.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(license).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *licenses) UpdateStatus(ctx context.Context, license *v1.License, opts metav1.UpdateOptions) (result *v1.License, err error) {
	result = &v1.License{}
	err = c.client.Put().
		Resource("licenses").
		Name(license.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(license).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the license and deletes it. Returns an error if one occurs.
func (c *licenses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("licenses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *licenses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("licenses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched license.
func (c *licenses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.License, err error) {
	result = &v1.License{}
	err = c.client.Patch(pt).
		Resource("licenses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// LicenseRequest takes the representation of a licenseRequest and creates it.  Returns the server's representation of the licenseRequest, and an error, if there is any.
func (c *licenses) LicenseRequest(ctx context.Context, licenseName string, licenseRequest *v1.LicenseRequest, opts metav1.CreateOptions) (result *v1.LicenseRequest, err error) {
	result = &v1.LicenseRequest{}
	err = c.client.Post().
		Resource("licenses").
		Name(licenseName).
		SubResource("request").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(licenseRequest).
		Do(ctx).
		Into(result)
	return
}
