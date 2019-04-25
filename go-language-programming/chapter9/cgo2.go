package main

import (
    "fmt"
)

/*
#include <stdio.h>
int hello() {
    printf("Hello, Cgo! \n");
    return 1;
}
*/
import "C"

func Hello() int {
    return int(C.hello())
}

func main() {
    fmt.Println(Hello())
}

/*
#cgo CFLAGS: -DPNG_DEBUG=1
#cgo linux CFLAGS: -DLINUX=1
#cgo LDFLAGS: -lpng
#include <png.h>
*/

//or

// #cgo pkg-config: png cairo
// #include <png.h>
