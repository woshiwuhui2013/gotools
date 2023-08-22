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

	fmt.Println(aa)
}
