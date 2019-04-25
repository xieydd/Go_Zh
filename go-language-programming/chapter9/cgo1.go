package main

import (
    "fmt"
)
/*
#include <stdlib.h>
*/
import "C"

func Random() int {
    return int(C.random())
}

func Seed(num int) {
    C.srandom(C.uint(num))
}

func main() {
    Seed(10)
    fmt.Println("Random:", Random())
}
