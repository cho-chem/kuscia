// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DomainRouteLister helps list DomainRoutes.
// All objects returned here must be treated as read-only.
type DomainRouteLister interface {
	// List lists all DomainRoutes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainRoute, err error)
	// DomainRoutes returns an object that can list and get DomainRoutes.
	DomainRoutes(namespace string) DomainRouteNamespaceLister
	DomainRouteListerExpansion
}

// domainRouteLister implements the DomainRouteLister interface.
type domainRouteLister struct {
	indexer cache.Indexer
}

// NewDomainRouteLister returns a new DomainRouteLister.
func NewDomainRouteLister(indexer cache.Indexer) DomainRouteLister {
	return &domainRouteLister{indexer: indexer}
}

// List lists all DomainRoutes in the indexer.
func (s *domainRouteLister) List(selector labels.Selector) (ret []*v1alpha1.DomainRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainRoute))
	})
	return ret, err
}

// DomainRoutes returns an object that can list and get DomainRoutes.
func (s *domainRouteLister) DomainRoutes(namespace string) DomainRouteNamespaceLister {
	return domainRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DomainRouteNamespaceLister helps list and get DomainRoutes.
// All objects returned here must be treated as read-only.
type DomainRouteNamespaceLister interface {
	// List lists all DomainRoutes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.DomainRoute, err error)
	// Get retrieves the DomainRoute from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.DomainRoute, error)
	DomainRouteNamespaceListerExpansion
}

// domainRouteNamespaceLister implements the DomainRouteNamespaceLister
// interface.
type domainRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DomainRoutes in the indexer for a given namespace.
func (s domainRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DomainRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DomainRoute))
	})
	return ret, err
}

// Get retrieves the DomainRoute from the indexer for a given namespace and name.
func (s domainRouteNamespaceLister) Get(name string) (*v1alpha1.DomainRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("domainroute"), name)
	}
	return obj.(*v1alpha1.DomainRoute), nil
}
