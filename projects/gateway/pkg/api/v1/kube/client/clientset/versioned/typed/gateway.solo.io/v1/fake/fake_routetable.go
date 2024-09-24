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

package fake

import (
	"context"

	v1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1/kube/apis/gateway.solo.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRouteTables implements RouteTableInterface
type FakeRouteTables struct {
	Fake *FakeGatewayV1
	ns   string
}

var routetablesResource = v1.SchemeGroupVersion.WithResource("routetables")

var routetablesKind = v1.SchemeGroupVersion.WithKind("RouteTable")

// Get takes name of the routeTable, and returns the corresponding routeTable object, and an error if there is any.
func (c *FakeRouteTables) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RouteTable, err error) {
	emptyResult := &v1.RouteTable{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(routetablesResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RouteTable), err
}

// List takes label and field selectors, and returns the list of RouteTables that match those selectors.
func (c *FakeRouteTables) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RouteTableList, err error) {
	emptyResult := &v1.RouteTableList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(routetablesResource, routetablesKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RouteTableList{ListMeta: obj.(*v1.RouteTableList).ListMeta}
	for _, item := range obj.(*v1.RouteTableList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested routeTables.
func (c *FakeRouteTables) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(routetablesResource, c.ns, opts))

}

// Create takes the representation of a routeTable and creates it.  Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Create(ctx context.Context, routeTable *v1.RouteTable, opts metav1.CreateOptions) (result *v1.RouteTable, err error) {
	emptyResult := &v1.RouteTable{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(routetablesResource, c.ns, routeTable, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RouteTable), err
}

// Update takes the representation of a routeTable and updates it. Returns the server's representation of the routeTable, and an error, if there is any.
func (c *FakeRouteTables) Update(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (result *v1.RouteTable, err error) {
	emptyResult := &v1.RouteTable{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(routetablesResource, c.ns, routeTable, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RouteTable), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRouteTables) UpdateStatus(ctx context.Context, routeTable *v1.RouteTable, opts metav1.UpdateOptions) (result *v1.RouteTable, err error) {
	emptyResult := &v1.RouteTable{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(routetablesResource, "status", c.ns, routeTable, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RouteTable), err
}

// Delete takes name of the routeTable and deletes it. Returns an error if one occurs.
func (c *FakeRouteTables) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(routetablesResource, c.ns, name, opts), &v1.RouteTable{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRouteTables) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(routetablesResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RouteTableList{})
	return err
}

// Patch applies the patch and returns the patched routeTable.
func (c *FakeRouteTables) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RouteTable, err error) {
	emptyResult := &v1.RouteTable{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(routetablesResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.RouteTable), err
}
