/*
create a program that has a loopthat prints every number from 0 to 90
modify the program  from the previous hands on exercise to run 100 times
*/

package main
import (
  "math/rand" 
  "fmt"
)
func main(){
	// for i:=0; i < 100; i++{
	// 	fmt.Println(i)
	// }
	for i :=0; i < 100; i++{
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Printf("iteration %v\t : the value of x & y is %v and %v\t",i,x,y)

		switch{
		case x < 4 && y < 4:
			fmt.Println("both are less then 4")
		case x > 6 && y > 6:
			fmt.Println("both are greater then 6")
		case x >=4 && x<=6:
			fmt.Println("x is from 4 to 6 inclusive of both numbers")
		case y!=5:
			fmt.Println("y is not equal to 5")
		default:
			fmt.Println("none of the previous were met")
		}
	}
}

