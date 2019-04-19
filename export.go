package texts

import (
	"strings"
)

func UUIDString() string {
	return strings.Replace(UUIDRawString(), "-", "", -1)
}

func UUIDRawString() string {
	return hObj.UUID().String()
}

func Md5(data interface{}) string {
	return hObj.Md5(data)
}

func UrlEncode(data interface{}) (string, error) {
	return hObj.UrlEncode(data)
}
