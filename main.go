package main

import (
	"go-spider/engine"
	"go-spider/fetcher"
	"go-spider/scheduler"
	"go-spider/storage"
	_type "go-spider/type"
)


func main() {
	scans:=make([]_type.Request,0)
	scans= append(scans, _type.Request{
		Url:     "https://cd.lianjia.com/xiaoqu/",
		Fetcher: fetcher.Fetch,
		MaxLevel: 3,
	})
	e:=engine.Engine{
		Scans:     scans,
		WorkerNum: 5,
		Scheduler: scheduler.NewSimpleScheduler(),
		DefaultStorage: storage.NewFileStorage(),
		RandomWaiting: 5,
		DuplicateType: _type.DuplicateType_Map,
	}
	e.Run()
}