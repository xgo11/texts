package texts

import (
	"crypto/md5"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

import (
	"github.com/satori/go.uuid"
)

const envKeyUUIDVersion = "UUID_Version"

type helper struct {
	uuidFunc func() uuid.UUID

	once sync.Once
}

var hObj = &helper{}

func (h *helper) initSelf() {
	if i, _ := strconv.Atoi(os.Getenv(envKeyUUIDVersion)); i == 1 {
		h.uuidFunc = uuid.NewV1
	} else {
		h.uuidFunc = uuid.NewV4
	}
}

func (h *helper) UUID() uuid.UUID {
	h.once.Do(h.initSelf)
	return h.uuidFunc()
}

func (h *helper) Md5(data interface{}) string {
	if bs, ok := data.([]byte); ok {
		return fmt.Sprintf("%x", md5.Sum(bs))
	}
	if s, ok := data.(string); ok {
		return fmt.Sprintf("%x", md5.Sum([]byte(s)))
	}
	return ""
}

func (h *helper) UrlEncode(data interface{}) (result string, err error) {
	if strData, ok := data.(string); ok {
		result = url.QueryEscape(strData)
		return
	}

	if m, ok := data.(url.Values); ok {
		result = m.Encode()
		return
	}

	if m, ok := data.(map[string]string); ok {
		result = h.urlEncodeMss(m)
		return
	}

	if m, ok := data.(map[string]interface{}); ok {
		result = h.urlEncodeMsi(m)
		return
	}

	if m, ok := data.(map[interface{}]interface{}); ok {
		result = h.urlEncodeMii(m)
		return
	}

	err = fmt.Errorf("UrlEncode data type not support")
	return

}

func (h *helper) urlEncodeMss(data map[string]string) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, v)
	}
	return values.Encode()
}

func (h *helper) urlEncodeMsi(data map[string]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(k, fmt.Sprintf("%v", v))
	}
	return values.Encode()
}

func (h *helper) urlEncodeMii(data map[interface{}]interface{}) string {
	values := url.Values{}
	for k, v := range data {
		values.Set(fmt.Sprintf("%v", k), fmt.Sprintf("%v", v))
	}
	return values.Encode()
}
