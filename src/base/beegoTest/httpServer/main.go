package main

import (
	"fmt"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Post() {
	// 设置不走默认的响应
	this.EnableRender = false

	// 获取请求消息的长度
	iLen := this.Ctx.Request.ContentLength

	// 申请数组接收
	data := make([]byte, iLen)

	// 读取数据
	iReadLen, err := this.Ctx.Request.Body.Read(data)

	if err != nil && err.Error() != "EOF" {
		fmt.Println("err :", err)
	} else {
		fmt.Println("Post: ", string(data), "iReadLen: ", iReadLen, "iLen: ", iLen)
		this.Ctx.WriteString("hello")
	}

	// 关闭请求
	this.Ctx.Request.Body.Close()
}

func (this *MainController) Get() {
	this.EnableRender = false
	this.Ctx.WriteString("hello")
}

func main() {
	mainTmp := &MainController{}

	// 映射/ 到该类
	beego.Router("/", mainTmp)

	beego.Run()
}
