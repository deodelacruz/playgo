//package github.com/deodelacruz

package main

import (
  "fmt" 
  //"reflect"
  //"strconv"
)

func main() {
  //var b bool = true
  var s string = "1"
  //var i int = 4
  //var f float64 = 2.3

  //fmt.Println(s, strconv.ParseInt(s))
}

func sayHello(friend string) string {
   return ("Hello " + friend)
}

func sum(a,b int) int {
  return a+b
}

    // Write your code here
    // get values of 1st diag
        // iterate over each row, max size =n
        // grab elem at these indices 
            // row0: col0
            // row1: col1
            // row2: col2
            // rown: coln 
          // var diag1 = [n]int32 
            
    // collect values of 2nd diag
        //iterate over each row
        // grab elem at these indices
            // row0: coln
            // row1: coln-1
            // row2: coln-2
            // rown: coln-n
    // get sums for each diag
    // get absolute value of difference