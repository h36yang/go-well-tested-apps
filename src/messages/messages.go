package messages

import (
	"fmt"
)

// Greet says Hello
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!\n", name)
}

// depart says Goodbye
func depart(name string) string {
	return fmt.Sprintf("Goodbye, %s\n", name)
}
