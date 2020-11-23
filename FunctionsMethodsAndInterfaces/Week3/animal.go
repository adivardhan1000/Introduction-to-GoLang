package main

import "fmt"

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	m := map[string]Animal{
		"cow" : Animal{"grass","walk","moo"},
		"bird" : Animal{"worms","fly","peep"},
		"snake" : Animal{"mice","slither","hsss"},
	}
	for{
		fmt.Print(">")
		input1:="0"
		input2:="0"
		fmt.Scan(&input1,&input2)

		if input2=="eat"{
			m[input1].Eat()
		} else if input2=="move"{
			m[input1].Move()
		} else if input2=="speak"{
			m[input1].Speak()
		}
	}
}