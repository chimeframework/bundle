package listeners

import (
	dispatcher "chime/components/eventdispatcher"
	"chime/components/httpcontext"
	kernel "chime/components/httpkernel"
	"chime/components/routing"
	"fmt"
)

type RouteListener struct {
	dispatcher.EventSubscriber
	router    *routing.Router
	callbacks dispatcher.EventCallbacks
}

func NewRouteListener(router *routing.Router) *RouteListener {
	calls := make(dispatcher.EventCallbacks, 0)
	calls = append(calls, dispatcher.NewEventCallback(kernel.KERNEL_EVENTS_REQUEST, OnKernelRequest, 255))
	return &RouteListener{router: router, callbacks: calls}
}

func (this *RouteListener) GetSubscribedEvents() dispatcher.EventCallbacks {
	return this.callbacks
}

func OnKernelRequest(receiver dispatcher.EventSubscriber, e dispatcher.Eventer) {
	this := receiver.(*RouteListener)
	event := e.(*kernel.ResponseEvent)
	fmt.Printf("%v\n", this, event)

	request := event.GetRequest()

	// already routed
	if request.HasAttribute(httpcontext.CONTROLLER_PARAM) {
		return
	}

	parameters := this.router.MatchRequest(request)
	request.AddAttributes(parameters)

	// TAIL PENDING
}
