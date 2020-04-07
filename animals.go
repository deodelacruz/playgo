/*
Write a program which allows the user to get information about a predefined set of animals. Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. The user can issue a request to find out one of three things about an animal: 1) the food that it eats, 2) its method of locomotion, and 3) the sound it makes when it speaks. The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal	Food eaten	Locomotion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt. Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings. The first string is the name of an animal, either “cow”, “bird”, or “snake”. The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. Make three methods called Eat(), Move(), and Speak(). The receiver type of all of your methods should be your Animal type. The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. Your program should call the appropriate method when the user makes a request.
*/

package playgo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	readInput()
}

// Animal struct
type Animal struct {
	food       string
	locomotion string
	noise      string
}

func readInput() {
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
					myAnimal = Animal{"grass", "walk", "moo"}
				case "bird":
					myAnimal = Animal{"worms", "fly", "peep"}
				case "snake":
					myAnimal = Animal{"mice", "slither", "hiss"}
				default:
					{
						fmt.Println("Oops. You have to specify either cow, bird or snake.")
					}
				}
				switch animalAction {
				case "eat":
					fmt.Printf("This animal eats: %s \n", myAnimal.Eat())
				case "move":
					fmt.Printf("This animal's movement: %s \n", myAnimal.Move())
				case "speak":
					fmt.Printf("This animal speaks: %s \n", myAnimal.Speak())
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

// Eat function
func (myAnimal Animal) Eat() string {
	return myAnimal.food
}

// Move function
func (myAnimal Animal) Move() string {
	return myAnimal.locomotion
}

// Speak function
func (myAnimal Animal) Speak() string {
	return myAnimal.noise
}
