package main

// Topic: CGO – Calling C
// Task: Call C's printf from Go.
// TODO: Use import "C" and C.printf
// Note: Requires GCC installed
/*
#include <stdio.h>
*/
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello from C!\n")
	defer C.free(unsafe.Pointer(cs))
	C.printf(cs)
}
