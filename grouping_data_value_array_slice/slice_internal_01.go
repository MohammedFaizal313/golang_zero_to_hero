package main
import "fmt"

func main(){
	a := []int{0,1,2,3,4,5}
	b := a
	fmt.Println(a)
	fmt.Println(b)

	a[0] = 7

	fmt.Println("a - ",a)
	fmt.Println("b - ",b)
}

/*
to overcome above code 
*/

// package main
// import "fmt"
// func main(){
// 	a := []int{0,1,2,3,4,5}
// 	b := make([]int,6)
// 	copy(b,a)

// 	fmt.Println(a)
// 	fmt.Println(b)

// 	a[0] = 7
// 	 fmt.Println("a - ",a)
// 	 fmt.Println("b - ",b)
// }