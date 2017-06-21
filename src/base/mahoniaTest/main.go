package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/axgle/mahonia"
)

func ConvertToString(src string, srcCode string, tagCode string) (string, error) {
	// 创建源字符集编码器
	srcCoder := mahonia.NewDecoder(srcCode)

	// 判断是否成功
	if nil == srcCoder {
		return "", errors.New("Src NewDecoder Err")
	}

	// 将源字符编码成UTF-8
	srcUTF8 := srcCoder.ConvertString(src)

	// 创建目标字符集解码器
	tagCoder := mahonia.NewEncoder(tagCode)

	// 判断是否成功
	if nil == tagCoder {
		return "", errors.New("Tag NewEncoder Err")
	}

	// 将UTF8转换成目标字符集
	tag := tagCoder.ConvertString(srcUTF8)

	return tag, nil
}

func main() {
	strTag, err := ConvertToString("你好", "utf-8", "big5")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(strTag)
	}

	// 保存到文件
	f, _ := os.Create("./test.txt")
	f.WriteString(strTag)
	defer f.Close()
}
