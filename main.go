package main

import "fmt"

func main() {
	Print("git learning")
}

func Print(a ...interface{}) {
	fmt.Println(a...)
}
