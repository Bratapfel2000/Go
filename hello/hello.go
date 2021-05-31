package main

import "fmt"

// adding constants
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world"))
}