package engine

import (
	"go-spider/duplicate"
	"go-spider/scheduler"
	_type "go-spider/type"
	"go-spider/worker"
)

type Engine struct {
	Scans     []_type.Request
	WorkerNum int
	Scheduler scheduler.Scheduler
	DefaultStorage _type.Storage
	RandomWaiting int
	DuplicateType string
}



func (e *Engine) Run() {
	in := make(chan _type.Request)
	out := make(chan _type.Resources)

	for _, v := range e.Scans {
		e.Scheduler.Submit(in, v)
	}

	e.Scheduler.CreateWorkerChan(out)

	for i := 0; i < e.WorkerNum; i++ {
		workerIn := make(chan _type.Request)
		go worker.Worker{
			RandomWaiting: e.RandomWaiting,
		}.Work(workerIn, out, e.Scheduler)
	}

	go e.Scheduler.Dispatch(in)

	for {
		resources := <-out
		for _, v := range resources.Requests {
			if !duplicate.IsDuplicate(e.DuplicateType,v.Url) {
				e.Scheduler.Submit(in, v)
			}
		}

		for _, v := range resources.Datas {
			// 数据处理
			if v.Storage != nil {
				v.Storage.Store(v)
			} else {
				e.DefaultStorage.Store(v)
			}
		}

	}
}
