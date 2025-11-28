package main
import "fmt"

func main(){
	xi := make([]int,0,10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	xi = append(xi,0,1,2,3,4,5,6,7,8,9,11)
	fmt.Println(xi)
	fmt.Println(cap(xi))
	fmt.Println("-----------------------------")
	// When a slice grows beyond its capacity, Go creates a new underlying array
	// with larger capacity, usually doubling the size.
	xi = append(xi,11,2,13,14,15,16,17,18,19,29)
	fmt.Println(len(xi))
	fmt.Printf("xi - %T",xi)
	fmt.Println(cap(xi))
	
}