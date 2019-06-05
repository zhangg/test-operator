/*
 Copyright 2019 The CloudDB Authors

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
	v1alpha1 "github.com/zhangg/test-operator/pkg/apis/rds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMysqls implements MysqlInterface
type FakeMysqls struct {
	Fake *FakeRdsV1alpha1
	ns   string
}

var mysqlsResource = schema.GroupVersionResource{Group: "rds.github.com", Version: "v1alpha1", Resource: "mysqls"}

var mysqlsKind = schema.GroupVersionKind{Group: "rds.github.com", Version: "v1alpha1", Kind: "Mysql"}

// Get takes name of the mysql, and returns the corresponding mysql object, and an error if there is any.
func (c *FakeMysqls) Get(name string, options v1.GetOptions) (result *v1alpha1.Mysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mysqlsResource, c.ns, name), &v1alpha1.Mysql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mysql), err
}

// List takes label and field selectors, and returns the list of Mysqls that match those selectors.
func (c *FakeMysqls) List(opts v1.ListOptions) (result *v1alpha1.MysqlList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mysqlsResource, mysqlsKind, c.ns, opts), &v1alpha1.MysqlList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MysqlList{ListMeta: obj.(*v1alpha1.MysqlList).ListMeta}
	for _, item := range obj.(*v1alpha1.MysqlList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mysqls.
func (c *FakeMysqls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mysqlsResource, c.ns, opts))

}

// Create takes the representation of a mysql and creates it.  Returns the server's representation of the mysql, and an error, if there is any.
func (c *FakeMysqls) Create(mysql *v1alpha1.Mysql) (result *v1alpha1.Mysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mysqlsResource, c.ns, mysql), &v1alpha1.Mysql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mysql), err
}

// Update takes the representation of a mysql and updates it. Returns the server's representation of the mysql, and an error, if there is any.
func (c *FakeMysqls) Update(mysql *v1alpha1.Mysql) (result *v1alpha1.Mysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mysqlsResource, c.ns, mysql), &v1alpha1.Mysql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mysql), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMysqls) UpdateStatus(mysql *v1alpha1.Mysql) (*v1alpha1.Mysql, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mysqlsResource, "status", c.ns, mysql), &v1alpha1.Mysql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mysql), err
}

// Delete takes name of the mysql and deletes it. Returns an error if one occurs.
func (c *FakeMysqls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mysqlsResource, c.ns, name), &v1alpha1.Mysql{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMysqls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mysqlsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MysqlList{})
	return err
}

// Patch applies the patch and returns the patched mysql.
func (c *FakeMysqls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mysql, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mysqlsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Mysql{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Mysql), err
}
