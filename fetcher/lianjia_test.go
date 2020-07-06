package fetcher

import (
	"fmt"
	"log"
	"testing"

	"go-spider/download"
)

func TestFetch(t *testing.T) {
	body, err := download.Get(download.NewClient(), "https://cd.lianjia.com/xiaoqu")
	if err != nil {
		log.Println("Work err:", err)
	}
	fetch := Fetch(body)

	for k,v:=range fetch.Requests {
		fmt.Println(k,v.Url)
	}

	for k,v:=range fetch.Datas {
		fmt.Println(k,v)
	}

}
