package main

import (
    "fmt"
)

func Add(x int, y int) {
    z := x + y
    fmt.Println(z)
}

func main() {
    for i := 0; i < 10; i++ {
        go Add(i,i)
    }
}
