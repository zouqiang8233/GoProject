package main

import (
	"fmt"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func writeTest() {
	// 打开文件
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println("OpenFile Err: ", err.Error())
		return
	}

	// 关闭文件
	defer db.Close()

	// 设置写入为落地模式
	writeOp := &opt.WriteOptions{
		Sync: true,
	}

	fmt.Println("writeTest")

	// 打印写入数据时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))

	for j := 0; j < 1000; j++ {

		batch := &leveldb.Batch{}

		// 开启每1000条数据写入一次
		for i := 0; i < 1000; i++ {
			batch.Put([]byte(string(j*1000+i)), []byte(string(j*1000+i)))
		}

		// 写入文件，文件写入后返回
		err := db.Write(batch, writeOp)

		if err != nil {
			fmt.Println("Write Err: ", err.Error())
			return
		}
	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}

func readTest() {
	// 打开文件
	db, err := leveldb.OpenFile("db", nil)
	if err != nil {
		fmt.Println("OpenFile Err: ", err.Error())
		return
	}

	// 关闭文件
	defer db.Close()

	fmt.Println("readTest")

	// 打印写入数据时间
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))

	for j := 0; j < 1000; j++ {
		for i := 0; i < 1000; i++ {
			_, err := db.Get([]byte(string(j*1000+i)), nil)

			if err != nil {
				fmt.Println("Get Err :", err.Error())
				return
			}

			//fmt.Println(string(value))
		}

	}

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.000"))
}

func main() {
	writeTest()
	readTest()
}
