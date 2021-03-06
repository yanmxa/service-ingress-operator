package pkg

import (
	coreInformer "k8s.io/client-go/informers/core/v1"
	netInformer "k8s.io/client-go/informers/networking/v1"
	"k8s.io/client-go/kubernetes"
	coreLister "k8s.io/client-go/listers/core/v1"
	netLister "k8s.io/client-go/listers/networking/v1"
	"k8s.io/client-go/tools/cache"
)

type controller struct {
	client kubernetes.Interface
	ingressLister netLister.IngressLister
	serviceLister coreLister.ServiceLister
}

func (c *controller) updateService(obj interface{}, newObj interface{}) {

}

func (c *controller) addService(obj interface{}) {

}

func (c *controller) deleteIngress(obj interface{}) {

}

func (c *controller) Run(stopCh chan struct{}) {
	<- stopCh
}

func NewController(client kubernetes.Interface, serviceInformer coreInformer.ServiceInformer, ingressInformer netInformer.IngressInformer) controller {
	c := controller {
		client: client,
		ingressLister: ingressInformer.Lister(),
		serviceLister: serviceInformer.Lister(),
	}
	serviceInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: c.addService,
		UpdateFunc: c.updateService,
	})
	ingressInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: c.deleteIngress,
	})  
	return c
}