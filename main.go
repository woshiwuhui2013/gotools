package main

import (
	"fmt"
	hello "github.com/woshiwuhui2013/gotools/hello"
	"github.com/woshiwuhui2013/gotools/unsafemod"
)

//go:genarate echo "test genarate"
func main() {
	fmt.Println(hello.Hello())

	var aa byte

	aa = 12
	unsafemod.SetVal(&aa)

	i := 20
	ret := unsafemod.RetSizeOf(i)
	fmt.Println(ret)
	fmt.Println(aa)
}
