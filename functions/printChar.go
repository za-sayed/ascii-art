package functions

import (
	"fmt"
	"os"
)

func PrintChar(str string, asciiRep [][]string) {
	arr := LineNum(str)
	for i := 0; i < 8; i++ {
		for _, num := range arr {
			fmt.Print(asciiRep[num][i])
		}
		fmt.Println("")
	}
}

func LineNum(str string) []int {
	var arr []int
	for _, char := range str {
		if char >= 32 && char <= 126 {
			arr = append(arr, (int(char) - 32))
		} else {
			fmt.Println("Error: a non ascii character found")
			os.Exit(1);
		}	
	}
	return arr
}
