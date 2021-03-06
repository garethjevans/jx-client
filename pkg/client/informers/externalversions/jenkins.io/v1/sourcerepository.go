// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	time "time"

	jenkinsiov1 "github.com/jenkins-x/jx-api/pkg/apis/jenkins.io/v1"
	versioned "github.com/jenkins-x/jx-api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/jenkins-x/jx-api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/jenkins-x/jx-api/pkg/client/listers/jenkins.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SourceRepositoryInformer provides access to a shared informer and lister for
// SourceRepositories.
type SourceRepositoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SourceRepositoryLister
}

type sourceRepositoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSourceRepositoryInformer constructs a new informer for SourceRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSourceRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSourceRepositoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSourceRepositoryInformer constructs a new informer for SourceRepository type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSourceRepositoryInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().SourceRepositories(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.JenkinsV1().SourceRepositories(namespace).Watch(options)
			},
		},
		&jenkinsiov1.SourceRepository{},
		resyncPeriod,
		indexers,
	)
}

func (f *sourceRepositoryInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSourceRepositoryInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sourceRepositoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&jenkinsiov1.SourceRepository{}, f.defaultInformer)
}

func (f *sourceRepositoryInformer) Lister() v1.SourceRepositoryLister {
	return v1.NewSourceRepositoryLister(f.Informer().GetIndexer())
}
