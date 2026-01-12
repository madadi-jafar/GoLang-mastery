package main

// Topic: Embedding Files (Go 1.16+)
// Task: Embed a text file and print its content.
// TODO: Create hello.txt in task82/ with "Embedded content"
// TODO: Use //go:embed hello.txt
// TODO: Print content
import _ "embed"

//go:embed hello.txt
var embeddedContent string

func main() {
	// Print embeddedContent
}
