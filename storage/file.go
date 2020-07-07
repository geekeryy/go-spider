package storage

import (
	"fmt"
	"log"
	"os"

	_type "go-spider/type"
)

type FileStorage struct {
	fileName string
}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (f *FileStorage) Store(v interface{}) {
	file, err := os.OpenFile("./info.txt",os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err!=nil {
		log.Println("Store err:",err)
		return
	}
	defer file.Close()

	if value, ok := v.(_type.Data);ok {
		switch value.Type {
		case "info":

		default:
			file.WriteString(value.Content.(string)+"\r\n")
		}
	}
	fmt.Println(v)
}
