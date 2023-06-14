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

package clusterdomainroute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"

	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	kusciafake "github.com/secretflow/kuscia/pkg/crd/clientset/versioned/fake"
	informers "github.com/secretflow/kuscia/pkg/crd/informers/externalversions"
)

func Test_compareSpec(t *testing.T) {
	testcdr := &kusciaapisv1alpha1.ClusterDomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.ClusterDomainRouteSpec{
			DomainRouteSpec: kusciaapisv1alpha1.DomainRouteSpec{
				AuthenticationType: "Token",
			},
		},
	}
	testdr := &kusciaapisv1alpha1.DomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.DomainRouteSpec{
			AuthenticationType: "Token",
		},
	}
	assert.True(t, compareSpec(testcdr, testdr))
	testcdr.Labels["auth2"] = "test"
	assert.False(t, compareSpec(testcdr, testdr))
	delete(testcdr.Labels, "auth2")
	assert.True(t, compareSpec(testcdr, testdr))
	testcdr.Spec.Destination = "alice"
	assert.False(t, compareSpec(testcdr, testdr))
	testcdr.Spec.Destination = ""
	assert.True(t, compareSpec(testcdr, testdr))
}

func Test_compareTokens(t *testing.T) {
	testdr1 := &kusciaapisv1alpha1.DomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.DomainRouteSpec{
			AuthenticationType: "Token",
		},
		Status: kusciaapisv1alpha1.DomainRouteStatus{
			TokenStatus: kusciaapisv1alpha1.DomainRouteTokenStatus{
				Tokens: []kusciaapisv1alpha1.DomainRouteToken{
					{Revision: 1},
					{Revision: 2},
				},
			},
		},
	}
	testdr2 := &kusciaapisv1alpha1.DomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.DomainRouteSpec{
			AuthenticationType: "Token",
		},
		Status: kusciaapisv1alpha1.DomainRouteStatus{
			TokenStatus: kusciaapisv1alpha1.DomainRouteTokenStatus{
				Tokens: []kusciaapisv1alpha1.DomainRouteToken{
					{Revision: 1},
					{Revision: 2},
				},
			},
		},
	}
	assert.True(t, compareTokens(testdr1.Status.TokenStatus.Tokens, testdr2.Status.TokenStatus.Tokens))
	testdr2.Status.TokenStatus.Tokens[1].Revision = 3
	assert.False(t, compareTokens(testdr1.Status.TokenStatus.Tokens, testdr2.Status.TokenStatus.Tokens))
	testdr2.Status.TokenStatus.Tokens[1].Revision = 2
	assert.True(t, compareTokens(testdr1.Status.TokenStatus.Tokens, testdr2.Status.TokenStatus.Tokens))
	testdr2.Status.TokenStatus.Tokens = append(testdr2.Status.TokenStatus.Tokens, kusciaapisv1alpha1.DomainRouteToken{Revision: 3})
	assert.False(t, compareTokens(testdr1.Status.TokenStatus.Tokens, testdr2.Status.TokenStatus.Tokens))
}

func Test_checkReadyCondition(t *testing.T) {
	testcdr := &kusciaapisv1alpha1.ClusterDomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.ClusterDomainRouteSpec{
			DomainRouteSpec: kusciaapisv1alpha1.DomainRouteSpec{
				AuthenticationType: "Token",
			},
		},
		Status: kusciaapisv1alpha1.ClusterDomainRouteStatus{
			Conditions: []kusciaapisv1alpha1.ClusterDomainRouteCondition{},
		},
	}
	assert.False(t, checkReadyCondition(testcdr))
	testcdr.Status.Conditions = append(testcdr.Status.Conditions, kusciaapisv1alpha1.ClusterDomainRouteCondition{Type: kusciaapisv1alpha1.ClusterDomainRouteReady, Status: corev1.ConditionTrue})
	assert.True(t, checkReadyCondition(testcdr))
}

func Test_checkEffectiveInstances(t *testing.T) {
	testNamespace := "test"
	testdr := &kusciaapisv1alpha1.DomainRoute{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testdr",
			Namespace: testNamespace,
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Spec: kusciaapisv1alpha1.DomainRouteSpec{
			AuthenticationType: "Token",
		},
		Status: kusciaapisv1alpha1.DomainRouteStatus{
			TokenStatus: kusciaapisv1alpha1.DomainRouteTokenStatus{},
		},
	}
	testgateway := &kusciaapisv1alpha1.Gateway{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testgw",
			Namespace: testNamespace,
			Labels: map[string]string{
				"auth": "test",
			},
		},
		Status: kusciaapisv1alpha1.GatewayStatus{
			HeartbeatTime: metav1.Time{
				Time: time.Now(),
			},
		},
	}
	kusciaClient := kusciafake.NewSimpleClientset()
	kusciaInformerFactory := informers.NewSharedInformerFactory(kusciaClient, time.Second*30)
	gatewayInformer := kusciaInformerFactory.Kuscia().V1alpha1().Gateways()
	controller := &Controller{
		gatewayLister:       gatewayInformer.Lister(),
		gatewayListerSynced: gatewayInformer.Informer().HasSynced,
	}
	kusciaInformerFactory.Start(context.Background().Done())
	cache.WaitForCacheSync(context.Background().Done(), controller.gatewayListerSynced)
	assert.False(t, controller.checkEffectiveInstances(testdr))
	testdr.Status.TokenStatus.Tokens = []kusciaapisv1alpha1.DomainRouteToken{
		{Revision: 1},
		{Revision: 2},
	}
	assert.False(t, controller.checkEffectiveInstances(testdr))
	kusciaClient.KusciaV1alpha1().Gateways(testNamespace).Create(context.Background(), testgateway, metav1.CreateOptions{})
	time.Sleep(time.Millisecond)
	assert.False(t, controller.checkEffectiveInstances(testdr))
	testdr.Status.TokenStatus.Tokens[1].EffectiveInstances = []string{"testgw"}
	assert.True(t, controller.checkEffectiveInstances(testdr))
}
