/*
MIT License

Copyright (c) 2018 PodKubervisor

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by informer-gen. DO NOT EDIT.

// This file was automatically generated by informer-gen

package v1

import (
	time "time"

	kubervisor_v1 "github.com/amadeusitgroup/podkubervisor/pkg/api/kubervisor/v1"
	versioned "github.com/amadeusitgroup/podkubervisor/pkg/client/clientset/versioned"
	internalinterfaces "github.com/amadeusitgroup/podkubervisor/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/amadeusitgroup/podkubervisor/pkg/client/listers/kubervisor/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubervisorServiceInformer provides access to a shared informer and lister for
// KubervisorServices.
type KubervisorServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubervisorServiceLister
}

type kubervisorServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubervisorServiceInformer constructs a new informer for KubervisorService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubervisorServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubervisorServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubervisorServiceInformer constructs a new informer for KubervisorService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubervisorServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BreakerV1().KubervisorServices(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.BreakerV1().KubervisorServices(namespace).Watch(options)
			},
		},
		&kubervisor_v1.KubervisorService{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubervisorServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubervisorServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubervisorServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubervisor_v1.KubervisorService{}, f.defaultInformer)
}

func (f *kubervisorServiceInformer) Lister() v1.KubervisorServiceLister {
	return v1.NewKubervisorServiceLister(f.Informer().GetIndexer())
}