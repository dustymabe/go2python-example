package main

import "C"
import "github.com/fatih/color"

//export Hello
func Hello() {
	c := color.New(color.FgCyan).Add(color.Underline)
	c.Println("Hello! I can print cyan text with an underline.")
}

func main() {}
