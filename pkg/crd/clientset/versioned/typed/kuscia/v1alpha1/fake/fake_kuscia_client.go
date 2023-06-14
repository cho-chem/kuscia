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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/secretflow/kuscia/pkg/crd/clientset/versioned/typed/kuscia/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKusciaV1alpha1 struct {
	*testing.Fake
}

func (c *FakeKusciaV1alpha1) AppImages() v1alpha1.AppImageInterface {
	return &FakeAppImages{c}
}

func (c *FakeKusciaV1alpha1) ClusterDomainRoutes() v1alpha1.ClusterDomainRouteInterface {
	return &FakeClusterDomainRoutes{c}
}

func (c *FakeKusciaV1alpha1) DataObjects(namespace string) v1alpha1.DataObjectInterface {
	return &FakeDataObjects{c, namespace}
}

func (c *FakeKusciaV1alpha1) DataSources(namespace string) v1alpha1.DataSourceInterface {
	return &FakeDataSources{c, namespace}
}

func (c *FakeKusciaV1alpha1) DataTables(namespace string) v1alpha1.DataTableInterface {
	return &FakeDataTables{c, namespace}
}

func (c *FakeKusciaV1alpha1) Domains() v1alpha1.DomainInterface {
	return &FakeDomains{c}
}

func (c *FakeKusciaV1alpha1) DomainAppImages(namespace string) v1alpha1.DomainAppImageInterface {
	return &FakeDomainAppImages{c, namespace}
}

func (c *FakeKusciaV1alpha1) DomainRoutes(namespace string) v1alpha1.DomainRouteInterface {
	return &FakeDomainRoutes{c, namespace}
}

func (c *FakeKusciaV1alpha1) Gateways(namespace string) v1alpha1.GatewayInterface {
	return &FakeGateways{c, namespace}
}

func (c *FakeKusciaV1alpha1) InteropConfigs() v1alpha1.InteropConfigInterface {
	return &FakeInteropConfigs{c}
}

func (c *FakeKusciaV1alpha1) KusciaJobs() v1alpha1.KusciaJobInterface {
	return &FakeKusciaJobs{c}
}

func (c *FakeKusciaV1alpha1) KusciaTasks() v1alpha1.KusciaTaskInterface {
	return &FakeKusciaTasks{c}
}

func (c *FakeKusciaV1alpha1) TaskResources(namespace string) v1alpha1.TaskResourceInterface {
	return &FakeTaskResources{c, namespace}
}

func (c *FakeKusciaV1alpha1) TaskResourceGroups() v1alpha1.TaskResourceGroupInterface {
	return &FakeTaskResourceGroups{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKusciaV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
