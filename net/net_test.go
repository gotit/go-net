package net

import (
	"log"
	"testing"
)

func Test_Get(t *testing.T) {
	_, err := New().Get("http://localhost:80/get").End(nil, nil)
	if err != nil {
		t.Errorf("Get returned error: %v", err)
	} else {
		log.Print("send get request success.")
	}
}

func Test_Post(t *testing.T) {
	body := &struct {
		Param string `json:"param"`
	}{Param: "请求内容"}
	result := &struct {
		Param string `json:"param"`
		Msg   string `json:"msg"`
	}{}
	_, err := New().Post("http://localhost:80/post").JSON(body).End(nil, result)
	if err != nil {
		t.Errorf("Post returned error: %v", err)
	} else {
		log.Printf("send post request success. response: %+#v", result)
	}
}
