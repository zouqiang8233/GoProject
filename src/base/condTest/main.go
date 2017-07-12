package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var mutexTest *sync.Mutex
var condTest *sync.Cond

func TestFunOne() {
	time.Sleep(time.Second * 1)

	mutexTest.Lock()
	condTest.Broadcast()
	mutexTest.Unlock()
}

func TestFunTwo() {
	mutexTest.Lock()
	defer mutexTest.Unlock()
	for {
		condTest.Wait()
		fmt.Println("TestFunTwo1111")
		time.Sleep(time.Second * 2)
		fmt.Println("TestFunTwo")
	}
}

func TestFunThree() {
	mutexTest.Lock()
	defer mutexTest.Unlock()
	for {
		condTest.Wait()
		fmt.Println("TestFunThree1111")
		time.Sleep(time.Second * 3)
		fmt.Println("TestFunThree")
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	mutexTest = &sync.Mutex{}
	condTest = sync.NewCond(mutexTest)
	go TestFunOne()
	go TestFunTwo()
	go TestFunThree()
	time.Sleep(time.Second * 30)
}
