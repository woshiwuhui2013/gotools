package myreflect

import (
	"fmt"
	"reflect"
)

func add(i int, j int) int {

	return i + j
}

func UseCase() {

	var i int
	i = 20
	rtype := reflect.TypeOf(i)
	fmt.Println(rtype.Name())

	funval := reflect.ValueOf(add)

	var j int
	j = 39
	arr := []reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)}
	z := funval.Call(arr)

	fmt.Println("z {}", z[0].Int())

}
