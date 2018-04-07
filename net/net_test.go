package net

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_NetGet(t *testing.T) {
	Convey("get logics", t, func() {
		resp, err := New().Get("http://localhost:8080/get").End(nil, nil)
		if err != nil {
			Printf("err: %v", err)
			return
		}
		Printf("resp: %v", resp)
		So(resp.StatusCode, ShouldEqual, 200)
	})

}

func Test_NetPost(t *testing.T) {
	Convey("post logics", t, func() {
		body := &struct {
			Param string `json:"param"`
		}{Param: "请求内容"}
		result := &struct {
			Param string `json:"param"`
			Msg   string `json:"msg"`
		}{}
		resp, err := New().Post("http://localhost:8080/post").JSON(body).End(nil, result)
		Printf("resp: %v\n", resp)
		Printf("err: %v\n", err)
		Printf("result: %v\n", result)
		So(resp.StatusCode, ShouldEqual, 200)
	})
}
