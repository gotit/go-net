package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kataras/iris"
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
		ctx.JSON(iris.StatusOK, nil)
	})
	router.POST("/post", func(ctx *gin.Context) {
		req := &testRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			ctx.JSON(iris.StatusBadRequest, &testRespose{Msg: "请求结构有问题"})
			return
		}
		ctx.JSON(iris.StatusOK, &testRespose{Param: req.Param, Msg: "请求成功"})
	})
	log.Fatal(router.Run(":80"))
}
