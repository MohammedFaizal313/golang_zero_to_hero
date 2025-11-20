/*
A functions can take zero or more arguments
in this example add takes two parameter as a int
Notice that type comes after the variable name
*/

// package main
// import "fmt"

// func add(x,y int) int{
// 	return x+y
// }

// func main(){
// 	fmt.Println(add(4,5))
// 	say_hello()

// }

// func say_hello(){
// 	fmt.Println("hello")
// }


// package main
// import "fmt"

// func swap(a string,b string)(string, string){
// 	return b,a
// }

// func main(){
// 	a,b := swap("hello","world")
// 	fmt.Println(a,b)

// }


// package main
// import "fmt"

// const d int = 42
// var c, python, java bool
// func main(){
// 	var i ,j float32 = 4,2
// 	fmt.Println(i,j, c, python,java, d)
// }


// package main
// import "fmt"

// const d int = 42
// var c, python, java bool
// func main(){
// 	var i ,j float32 = 4,2
// 	// fmt.Printf("%v \t %v \t %v \t %v \t %v \t %v \t",i,j, c, python,java, d)
// 	fmt.Printf("%T \t %T \t %T \t %T \t %T \t %T \t",i,j, c, python,java, d)
// }


// package main
// import "fmt"

// const d int = 42
// var c, python, java bool

// func main(){
// 	k:=3
// 	var i ,j float32 = 4,2
// 	fmt.Println(i, j, k, c, python,java, d)
// }


// package main
// import(
// 	"fmt"
// 	"math"
// )
// func main(){
// 	var x,y float32 = 3,4
// 	var f float64 = math.Sqrt(float64(x*x + y*y))    //because math.Sqrt() only accepts float64
// 	var z uint = uint(f)    // - It is a whole number - It cannot be negative
// 	fmt.Println(x,y,z)
// }




// package main
// import "fmt"

// func main(){
// 	const x float32 = 25   #constant cannot be declared using := syntax
// 	fmt.Println(x)
// }



// need to go throw this code deeper
package main
import "fmt"

const(
	Big = 1 << 100
	small = Big >> 99
)

func needInt(x int) int { return x*10 + 1}
func needFloat(x float64) float64{
	return x * 0.1
}

func main(){
	fmt.Println(needInt(small))
	fmt.Println(needFloat(small))
	fmt.Println(needFloat(Big))
}