package _type

import (
	"io"
)

const (
	DuplicateType_Map = "map"
)

type Resources struct {
	Requests []Request
	Datas    []Data
}

type Data struct {
	Storage Storage
	Type    string
	Content interface{}
}

type Request struct {
	Url     string
	Level   int
	MaxLevel   int
	Fetcher func (body io.ReadCloser,req Request) (resources Resources)
}

type Storage interface {
	Store(v interface{})
}

