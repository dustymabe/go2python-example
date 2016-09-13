package main

import "C"
import "github.com/fatih/color"

//export Hello
func Hello() {
	c := color.New().Add(color.Underline)
	c.Println("Hello! I can print text with an underline.")
}

func main() {
	Hello()
}
