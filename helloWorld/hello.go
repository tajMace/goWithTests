package main

import "fmt"

func main() {
	fmt.Println(Hello("Taj"))
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}
