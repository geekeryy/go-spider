package worker

import (
	"log"
	"math/rand"
	"time"

	"go-spider/download"
	"go-spider/scheduler"
	_type "go-spider/type"
)

type Worker struct {
}

func New() *Worker {
	return &Worker{}
}

func (w *Worker) Work(in chan _type.Request, out chan _type.Resources, ready scheduler.ReadyNotify) {
	c := download.NewClient()
	ticker := time.NewTicker(time.Duration(rand.Intn(5)) * time.Second)
	for {
		ready.WorkerReady(in)
		request := <-in
		<-ticker.C
		body, err := download.Get(c, request.Url)
		if err != nil {
			log.Println("Work err:", err, request.Url)
			continue
		}
		resources := request.Fetcher(body)
		out <- resources
	}
}
