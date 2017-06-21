package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	c := exec.Command("cmd", "/c", "del", "d:\\a.txt")

	if err := c.Run(); err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("OK")
	}

	result, err := exec.Command("ping", "127.0.0.1").Output()

	f, _ := os.Create("./test.txt")
	f.WriteString(string(result))
	defer f.Close()

	fmt.Println(string(result), err)
}
