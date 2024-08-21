package functions

import (
	"fmt"
	"os"
	"strings"
)

func PrintStr(str string, asciiRep [][]string) {
	subStrs := SubStrs(str)
	for _, subStr := range subStrs {
		if !ValidateASCII(subStr) {
			fmt.Println("Error: Non-ASCII character detected")
			os.Exit(1) 
		}
	}
	len := len(subStrs)
	for index, subStr := range subStrs {
		if subStr == "\\n" {
			if index == len - 1 {
				fmt.Print("\n")
			} else if index == 0 {
				fmt.Print("\n") 
			} else {
				if subStrs[index - 1] == "\\n" {
					fmt.Print("\n")
				}
			}
		} else {
			PrintChar(subStr, asciiRep)
		}
	}
}

func SubStrs(str string) []string {
	parts := strings.Split(str, "\\n")
	var result []string
	for i, part := range parts {
		if part != "" {
			result = append(result, part)
		}
		if i < len(parts) - 1 {
			result = append(result, "\\n")
		}
	}
	return result
}

func ValidateASCII(str string) bool {
	for _, char := range str {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}
