package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/PuerkitoBio/goquery"
)

type Resources struct {
	Requests []Request
	Datas    []Data
}

type Data struct {
	Name       string
	TotalPrice int
}

type Request struct {
	Url     string
	Fetcher func(body io.ReadCloser) (resources *Resources)
}

func main() {
	// 下载页面
	client := NewClient()
	data, err := Get(client, "https://cd.lianjia.com/xiaoqu/")
	fmt.Println(err)
	
	// 取出数据
	fetch := Fetch(data)
	fmt.Println(fetch.Requests)


	// 存储数据
}

type Engine struct {
	Scans []string
	WorkerNum int
}

func (e *Engine) Run() {
	for  {

	}
}

type Worker struct {
	Client *http.Client
	Url string
	Resources
}

func NewWorker(in chan string,out chan Resources)  {

}

func (w *Worker) Work() {

}



func NewClient() *http.Client {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil
	}
	client := http.DefaultClient
	client.Jar = jar
	return client
}

func Get(client *http.Client, url string) (io.ReadCloser, error) {
	return download(client,"GET",url)
}

func download(client *http.Client, method string, url string) (io.ReadCloser, error) {
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.115 Safari/537.36")

	res, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return res.Body, nil

}

func Fetch(body io.ReadCloser) (resources *Resources) {
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Fetch err: ", err)
		return
	}
	resources=&Resources{}
	resources.Requests=make([]Request,0)

	doc.Find("[data-role='ershoufang'] [href]").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, Request{
				Url:     href,
				Fetcher: Fetch,
			})
		}
	})

	doc.Find(".page-box  [href]").Each(func(i int, selection *goquery.Selection) {
		if href,ok:=selection.Attr("href");ok{
			resources.Requests=append(resources.Requests,Request{
				Url: href,
				Fetcher: Fetch,
			})
		}
	})

	return
}
