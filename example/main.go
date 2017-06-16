package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	type testRequest struct {
		Param string `json:"param"`
	}
	type testRespose struct {
		Param string `json:"param"`
		Msg   string `json:"msg"`
	}

	router := gin.Default()
	router.GET("/get", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})
	router.POST("/post", func(ctx *gin.Context) {
		req := &testRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &testRespose{Msg: "请求结构有问题"})
			return
		}
		ctx.JSON(http.StatusOK, &testRespose{Param: req.Param, Msg: "请求成功"})
	})
	log.Fatal(router.Run(":80"))
}
