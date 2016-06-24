package hello

import "fmt"

const testVersion = 2

// HelloWorld greets the user by name, or by saying "Hello, World!" if no name is given.
func HelloWorld(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprintf("Hello, %v!", name)
}
