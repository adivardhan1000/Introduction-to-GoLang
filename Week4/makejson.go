package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var name string
	var addr string

	person := make(map[string]string)

	fmt.Printf("\nPlease enter a name: ")
	fmt.Scan(&name)

	fmt.Printf("Please enter an address: ")
	fmt.Scan(&addr)

	person["name"] = name
	person["address"] = addr

	jsonperson, _ := json.Marshal(person)

	fmt.Println("Your JSON Object is: ")
	fmt.Println(string(jsonperson))
}
