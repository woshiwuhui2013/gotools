package unsafemod

import "unsafe"

func SetVal(elem *byte) {
	*elem = 30
}

func RetSizeOf(value int) uintptr {
	var i int = value
	return unsafe.Sizeof(i)
}
