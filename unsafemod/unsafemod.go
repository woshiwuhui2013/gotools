package unsafemod

import (
	"fmt"
	"unsafe"
)

func SetVal(elem *byte) {
	*elem = 30
}

type Foo struct {
	a int
	b string
	c [10]int
	d float64
}

func RetSizeOf(value int) uintptr {
	var i int = value

	var foo Foo
	fmt.Println(unsafe.Sizeof(foo))
	return unsafe.Sizeof(i)
}
