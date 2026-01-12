package main

// Topic: unsafe.Pointer
// Task: Convert []byte to string without allocation (unsafe).
// TODO: Use unsafe.String and unsafe.Slice (Go 1.20+)
// Or use old method with uintptr if needed
import "unsafe"

func BytesToString(b []byte) string {
	// Your code here (unsafe)
	return ""
}

func main() {
	b := []byte("unsafe string")
	s := BytesToString(b)
	println(s)
}
