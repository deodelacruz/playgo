/*
Write a program which prompts the user to first enter a name, and then enter an
address. Your program should create a map and add the name and address to the
map using the keys “name” and “address”, respectively. Your program should use
Marshal() to create a JSON object from the map, and then your program should
print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Hi. What is your name?")
	name, _ := reader.ReadString('\n')
	fmt.Printf("Name: %v\n", name)

	fmt.Println("What is your address?")
	address, _ := reader.ReadString('\n')
	fmt.Printf("Address: %v\n", address)

	map1 := make(map[string]string)
	map1["name"] = name
	map1["address"] = address
	//fmt.Printf("Map Person: %v\n", map1)

	json1, _ := json.Marshal(map1)
	fmt.Printf("Json: %v\n", string(json1))

}
