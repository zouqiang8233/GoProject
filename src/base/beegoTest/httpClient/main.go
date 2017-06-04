package main

import (
	"encoding/xml"
	"fmt"

	"github.com/astaxie/beego/httplib"
)

func main() {
	// 打包协议
	signReq := SignReq{
		Header: &MsgHeader{
			MsgCode: "10301",
			Version: "1.0",
		},
		Bodyer: &SignReqBody{
			ExchNo:   "ExchNo",
			ExchDate: "ExchDate",
			Exinfo: &ExtendInfo{
				BankPassword: "BankPassword",
				Sex:          "Sex",
			},
		},
	}

	//	打包数据
	data, err := xml.Marshal(&signReq)

	if err != nil {
		fmt.Println("xml.Marshal, err :", err.Error())
		return
	}

	// 设置POST地址
	req := httplib.Post("http://127.0.0.1:8080/")

	// POST数据并且接收回应
	revData, err := req.Body(data).String()

	if err != nil {
		fmt.Println("post data err, err :", err.Error())
	} else {
		fmt.Println("get data :", revData)
	}

	// 获取返回码
	rsp, err := req.Response()

	if err != nil {
		fmt.Println("get response err, err :", err.Error())
	} else {
		fmt.Println("retcode :", rsp.StatusCode)
	}
}
