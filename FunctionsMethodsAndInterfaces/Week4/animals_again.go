package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (objectAnimal animal) Eat() {
	fmt.Println(objectAnimal.food)
	return
}

func (objectAnimal animal) Move() {
	fmt.Println(objectAnimal.locomotion)
	return
}

func (objectAnimal animal) Speak() {
	fmt.Println(objectAnimal.noise)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	var tempAnimal animalInterface
	for {
		var command, requestAnimal, requestType string
		fmt.Print(">")
		fmt.Scan(&command, &requestAnimal, &requestType)
		if command == "query" {
			tempAnimal = animalMap[requestAnimal]
			switch requestType {
			case "eat":
				tempAnimal.Eat()
			case "move":
				tempAnimal.Move()
			case "speak":
				tempAnimal.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[requestAnimal] = animalMap[requestType]
			fmt.Println("Created")
		}
	}
}