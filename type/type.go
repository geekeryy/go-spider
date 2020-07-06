package _type

import (
	"io"
)

type Resources struct {
	Requests []Request
	Datas    []Data
}

type Data struct {
	Type    string
	Content interface{}
}

type Request struct {
	Url     string
	Fetcher func(body io.ReadCloser) (resources Resources)
}
