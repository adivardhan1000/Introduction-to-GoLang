package main

import "fmt"

func main(){

    var tempnum float64

    fmt.Printf("Please enter a floating point number: ")
    num, err := fmt.Scan(&tempnum)

    trunc := int(tempnum)

    if (num == 1 && err == nil){
        fmt.Printf("Your integer is: %d\n", trunc)
    } else{
        fmt.Printf("Input Error\n")
    }
}