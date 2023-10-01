package main

import (
	"fmt"
	"strings"
)

//In order to run in test.go, must be of the same package name and run .go files must include this file along with main() files.

func replace_string() {
	//Declare string.
	var full_string string = "Here is a potential result from concat. string1. string2. string3."
	//Replace string.
	replaced_str := strings.Replace(full_string, "string", "String", -1) // int < 0, results in infinite replace.

	//Print formatting.
	fmt.Printf("Before replace: %s \n", full_string)
	fmt.Printf("After replace: %s", replaced_str)
}
