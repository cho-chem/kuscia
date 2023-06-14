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

package handler

import (
	"k8s.io/client-go/kubernetes"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/record"

	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	kusciaclientset "github.com/secretflow/kuscia/pkg/crd/clientset/versioned"
	kuscialistersv1alpha1 "github.com/secretflow/kuscia/pkg/crd/listers/kuscia/v1alpha1"
)

// Dependencies defines some parameter dependencies of functions.
type Dependencies struct {
	KubeClient      kubernetes.Interface
	KusciaClient    kusciaclientset.Interface
	TrgLister       kuscialistersv1alpha1.TaskResourceGroupLister
	TrLister        kuscialistersv1alpha1.TaskResourceLister
	PodsLister      corelisters.PodLister
	ServicesLister  corelisters.ServiceLister
	ConfigMapLister corelisters.ConfigMapLister
	AppImagesLister kuscialistersv1alpha1.AppImageLister
	Recorder        record.EventRecorder
}

// KusciaTaskPhaseHandler is an interface to handle kuscia task.
type KusciaTaskPhaseHandler interface {
	Handle(kusciaTask *kusciaapisv1alpha1.KusciaTask) (bool, error)
}

// NewKusciaTaskPhaseHandlerFactory returns a KusciaTaskPhaseHandlerFactory instance.
func NewKusciaTaskPhaseHandlerFactory(deps *Dependencies) *KusciaTaskPhaseHandlerFactory {
	finishedHandler := NewFinishedHandler(deps)
	pendingHandler := NewPendingHandler()
	creatingHandler := NewCreatingHandler(deps)
	runningHandler := NewRunningHandler(deps)
	succeededHandler := NewSucceededHandler(deps, finishedHandler)
	failedHandler := NewFailedHandler(deps, finishedHandler)
	KusciaTaskStateHandlerMap := map[kusciaapisv1alpha1.KusciaTaskPhase]KusciaTaskPhaseHandler{
		kusciaapisv1alpha1.TaskPending:   pendingHandler,
		kusciaapisv1alpha1.TaskCreating:  creatingHandler,
		kusciaapisv1alpha1.TaskRunning:   runningHandler,
		kusciaapisv1alpha1.TaskSucceeded: succeededHandler,
		kusciaapisv1alpha1.TaskFailed:    failedHandler,
	}
	return &KusciaTaskPhaseHandlerFactory{KusciaTaskStateHandlerMap: KusciaTaskStateHandlerMap}
}

// KusciaTaskPhaseHandlerFactory is a factory to get phase handler by task resource group phase.
type KusciaTaskPhaseHandlerFactory struct {
	KusciaTaskStateHandlerMap map[kusciaapisv1alpha1.KusciaTaskPhase]KusciaTaskPhaseHandler
}

// GetKusciaTaskPhaseHandler is used to get KusciaTaskPhaseHandler by KusciaTaskPhase.
func (m *KusciaTaskPhaseHandlerFactory) GetKusciaTaskPhaseHandler(conditionType kusciaapisv1alpha1.KusciaTaskPhase) KusciaTaskPhaseHandler {
	return m.KusciaTaskStateHandlerMap[conditionType]
}

func getKusciaTaskCondition(kusciaTaskStatus *kusciaapisv1alpha1.KusciaTaskStatus, conditionType kusciaapisv1alpha1.KusciaTaskConditionType) *kusciaapisv1alpha1.KusciaTaskCondition {
	for i, condition := range kusciaTaskStatus.Conditions {
		if condition.Type == conditionType {
			return &kusciaTaskStatus.Conditions[i]
		}
	}

	kusciaTaskStatus.Conditions = append(kusciaTaskStatus.Conditions, kusciaapisv1alpha1.KusciaTaskCondition{
		Type: conditionType,
	})

	return &kusciaTaskStatus.Conditions[len(kusciaTaskStatus.Conditions)-1]
}
