package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Name           string
	LifeExpectance int
}

func (b *Bird) Fly() {
	fmt.Println("I am Flaying")
}

func main() {
	bird := &Bird{"zsz", 100}
	e := reflect.ValueOf(bird).Elem()
	typeOfT := e.Type()
	for i := 0; i < e.NumField(); i++ {
		f := e.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
