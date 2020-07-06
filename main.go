package main

import (
	"go-spider/engine"
	"go-spider/fetcher"
	"go-spider/scheduler"
	_type "go-spider/type"
)


func main() {
	scans:=make([]_type.Request,0)
	scans= append(scans, _type.Request{
		Url:     "https://cd.lianjia.com/xiaoqu/",
		Fetcher: fetcher.Fetch,
	})
	e:=engine.Engine{
		Scans:     scans,
		WorkerNum: 5,
		Scheduler: scheduler.NewSimpleScheduler(),
	}
	e.Run()
}