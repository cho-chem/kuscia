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

package election

import (
	"context"

	"go.uber.org/atomic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"

	"github.com/secretflow/kuscia/pkg/utils/nlog"
)

type Elector interface {
	IsLeader() bool
	GetLeader() string
	MyIdentity() string
	Run(ctx context.Context)
	Elect(ctx context.Context)
	Stop()
	Stopped() <-chan struct{}
}

func NewElector(kubeClient kubernetes.Interface, name string, opt ...Option) Elector {
	options := defaultOptions()
	options.KubeClient = kubeClient
	options.Name = name
	for _, o := range opt {
		o(options)
	}

	if err := checkOption(options); err != nil {
		nlog.Fatal(err)
		return nil
	}

	e := k8sElector{
		options:   options,
		chStopped: make(chan struct{}),
	}

	leaderElector, err := buildLeaderElector(options, e.onStartedLeading, e.onStoppedLeading, e.onNewLeader)
	if err != nil {
		nlog.Fatal(err)
		return nil
	}

	e.leaderElector = leaderElector

	return &e
}

func buildLeaderElector(options *Options, onStartedLeading func(context.Context), onStoppedLeading func(), onNewLeader func(string)) (*leaderelection.LeaderElector, error) {
	rl, err := resourcelock.New(
		resourcelock.LeasesResourceLock,
		options.Namespace,
		options.Name,
		options.KubeClient.CoreV1(),
		options.KubeClient.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity:      options.Identity,
			EventRecorder: options.EventRecorder,
		},
	)
	if err != nil {
		return nil, err
	}

	lec := leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: options.LeaseDuration,
		RenewDeadline: options.RenewDuration,
		RetryPeriod:   options.RetryPeriod,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: onStartedLeading,
			OnStoppedLeading: onStoppedLeading,
			OnNewLeader:      onNewLeader,
		},
		WatchDog: options.HealthChecker,
		Name:     options.Name,
	}

	leaderElector, err := leaderelection.NewLeaderElector(lec)
	if err != nil {
		return nil, err
	}

	return leaderElector, nil
}

type k8sElector struct {
	options       *Options
	leaderID      atomic.String
	isLeader      atomic.Bool
	leaderElector *leaderelection.LeaderElector
	electionCtx   context.Context
	cancel        context.CancelFunc
	stop          atomic.Bool
	chStopped     chan struct{}
}

func (e *k8sElector) IsLeader() bool {
	return e.isLeader.Load()
}

func (e *k8sElector) GetLeader() string {
	return e.leaderID.Load()
}

func (e *k8sElector) MyIdentity() string {
	return e.options.Identity
}

// Elect leader and wait for leader
func (e *k8sElector) Elect(ctx context.Context) {
	e.electionCtx, e.cancel = context.WithCancel(ctx)
	go e.leaderElector.Run(e.electionCtx)
}

func (e *k8sElector) Stop() {
	e.stop.Store(true)
	e.cancel()
}

func (e *k8sElector) Stopped() <-chan struct{} {
	return e.chStopped
}

func (e *k8sElector) onNewLeader(identity string) {
	e.leaderID.Store(identity)
	if e.options.OnNewLeader != nil {
		e.options.OnNewLeader(identity)
	}
}

func (e *k8sElector) onStartedLeading(ctx context.Context) {
	e.isLeader.Store(true)
	if e.options.OnStartedLeading != nil {
		e.options.OnStartedLeading(ctx)
	}
}

func (e *k8sElector) onStoppedLeading() {
	e.isLeader.Store(false)
	if e.options.OnStoppedLeading != nil {
		e.options.OnStoppedLeading()
	}

	if e.stop.Load() {
		close(e.chStopped)
		return
	}
	e.Elect(e.electionCtx)
}

func (e *k8sElector) Run(ctx context.Context) {
	e.electionCtx, e.cancel = context.WithCancel(ctx)
	e.leaderElector.Run(e.electionCtx)
	<-e.electionCtx.Done()
	e.stop.Store(true)
	<-e.Stopped()
}
