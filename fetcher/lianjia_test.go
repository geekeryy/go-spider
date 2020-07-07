package fetcher

import (
	"fmt"
	"log"
	"testing"

	"go-spider/download"
	_type "go-spider/type"
)

func TestFetch(t *testing.T) {
	body, err := download.Get(download.NewClient(), "https://cd.lianjia.com/xiaoqu")
	if err != nil {
		log.Println("Work err:", err)
	}
	fetch := Fetch(body,_type.Request{
		MaxLevel: 3,
	})

	for k,v:=range fetch.Requests {
		fmt.Println(k,v.Url)
	}

	for k,v:=range fetch.Datas {
		fmt.Println(k,v)
	}

}
