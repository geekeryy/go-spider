package scheduler

import (
	_type "go-spider/type"
)

type Scheduler interface {
	Submit(in chan _type.Request,req _type.Request)
	CreateWorkerChan(out chan _type.Resources)
	WorkerReady(in chan _type.Request)
	Dispatch(in chan _type.Request)
}
