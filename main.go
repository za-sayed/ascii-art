package main

import (
	"ascii-art/functions"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"<text>\"")
		return
	} else {
		text := os.Args[1]
		if !functions.ValidateASCII(text) {
			fmt.Println("Error: Non-ASCII character detected")
			os.Exit(1) 
		}
		fileLines := functions.Read()
		asciiRep := functions.AsciiRep(fileLines)
		functions.PrintStr(text, asciiRep)
	}
}
