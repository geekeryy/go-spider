package engine

import (
	"fmt"

	"go-spider/scheduler"
	_type "go-spider/type"
	"go-spider/worker"
)

type Engine struct {
	Scans     []_type.Request
	WorkerNum int
	Scheduler scheduler.Scheduler
}

func (e *Engine) Run() {
	in := make(chan _type.Request)
	out := make(chan _type.Resources)

	for _, v := range e.Scans {
		go e.Scheduler.Submit(in, v)
	}

	e.Scheduler.CreateWorkerChan(out)

	for i := 0; i < e.WorkerNum; i++ {
		workerIn := make(chan _type.Request)
		go worker.New().Work(workerIn, out, e.Scheduler)
	}

	go e.Scheduler.Dispatch(in)

	for {
		resources := <-out
		for _, v := range resources.Requests {
			e.Scheduler.Submit(in, v)
		}

		for k, v := range resources.Datas {
			// 数据处理
			fmt.Println(k, v)
		}

	}
}
