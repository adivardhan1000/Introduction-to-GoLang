package main

import "fmt"

func GenDisplaceFn(a,v,s float64) func(float64) float64{
	return func(t float64) float64 {
		return v*t+1/2*a*t*t+s
	}
}

func main() {
	var a float64
		fmt.Print("Enter acceleration:")
		fmt.Scanln(&a)
	var v float64
		fmt.Print("Enter the initial velocity:")
		fmt.Scanln(&v)
	var s float64
		fmt.Print("Enter initial displacement:")
		fmt.Scanln(&s)
	var t float64
		fmt.Print("Enter time:")
		fmt.Scanln(&t)
	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

