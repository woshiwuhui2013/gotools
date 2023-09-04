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
	v := reflect.ValueOf(i)
	iv := v.Interface()
	fmt.Println("iv", iv)

	funval := reflect.ValueOf(add)

	var j int
	j = 39
	arr := []reflect.Value{reflect.ValueOf(i), reflect.ValueOf(j)}
	z := funval.Call(arr)

	fmt.Println("z {}", z[0].Int())

}

func ParseStu(itf interface{})  {
	reflectVal := reflect.ValueOf(itf)
	reflectType := reflect.TypeOf(itf)

	for i := 0; i < reflectType.NumField(); i++ {
		val := reflectVal.Field(i)
		sttype := reflectType.Field(i)

		fmt.Println("name ", sttype.Name)
		fmt.Println("type ", sttype.Type)

		fmt.Println("val ", val.Interface())

	}

}
