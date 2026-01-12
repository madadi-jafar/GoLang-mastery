package main

// Topic: Build Tags
// Task: Create two implementations (debug/prod) selected via build tag.
// TODO: Later, create task83_debug.go and task83_prod.go
// For now, just know Mode() should return string
func Mode() string {
	return "unknown"
}

func main() {
	println("Mode:", Mode())
}
