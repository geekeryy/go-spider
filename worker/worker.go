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
	MaxLevel      int
	RandomWaiting int
}

func (w Worker) Work(in chan _type.Request, out chan _type.Resources, ready scheduler.ReadyNotify) {
	c := download.NewClient()
	ticker := time.NewTicker(time.Duration(rand.Intn(w.RandomWaiting)) * time.Second)
	for {
		ready.WorkerReady(in)

		request := <-in

		<-ticker.C

		if request.Level >= request.MaxLevel {
			log.Println("Max Level:", request.Url)
			continue
		}

		log.Println("GET:", request.Url, request.Level)

		body, err := download.Get(c, request.Url)
		if err != nil {
			log.Println("Work err:", err, request.Url)
			continue
		}
		resources := request.Fetcher(body, request)
		out <- resources
	}
}
