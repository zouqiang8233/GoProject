package main

import (
	"base/DBAccess"
	"base/dbaccess/protocol"
	"fmt"
)

func main() {
	db := dbaccess.NewDBAccess("mysql", "root:muchinfo@/test")
	db.ConnDB()

	data, err := db.Query("student", []string{}, "", &test.Student{})

	if err == nil {
		for _, v := range data {
			tmp := v.(*test.Student)
			fmt.Println(tmp)
			fmt.Println(tmp.GetAge(), tmp.GetName(), tmp.GetNO())
		}

	} else {
		fmt.Println("Query err ", err)
	}
}
