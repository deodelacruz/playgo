/*

Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

*/
package playgo

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	fmt.Println("Hi. What is the name of the text file?")
	currentPath, _ := os.Getwd()
	stdInReader := bufio.NewReader(os.Stdin)
	fileName, _ := stdInReader.ReadString('\n')
	filePath := strings.TrimSuffix(currentPath+string(os.PathSeparator)+fileName, "\n")
	fmt.Printf("file: %v\n", filePath)

	sliceNames := make([]name, 0)
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	fileReader := bufio.NewReader(file)
	for {
		line, _, err := fileReader.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Printf("%s \n", line)
		fields := strings.Fields(string(line))
		nameStruct := name{fields[0], fields[1]}
		sliceNames = append(sliceNames, nameStruct)
	}

	for _, fullName := range sliceNames {
		fmt.Println(fullName)
	}

}

type name struct {
	fname string
	lname string
}
