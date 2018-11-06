package controller

import (
	"log"

	"github.com/operator-framework/operator-sdk/pkg/ansible/events"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"
)

type Controller struct {
	logger       log.Logger
	client       kubernetes.Interface
	queue        workqueue.Interface
	informer     cache.SharedInformer
	eventHandler events.EventHandler
}
