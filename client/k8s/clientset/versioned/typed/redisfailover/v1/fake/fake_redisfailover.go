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
	redisfailover_v1 "github.com/jesstracy/redis-operator/api/redisfailover/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRedisFailovers implements RedisFailoverInterface
type FakeRedisFailovers struct {
	Fake *FakeDatabasesV1
	ns   string
}

var redisfailoversResource = schema.GroupVersionResource{Group: "databases.spotahome.com", Version: "v1", Resource: "redisfailovers"}

var redisfailoversKind = schema.GroupVersionKind{Group: "databases.spotahome.com", Version: "v1", Kind: "RedisFailover"}

// Get takes name of the redisFailover, and returns the corresponding redisFailover object, and an error if there is any.
func (c *FakeRedisFailovers) Get(name string, options v1.GetOptions) (result *redisfailover_v1.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(redisfailoversResource, c.ns, name), &redisfailover_v1.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*redisfailover_v1.RedisFailover), err
}

// List takes label and field selectors, and returns the list of RedisFailovers that match those selectors.
func (c *FakeRedisFailovers) List(opts v1.ListOptions) (result *redisfailover_v1.RedisFailoverList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(redisfailoversResource, redisfailoversKind, c.ns, opts), &redisfailover_v1.RedisFailoverList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &redisfailover_v1.RedisFailoverList{ListMeta: obj.(*redisfailover_v1.RedisFailoverList).ListMeta}
	for _, item := range obj.(*redisfailover_v1.RedisFailoverList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested redisFailovers.
func (c *FakeRedisFailovers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(redisfailoversResource, c.ns, opts))

}

// Create takes the representation of a redisFailover and creates it.  Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *FakeRedisFailovers) Create(redisFailover *redisfailover_v1.RedisFailover) (result *redisfailover_v1.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(redisfailoversResource, c.ns, redisFailover), &redisfailover_v1.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*redisfailover_v1.RedisFailover), err
}

// Update takes the representation of a redisFailover and updates it. Returns the server's representation of the redisFailover, and an error, if there is any.
func (c *FakeRedisFailovers) Update(redisFailover *redisfailover_v1.RedisFailover) (result *redisfailover_v1.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(redisfailoversResource, c.ns, redisFailover), &redisfailover_v1.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*redisfailover_v1.RedisFailover), err
}

// Delete takes name of the redisFailover and deletes it. Returns an error if one occurs.
func (c *FakeRedisFailovers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(redisfailoversResource, c.ns, name), &redisfailover_v1.RedisFailover{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRedisFailovers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(redisfailoversResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &redisfailover_v1.RedisFailoverList{})
	return err
}

// Patch applies the patch and returns the patched redisFailover.
func (c *FakeRedisFailovers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *redisfailover_v1.RedisFailover, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(redisfailoversResource, c.ns, name, pt, data, subresources...), &redisfailover_v1.RedisFailover{})

	if obj == nil {
		return nil, err
	}
	return obj.(*redisfailover_v1.RedisFailover), err
}
