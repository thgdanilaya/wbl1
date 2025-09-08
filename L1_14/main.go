package main

import (
	"fmt"
	"reflect"
)

func detect(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int: // Switch распознает только определенный канал (int, string, ...)
		fmt.Println("chan int")
	default:
		fmt.Println("unknown")
	}
}

func reflectDetect(v interface{}) {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Chan:
		fmt.Printf("chan type %v \n", reflect.TypeOf(v).Elem().Kind())
	default:
		fmt.Println("unknown")
	}
}

func main() {
	c := make(chan int)
	//y := 6
	reflectDetect(c)
	detect(c)
}
