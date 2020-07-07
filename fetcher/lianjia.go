package fetcher

import (
	"io"
	"log"

	"github.com/PuerkitoBio/goquery"

	"go-spider/storage"
	_type "go-spider/type"
)





func Fetch(body io.ReadCloser,req _type.Request) (resources _type.Resources) {
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Fetch err: ", err)
		return
	}
	resources = _type.Resources{}
	resources.Requests = make([]_type.Request, 0)
	resources.Datas = make([]_type.Data, 0)

	doc.Find("[data-role='ershoufang'] [href]").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     "https://cd.lianjia.com" + href,
				Fetcher: Fetch,
				Level: req.Level+1,
				MaxLevel: req.MaxLevel,
			})
		}
	})

	doc.Find(".house-lst-page-box  [href]").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     "https://cd.lianjia.com" + href,
				Fetcher: Fetch,
				Level: req.Level+1,
				MaxLevel: req.MaxLevel,
			})
		}
	})

	doc.Find(".listContent   li > a ").Each(func(i int, selection *goquery.Selection) {
		if href, ok := selection.Attr("href"); ok {
			resources.Requests = append(resources.Requests, _type.Request{
				Url:     href,
				Fetcher: FetchInfo,
				Level: req.Level+1,
				MaxLevel: req.MaxLevel,
			})
		}
	})

	doc.Find(".listContent   li ").Each(func(i int, selection *goquery.Selection) {
		resources.Datas = append(resources.Datas, _type.Data{
			Type:    "name",
			Content: selection.Find(".info > .title > a").Text()+"---"+selection.Find(".xiaoquListItemRight > .xiaoquListItemPrice > .totalPrice > span").Text(),
		})
	})

	return
}

func FetchInfo(body io.ReadCloser,req _type.Request) (resources _type.Resources) {
	defer body.Close()
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Println("Fetch err: ", err)
		return
	}
	resources = _type.Resources{}
	resources.Requests = make([]_type.Request, 0)
	resources.Datas = make([]_type.Data, 0)

	resources.Datas = append(resources.Datas, _type.Data{
		Storage: storage.NewFileStorage(),
		Type:    "info",
		Content: doc.Find(".xiaoquInfo").Text(),
	})

	return resources
}
