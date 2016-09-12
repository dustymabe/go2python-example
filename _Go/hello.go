package main

// go build -buildmode=c-shared .

import "C"
import "github.com/fatih/color"


//export Hello
func Hello() {
    c := color.New(color.FgCyan).Add(color.Underline)
    c.Println("Prints cyan text with an underline.")
}


func main() {}
