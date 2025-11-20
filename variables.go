/*package main
import "fmt"
func main(){
	a := 5
	fmt.Println(a)
	b,c,d,_,f := 5,6,8,3,4
	fmt.Println(b,c,d,f)
}
*/


// this code will throw error bcz c is not declared
/*
package main
import "fmt"
func main(){
	a,b,c := 2,3,4
	fmt.Println(a,b)
}
	*/


/*
package main
import "fmt"
func main(){
	var a int
	fmt.Println(a)
	a = 10
	fmt.Println(a)
}
*/


/*
%v  normal decimal value
%b  binary representation
%X  hexadecimal (uppercase letters)

\t is an escape sequence in Go (and many other languages) that represents a horizontal tab space.

📌 What \t does

It adds a big space between values so the output becomes neatly aligned in columns.
👍
Example:

fmt.Println("a\tb\tc")


Output:

a       b       c


Each \t inserts a tab space (much bigger than a regular space).

🔍 In your code
fmt.Printf("%v \t %b \t %X\n", a, a, a)


\t separates the decimal, binary, and hexadecimal output in a clear table-like format.

Example output:

7      111      7


Without \t, it would look messy:

7 111 7
*/
// package main
// import "fmt"
// func main(){
// 	a,b,c,d,e := 1,2,3,4,7
// 	fmt.Printf("%v \t %b \t %X\n",a,a,a)
// 	fmt.Printf("%v \t %b \t %X\n",b,b,b)
// 	fmt.Printf("%v \t %b \t %X\n",c,c,c)
// 	fmt.Printf("%v \t %b \t %X\n",d,d,d)
// 	fmt.Printf("%v \t %b \t %X\n",e,e,e)

// }


// package main
// import "fmt"
// func main(){
// 	y:=42
// 	z:= "faizal"
// 	fmt.Printf("%v of the type %T\n",y,y)
// 	fmt.Printf("%v of the type %T\n",z,z)

// 	var m float32 = 43.53
// 	fmt.Printf("%v of the type %T\n",m,m)
// }


// package main
// import "fmt"
// func main(){
// 	y:=42
// 	z:= "faizal"
// 	fmt.Printf("%v of the type %T\n",y,y)
// 	fmt.Printf("%v of the type %T\n",z,z)

// 	var m float32 = 43.53
// 	// z= m
// 	// fmt.Printf("%v of the type %T\n",z,z)
// 	f := float64(m)
// 	fmt.Printf("%v of the type %T\n",f,f)
// }



	
