package unsafemod

import (
	"fmt"
	"unsafe"
)

func SetVal(elem *byte) {
	*elem = 30
}

type Foo struct {
	A int
	B string
	C [2]int
	D float64
}

func RetSizeOf(value int) uintptr {
	var i int = value

	var foo Foo
	fmt.Println(unsafe.Sizeof(foo))
	return unsafe.Sizeof(i)
}

func PointTo(foo *Foo) {

	p := unsafe.Pointer(foo)

	var a = (*int)(p)

	fmt.Println(*a)

}

func UintptrHandler() {
	foo := Foo{A: 1, B: "hhhh", C: [2]int{3, 5}, D: 1.1}
	a := uintptr(unsafe.Pointer(&foo))
	fmt.Println("x%", a)
}
