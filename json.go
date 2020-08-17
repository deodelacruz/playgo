/* Ask user for name and address, create map of it and print it out as json */

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//ask for name
	reader1 := bufio.NewReader(os.Stdin)
	fmt.Println("Hi, what is your name?")
	name, err := reader1.ReadString('\n')
	if err != nil {
		panic(err)
	}
	//ask for address
	fmt.Println("Hi, what is your address?")
	address, err := reader1.ReadString('\n')
	if err != nil {
		panic(err)
	}

	//create map with name as key
	personMap := make(map[string]string)
	personMap["name"] = name
	personMap["address"] = address
	fmt.Printf("PersonMap: %v", personMap)

	//encode map to json
	json1, err := json.Marshal(personMap)
	if err != nil {
		panic(err)
	}

	//print out json to screen and file
	fmt.Printf("json: %v\n", string(json1))
	ioutil.WriteFile("tmp.json", json1, 0644)
}
