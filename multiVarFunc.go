/* Call a function that rturns multiple values
 */
package main

import (
	"fmt"
	"time"
)

func getConfig() (bool, string, time.Time) {
	return false, "hello", time.Now()
}

func main() {

	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)

	Debug, ReturnedString, returnedTime := getConfig()
	fmt.Println(Debug, ReturnedString, returnedTime)

}
