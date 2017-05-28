package main

import (
	"base/DBAccess"
	"base/protobuf/protocol"
	"fmt"

	//"github.com/golang/protobuf/proto"
)

/* 使用说明:
   1.将src/base/third/protobuf-master.zip文件解压到
     src\github.com\golang下，修改protobuf-master的文件夹名字为protobuf。
   2.将src/base/third/protoc-gen-go.exe，
       src/base/third/protoc_win32.exe或者protoc_win64.exe拷贝到bin目录下。
   3.定义src\base\protobuf\protocol\test.proto测试协议文件，使用
     src\base\protobuf\protocol\generate.bat生成test.pb.go文件。
*/

func main() {

	db := dbaccess.NewDBAccess("mysql", "root:muchinfo@/test")
	db.ConnDB()

	data, err := db.Query("student", nil, "", &test.Student{})

	if err == nil {
		for _, v := range data {
			tmp := v.(*test.Student)
			fmt.Println(tmp.GetAge(), tmp.GetName(), tmp.GetNO())
		}

	} else {
		fmt.Println("Query err ", err)
	}
}
