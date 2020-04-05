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
	name       string
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
	name       string
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
	name       string
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
	fmt.Println("Hi. Please type in a cmd in either of the formats below: ")
	fmt.Println("a) newanimal animalname animaltype (where animaltype is either cow, bird or snake)")
	fmt.Println("b) quey animalname action (where action is either eat,move or speak)")

	mapAnimals := make(map[string]Animal)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			animalInput := strings.Fields(scanner.Text())
			//fmt.Printf("User Input: %v %v \n", animalInput, len(animalInput))
			if len(animalInput) != 3 {
				fmt.Println("Oops. You have to specify exactly 3 strings.")
			} else {
				var myAnimal Animal
				cmd := animalInput[0]
				animalName := animalInput[1]
				arg3 := animalInput[2]
				switch cmd {
				case "newanimal":
					{
						// instantiate new animal based on input
						isNewAnimalCreated := false
						animalType := arg3
						switch animalType {
						case "cow":
							cow := Cow{animalName, "grass", "walk", "moo"}
							myAnimal = cow
							isNewAnimalCreated = true
						case "bird":
							bird := Bird{animalName, "worms", "fly", "peep"}
							myAnimal = bird
							isNewAnimalCreated = true
						case "snake":
							snake := Snake{animalName, "mice", "slither", "hiss"}
							myAnimal = snake
							isNewAnimalCreated = true
						default:
							{
								fmt.Println("Oops. You have to specify either cow, bird or snake.")
								isNewAnimalCreated = false
							}

						} // switch animalType
						if isNewAnimalCreated {
							mapAnimals[animalName] = myAnimal // add new animal to map wih name as key
							fmt.Println("Created it!")
						}
						fmt.Printf("mapAnimals: %v \n", mapAnimals)
					} //case: newanimal
				case "query":
					{
						// get object for given name - implement this via map
						myAnimal = mapAnimals[animalName]
						animalAction := arg3
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
						} //switch animalAction
					} //case query
				default:
					{
						fmt.Println("Oops. Your 1st argument should be either newanimal or query.")
					}
				} // switch  cmd
			} //else
			break
		}
		if scanner.Err() != nil {
			// handle error.
		}
	}
}
