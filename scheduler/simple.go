package scheduler

import (
	"fmt"

	_type "go-spider/type"
)

type SimpleScheduler struct {
	WorkerNum int
	workerChan chan chan _type.Request
}

type ReadyNotify interface {
	WorkerReady(in chan _type.Request)
}

func NewSimpleScheduler() *SimpleScheduler {
	return &SimpleScheduler{}
}

func (s *SimpleScheduler) CreateWorkerChan(out chan _type.Resources)  {
	s.workerChan=make(chan chan _type.Request, s.WorkerNum)
}

func (s *SimpleScheduler) WorkerReady(in chan _type.Request) {
	s.workerChan <- in
}

func (s *SimpleScheduler) Submit(in chan _type.Request,req _type.Request)  {
	go func() {
		in<-req
	}()
}

func (s *SimpleScheduler) Dispatch(in chan _type.Request) {
	requestArr := make([]_type.Request, 0)
	workerArr := make([]chan _type.Request, 0)

	for {
		if len(requestArr) > 0 && len(workerArr) > 0 {
			workerArr[0] <- requestArr[0]
			workerArr = workerArr[1:]
			requestArr = requestArr[1:]
		}

		select {
		case request := <-in:
			requestArr = append(requestArr, request)
		case workers := <-s.workerChan:
			workerArr = append(workerArr, workers)
		}

		fmt.Println(len(s.workerChan), len(workerArr), len(requestArr))

	}
}
