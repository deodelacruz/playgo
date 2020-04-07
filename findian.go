package playgo

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hi. Please input a string, and I'll say 'Found!' if it starts with the character ‘i’,\n" + "ends with the character ‘n’, and contains the character ‘a’.")

	var userInput string
	if _, err := fmt.Scanln(&userInput); err != nil {
		os.Exit(1)
	} else {
		s := strings.ToLower(userInput)
		if strings.HasPrefix(s, "i") && strings.HasSuffix(s, "n") && strings.Contains(s, "a") {
			fmt.Println("Found!")
		} else {
			fmt.Println("Not Found!")
		}
	}
}
