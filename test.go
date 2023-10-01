package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func strings_examples() {
	//Combining strings, using '+' operator.
	var first_string string = "First String example, "
	var second_string string = "Second string example."
	var combined_strings = first_string + second_string
	fmt.Println(combined_strings)

	//Using string indexing.
	var new_string string = "0123456789"
	var character_string string = string(new_string[0])
	fmt.Println(character_string)

	//Finding a substring index, printing from substring start to end (using slicing).
	var str_index int = 0
	str_index = strings.Index(first_string, "String")
	fmt.Printf("String index: %d \n", str_index)
	fmt.Printf(first_string[str_index : len(first_string)-1])

}

func string_contains() {
	var string_to_search string = ""
	var sub_string string = ""

	fmt.Println("Please input a string to search.")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	string_to_search = scanner.Text()

	fmt.Println("Please input a substring to search for.")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sub_string = scanner.Text()

	if strings.Contains(string_to_search, sub_string) {
		fmt.Printf(`Substring '%s' is in '%s'.`, sub_string, string_to_search)
	}
	if !strings.Contains(string_to_search, sub_string) {
		fmt.Printf(`Substring '%s' is not in '%s'.`, sub_string, string_to_search)
	}
}

func cut_string() {
	fmt.Println("What string do you want to cut?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	full_string := scanner.Text()

	fmt.Println("What is the separator?")
	scanner = bufio.NewScanner(os.Stdin)
	scanner.Scan()
	string_sep := scanner.Text()

	before_str, after_str, found_bool := strings.Cut(full_string, string_sep)

	fmt.Printf(`Before string: %s; After string: %s; Found: %t`, before_str, after_str, found_bool)
}

func split_string() {
	fmt.Println("What string do you want to split?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	full_string := scanner.Text()

	split_string := strings.Fields(full_string)
	array_len := len(split_string)

	for i := 0; i < array_len; i++ {
		fmt.Println(split_string[i])
	}

}

func ints() {

}

func main() {
	strings_examples()
	//string_contains()
	//cut_string()
	//split_string()
}
