package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println("Done.")
	f.Write([]byte("Using Write function. \n"))
	f.WriteString("Using Writestring function.")

	message := []byte("Look!")
	err = ioutil.WriteFile("test2.txt", message, 0644)

	content1, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("File contents: ")
	fmt.Println(string(content1))
}
