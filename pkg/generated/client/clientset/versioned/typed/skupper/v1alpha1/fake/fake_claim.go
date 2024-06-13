/*
Copyright 2021 The Skupper Authors.

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

package fake

import (
	"context"

	v1alpha1 "github.com/skupperproject/skupper/pkg/apis/skupper/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClaims implements ClaimInterface
type FakeClaims struct {
	Fake *FakeSkupperV1alpha1
	ns   string
}

var claimsResource = schema.GroupVersionResource{Group: "skupper.io", Version: "v1alpha1", Resource: "claims"}

var claimsKind = schema.GroupVersionKind{Group: "skupper.io", Version: "v1alpha1", Kind: "Claim"}

// Get takes name of the claim, and returns the corresponding claim object, and an error if there is any.
func (c *FakeClaims) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Claim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(claimsResource, c.ns, name), &v1alpha1.Claim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Claim), err
}

// List takes label and field selectors, and returns the list of Claims that match those selectors.
func (c *FakeClaims) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClaimList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(claimsResource, claimsKind, c.ns, opts), &v1alpha1.ClaimList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClaimList{ListMeta: obj.(*v1alpha1.ClaimList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClaimList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested claims.
func (c *FakeClaims) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(claimsResource, c.ns, opts))

}

// Create takes the representation of a claim and creates it.  Returns the server's representation of the claim, and an error, if there is any.
func (c *FakeClaims) Create(ctx context.Context, claim *v1alpha1.Claim, opts v1.CreateOptions) (result *v1alpha1.Claim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(claimsResource, c.ns, claim), &v1alpha1.Claim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Claim), err
}

// Update takes the representation of a claim and updates it. Returns the server's representation of the claim, and an error, if there is any.
func (c *FakeClaims) Update(ctx context.Context, claim *v1alpha1.Claim, opts v1.UpdateOptions) (result *v1alpha1.Claim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(claimsResource, c.ns, claim), &v1alpha1.Claim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Claim), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClaims) UpdateStatus(ctx context.Context, claim *v1alpha1.Claim, opts v1.UpdateOptions) (*v1alpha1.Claim, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(claimsResource, "status", c.ns, claim), &v1alpha1.Claim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Claim), err
}

// Delete takes name of the claim and deletes it. Returns an error if one occurs.
func (c *FakeClaims) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(claimsResource, c.ns, name), &v1alpha1.Claim{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClaims) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(claimsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClaimList{})
	return err
}

// Patch applies the patch and returns the patched claim.
func (c *FakeClaims) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Claim, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(claimsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Claim{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Claim), err
}