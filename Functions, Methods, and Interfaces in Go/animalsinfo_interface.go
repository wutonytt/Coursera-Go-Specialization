package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow
type Cow struct {
	Name string
}

func (c *Cow) Eat() {
	fmt.Println("grass")
}
func (c *Cow) Move() {
	fmt.Println("walk")
}
func (c *Cow) Speak() {
	fmt.Println("moo")
}

// Bird
type Bird struct {
	Name string
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}
func (b *Bird) Move() {
	fmt.Println("fly")
}
func (b *Bird) Speak() {
	fmt.Println("peep")
}

// Snake
type Snake struct {
	Name string
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}
func (s *Snake) Move() {
	fmt.Println("slither")
}
func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		var command, name, s string
		fmt.Printf("> ")
		fmt.Scan(&command, &name, &s)

		switch command {
		case "newanimal":
			species := s
			switch species {
			case "cow":
				animals[name] = new(Cow)
			case "bird":
				animals[name] = new(Bird)
			case "snake":
				animals[name] = new(Snake)
			}
			fmt.Println("Created")

		case "query":
			info := s
			obj, ok := animals[name]
			if ok {
				switch info {
				case "eat":
					obj.Eat()
				case "move":
					obj.Move()
				case "speak":
					obj.Speak()
				}
			} else {
				fmt.Println("Cannot found!")
			}
		}

	}
}
