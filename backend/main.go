package main

import "fmt"

// GetHelloWorldString возвращает строку приветствия
func GetHelloWorldString(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func main() {
	fmt.Println(GetHelloWorldString("slimz"))
}
