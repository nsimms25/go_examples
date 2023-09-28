package main

import "fmt"

func strings() {
	var first_string string = "First String, "
	var second_string string = "Second string."
	var combined_strings = first_string + second_string
	fmt.Println(combined_strings)

	var new_string string = "0123456789"
	var character_string string = string(new_string[0])
	fmt.Println(character_string)
}

func main() {
	strings()
}
