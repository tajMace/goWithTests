package main

import "fmt"

const helloPrefix = "Hello, "

func main() {
	fmt.Println(Hello("Taj"))
}

func Hello(name string) string {
	return helloPrefix + name + "!"
}
