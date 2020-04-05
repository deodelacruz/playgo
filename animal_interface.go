/*
Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.

Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss
Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either “cow”, “bird”, or “snake”. Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. The second string is the name of the animal. The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.

Submit your Go program source code.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	processInput()
}

// Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow struct
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Eat function
func (myAnimal Cow) Eat() {
	fmt.Println(myAnimal.food)
}

// Move function
func (myAnimal Cow) Move() {
	fmt.Println(myAnimal.locomotion)
}

// Speak function
func (myAnimal Cow) Speak() {
	fmt.Println(myAnimal.noise)
}

// Bird struct
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Eat function
func (myAnimal Bird) Eat() {
	fmt.Println(myAnimal.food)
}

// Move function
func (myAnimal Bird) Move() {
	fmt.Println(myAnimal.locomotion)
}

// Speak function
func (myAnimal Bird) Speak() {
	fmt.Println(myAnimal.noise)
}

// Snake struct
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Eat function
func (myAnimal Snake) Eat() {
	fmt.Println(myAnimal.food)
}

// Move function
func (myAnimal Snake) Move() {
	fmt.Println(myAnimal.locomotion)
}

// Speak function
func (myAnimal Snake) Speak() {
	fmt.Println(myAnimal.noise)
}

func processInput() {
	fmt.Println("Hi. Please type in a request in this format: animal action")
	fmt.Println("Note: animal can be either cow, bird or snake")
	fmt.Println("Note: action can be either eat, move or speak")
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			animalInput := strings.Fields(scanner.Text())
			fmt.Printf("%v %v \n", animalInput, len(animalInput))
			if len(animalInput) != 2 {
				fmt.Println("Oops. You have to specify exactly 2 strings.")
			} else {
				animalType := animalInput[0]
				animalAction := animalInput[1]
				// instantiate new animal based on input
				var myAnimal Animal
				switch animalType {
				case "cow":
					cow := Cow{"grass", "walk", "moo"}
					myAnimal = cow
				case "bird":
					bird := Bird{"worms", "fly", "peep"}
					myAnimal = bird
				case "snake":
					snake := Snake{"mice", "slither", "hiss"}
					myAnimal = snake
				default:
					{
						fmt.Println("Oops. You have to specify either cow, bird or snake.")
					}
				}
				switch animalAction {
				case "eat":
					myAnimal.Eat()
				case "move":
					myAnimal.Move()
				case "speak":
					myAnimal.Speak()
				default:
					{
						fmt.Println("Oops. You have to specify either eat, move or speak.")
					}
				}
			}
			break
		}
		if scanner.Err() != nil {
			// handle error.
		}
	}
}
