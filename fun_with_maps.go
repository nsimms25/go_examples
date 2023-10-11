package main

import (
	"fmt"
)

func map_practice() {
	cust_contact := map[string]string{}
	cust_contact["John"] = "1234"
	cust_contact["Dave"] = "5609"
	fmt.Println(cust_contact)
	fmt.Println(cust_contact["John"])
	fmt.Println(cust_contact["Dave"])
}
