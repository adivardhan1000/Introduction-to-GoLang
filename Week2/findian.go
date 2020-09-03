package main

import (
	"fmt"
	"strings"
)

func main(){
	var inpString string = ""
	fmt.Print("Enter a string to search for \"ian\" pattern: ")
	fmt.Scan(&inpString)
	y := strings.ReplaceAll(strings.ToLower(inpString), " ", "")
	
	if strings.HasPrefix(y, "i")  && strings.Contains(y, "a") && strings.HasSuffix(y, "n"){
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}