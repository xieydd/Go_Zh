package main

/*
#include <stdio.h>
#include <stdlib.h>
void c_print(char *str) {
    printf("%s\n", str);
}
*/
import "C"
import "unsafe"

func main() {
	cstr := C.CString("Xieyuandong")
	C.puts(cstr)
    C.c_print(cstr)
	defer C.free(unsafe.Pointer(cstr))
}
