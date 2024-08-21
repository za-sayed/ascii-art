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
		fileLines := functions.Read()
		asciiRep := functions.AsciiRep(fileLines)
		functions.PrintStr(text, asciiRep)
	}
}
