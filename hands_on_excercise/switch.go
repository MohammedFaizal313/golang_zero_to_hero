/*
create a random int between 0 inclusive and 5 exclusive
 -> assign the value to a variable with the identifier x
use a switch statementsto print descriptionvof the variables and value
run the code 42 times and print the iteration number
*/

package main

import(
	"fmt"
	"math/rand"
)

func main(){
	for i:=0; i<42; i++{
		x := rand.Intn(5)
		fmt.Printf("iteration %v\t:",i)

		switch x{
		case 0:
			fmt.Println("x is 0")
		case 1:
			fmt.Println("x is 1")
		case 2:
			fmt.Println("x is 2")
		case 3:
			fmt.Println("x is 3")
		case 4:
			fmt.Println("x is 4")
		}
	}
}