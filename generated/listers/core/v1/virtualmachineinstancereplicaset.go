/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2023 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "kubevirt.io/api/core/v1"
)

// VirtualMachineInstanceReplicaSetLister helps list VirtualMachineInstanceReplicaSets.
// All objects returned here must be treated as read-only.
type VirtualMachineInstanceReplicaSetLister interface {
	// List lists all VirtualMachineInstanceReplicaSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualMachineInstanceReplicaSet, err error)
	// VirtualMachineInstanceReplicaSets returns an object that can list and get VirtualMachineInstanceReplicaSets.
	VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetNamespaceLister
	VirtualMachineInstanceReplicaSetListerExpansion
}

// virtualMachineInstanceReplicaSetLister implements the VirtualMachineInstanceReplicaSetLister interface.
type virtualMachineInstanceReplicaSetLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineInstanceReplicaSetLister returns a new VirtualMachineInstanceReplicaSetLister.
func NewVirtualMachineInstanceReplicaSetLister(indexer cache.Indexer) VirtualMachineInstanceReplicaSetLister {
	return &virtualMachineInstanceReplicaSetLister{indexer: indexer}
}

// List lists all VirtualMachineInstanceReplicaSets in the indexer.
func (s *virtualMachineInstanceReplicaSetLister) List(selector labels.Selector) (ret []*v1.VirtualMachineInstanceReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachineInstanceReplicaSet))
	})
	return ret, err
}

// VirtualMachineInstanceReplicaSets returns an object that can list and get VirtualMachineInstanceReplicaSets.
func (s *virtualMachineInstanceReplicaSetLister) VirtualMachineInstanceReplicaSets(namespace string) VirtualMachineInstanceReplicaSetNamespaceLister {
	return virtualMachineInstanceReplicaSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineInstanceReplicaSetNamespaceLister helps list and get VirtualMachineInstanceReplicaSets.
// All objects returned here must be treated as read-only.
type VirtualMachineInstanceReplicaSetNamespaceLister interface {
	// List lists all VirtualMachineInstanceReplicaSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.VirtualMachineInstanceReplicaSet, err error)
	// Get retrieves the VirtualMachineInstanceReplicaSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.VirtualMachineInstanceReplicaSet, error)
	VirtualMachineInstanceReplicaSetNamespaceListerExpansion
}

// virtualMachineInstanceReplicaSetNamespaceLister implements the VirtualMachineInstanceReplicaSetNamespaceLister
// interface.
type virtualMachineInstanceReplicaSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineInstanceReplicaSets in the indexer for a given namespace.
func (s virtualMachineInstanceReplicaSetNamespaceLister) List(selector labels.Selector) (ret []*v1.VirtualMachineInstanceReplicaSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.VirtualMachineInstanceReplicaSet))
	})
	return ret, err
}

// Get retrieves the VirtualMachineInstanceReplicaSet from the indexer for a given namespace and name.
func (s virtualMachineInstanceReplicaSetNamespaceLister) Get(name string) (*v1.VirtualMachineInstanceReplicaSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("virtualmachineinstancereplicaset"), name)
	}
	return obj.(*v1.VirtualMachineInstanceReplicaSet), nil
}
