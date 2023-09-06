package main

import (
	"context"
	"fmt"
	CtxPkg "github.com/woshiwuhui2013/gotools/context-test"
	"time"
)

//go:genarate echo "test genarate"
func main() {
	//fmt.Println(hello.Hello())
	//
	//var aa byte
	//
	//aa = 12
	//unsafemod.SetVal(&aa)
	//
	//i := 20
	//ret := unsafemod.RetSizeOf(i)
	//fmt.Println(ret)
	//
	//foo := unsafemod.Foo{A: 1, B: "aa", C: [2]int{1, 2}, D: 1.1}
	//
	//unsafemod.PointTo(&foo)
	//
	//unsafemod.UintptrHandler()
	//myreflect.UseCase()
	//myreflect.ParseStu(foo)
	CtxPkg.CtxHello()

	timeout, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Hello")
	case <-timeout.Done():
		fmt.Println("finish finish")

	}


}
