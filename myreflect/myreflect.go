package myreflect

import (
	"fmt"
	"reflect"
)

func UseCase() {

	var i int
	i = 20
	rtype := reflect.TypeOf(i)
	fmt.Println(rtype.Name())
}
