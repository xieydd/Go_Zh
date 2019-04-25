package main

import (
    "reflect"
    "fmt"
)

type T struct {
    A int
    B string
}

func main() {
    var x float64 = 0.3334
    fmt.Printf("%.2f type is %s\n", x, reflect.TypeOf(x))
    fmt.Printf("kind is float64: %v\n", reflect.TypeOf(x).Kind() == reflect.Float64)

    p := reflect.ValueOf(&x)
    fmt.Println("type of p:", p.Type())
    fmt.Println("settability of p", p.CanSet())

    m := p.Elem()
    fmt.Println("settability of m", m.CanSet())

    m.SetFloat(7.1)
    fmt.Println(m.Interface())
    fmt.Println(x)

    t := T{1, "xieyd"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
}
