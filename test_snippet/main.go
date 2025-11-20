/*
var for zero value
short declaration operator
multiple intializations
var when specifically is required
blank identifier
*/

package main
import "fmt"

func main(){
	var a int
	fmt.Println(a)

	b := 3
	fmt.Println(b)

	c,d := 4,5
	fmt.Println(c,d)

	var f float32 = 4.5
	fmt.Printf("%v is the type of %T\n",f,f)

	m,_,o := 4,5,6
	fmt.Println(m,o)
}