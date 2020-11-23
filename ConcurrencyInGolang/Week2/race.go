package main

import "fmt"

/*
Raca condition is when multiple threads are trying to access and manipulate the same variable.

Below, main, increament and decreament are all accessing and changing the value of i.

Due to the uncertainty of GoLang Routine scheduling mechanism, the results of the following program is unpredictable.
*/

func increament(pt *int) {
	(*pt)++
	fmt.Println(*pt)
}

func decreament(pt *int) {
	(*pt)--
	fmt.Println(*pt)
}

func main() {
	i := 0
	go increament(&i)
	go decreament(&i)

	i++
	fmt.Println()

}
