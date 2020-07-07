package duplicate

import (
	_type "go-spider/type"
)

var strMap map[string]struct{}

func IsDuplicate(duplicateType string,url string) bool {
	switch duplicateType {
	case _type.DuplicateType_Map:
		return duplicateMap(url)
	}
	return false
}

// 内存去重
func duplicateMap(url string) bool {
	if _,ok:=strMap[url];ok{
		return true
	}
	strMap[url]= struct{}{}
	return false
}
