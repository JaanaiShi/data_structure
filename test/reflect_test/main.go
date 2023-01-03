package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int
	a = 1
	typeOfA := reflect.TypeOf(a)
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
	var b *int
	valueOfA := reflect.ValueOf(b)
	fmt.Println(valueOfA.Type().Elem())
}