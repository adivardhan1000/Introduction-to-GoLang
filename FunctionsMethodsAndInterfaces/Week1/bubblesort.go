/*Write a Bubble Sort program in Go. The program should prompt the user to type in a sequence of up to 10 integers. The program should print the integers out on one line, in sorted order, from least to greatest. Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing. The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice. You should write a Swap() function which performs this operation. Your Swap() function should take two arguments, a slice of integers and an index value i which indicates a position in the slice. The Swap() function should return nothing, but it should swap the contents of the slice in position i with the contents in position i+1.
*/
package main

import (
	"fmt"
)

func main() {
	//var tempinput int
	var intArray [10]int

	fmt.Printf("Enter 10 integers")
	for i := 0; i < 10; i++ {
		fmt.Scan(&intArray[i])
	}

	fmt.Println(intArray)
	BubbleSort(intArray)
	fmt.Println("Bubble Sort Completed")

}

func BubbleSort(a [10]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j+1] < a[j] {
				//fmt.Println("inside condition")
				a = Swap(a, j)
			}
		}
	}
	fmt.Println("Bubble Sorted array:", a)
}

func Swap(array [10]int, j int) [10]int {
	var temp int
	//fmt.Println("to be swapped",a[j],a[j+1])
	temp = array[j]
	array[j] = array[j+1]
	array[j+1] = temp
	return array
}

/*
func bubblesort(intArray [10]int)  {
	for i := 0; i < len(intArray); i++ {
        for j := 0; j < len(intArray)-1-i; j++  {
			fmt.Println("INside")
            if intArray[j+1] < intArray[j] {
				fmt.Println("INside2")
                Swap(intArray,j)
            }
        }
	}
	fmt.Println("Bubble Sorted array:",intArray)
}

func Swap(intArray [10]int,j int) {
	var temp int
	temp = intArray[j]
	intArray[j] = intArray[j+1]
	intArray[j+1] = temp
	fmt.Println("Swapped")
}*/
