package main

import "fmt"

type Animal0 struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal0) Eat() {
	fmt.Println(a.food)
}
func (a Animal0) Move() {
	fmt.Println(a.locomotion)
}
func (a Animal0) Speak() {
	fmt.Println(a.noise)
}

func (animal Animal0) Motion(locomotion string) {
	switch locomotion {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	}
}

func main() {
	// predefine the info of animals
	cow := Animal0{"grass", "walk", "moo"}
	bird := Animal0{"worms", "fly", "peep"}
	snake := Animal0{"mice", "slither", "hsss"}

	for {
		var name, info string
		fmt.Println("Enter the name of an animal and the name of the information requested about the animal: ")
		fmt.Printf("> ")
		fmt.Scan(&name, &info)
		switch name {
		case "cow":
			cow.Motion(info)
		case "bird":
			bird.Motion(info)
		case "snake":
			snake.Motion(info)
		}
	}
}
