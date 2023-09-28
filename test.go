package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func strings_examples() {
	var first_string string = "First String, "
	var second_string string = "Second string."
	var combined_strings = first_string + second_string
	fmt.Println(combined_strings)

	var new_string string = "0123456789"
	var character_string string = string(new_string[0])
	fmt.Println(character_string)

}

func string_contains() {
	var string_to_search string = ""
	var sub_string string = ""

	fmt.Println("Please input a string to search.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	string_to_search = scanner.Text()

	fmt.Println("Please input a substring to search for.")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	sub_string = scanner.Text()

	if strings.Contains(string_to_search, sub_string) {
		fmt.Printf(`Substring '%s' is in '%s'.`, sub_string, string_to_search)
	}
	if !strings.Contains(string_to_search, sub_string) {
		fmt.Printf(`Substring '%s' is not in '%s'.`, sub_string, string_to_search)
	}
}

func ints() {

}

func main() {
	//strings()
	string_contains()
}
