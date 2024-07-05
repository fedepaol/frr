// SPDX-License-Identifier:Apache-2.0

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	apiv1beta1 "github.com/metallb/frr-k8s/api/v1beta1"
	versioned "github.com/metallb/frr-k8s/pkg/client/clientset/versioned"
	internalinterfaces "github.com/metallb/frr-k8s/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/metallb/frr-k8s/pkg/client/listers/core/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FRRConfigurationInformer provides access to a shared informer and lister for
// FRRConfigurations.
type FRRConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FRRConfigurationLister
}

type fRRConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFRRConfigurationInformer constructs a new informer for FRRConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFRRConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFRRConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFRRConfigurationInformer constructs a new informer for FRRConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFRRConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().FRRConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1beta1().FRRConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&apiv1beta1.FRRConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *fRRConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFRRConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fRRConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiv1beta1.FRRConfiguration{}, f.defaultInformer)
}

func (f *fRRConfigurationInformer) Lister() v1beta1.FRRConfigurationLister {
	return v1beta1.NewFRRConfigurationLister(f.Informer().GetIndexer())
}
