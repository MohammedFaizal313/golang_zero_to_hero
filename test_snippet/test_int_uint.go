// package main
// import (
// 	"fmt"
// 	"golang_pract/03-scopes"
// 	)
// var x = 42
// func main(){

// 	var i int  = 7
// 	var j uint = 29
// 	scopes.PrintMe(x)
// 	fmt.Printf("%v is the type of %T\n",i,i)
// 	fmt.Printf("%v is the type of %T\n",j,j)
// }


package main
import (
	// "fmt"
	"golang_pract/03-scopes"
)
var name = "faizal"
func main(){
	scopes.PrintName(name)
}
