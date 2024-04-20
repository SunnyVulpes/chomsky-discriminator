package pkg

import (
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func isUpperCase(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	} else {
		return false
	}
}

func isLowerCase(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	} else {
		return false
	}
}

func isNumber(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}

func isCase(c byte) bool {
	if isNumber(c) || isUpperCase(c) || isLowerCase(c) {
		return true
	} else {
		return false
	}
}
