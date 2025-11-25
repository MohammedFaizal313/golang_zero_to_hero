// package main
// import "fmt"
// func main(){
//     xi := []int{42,43,44,45,46,47}
// 	for i, v:= range xi{
// 		fmt.Printf("index %v \t value %v\n",i,v)
// 	}
    
// 	m := map[string]int{
// 		"faizal" : 26,
// 		"satish" : 29,
// 	}

// 	for k,v := range m{
// 		fmt.Printf("key %v\t value %v\n",k,v)
// 	}

// 	fmt.Println("--------------------------------------------------------")

// 	age1 := m["faizal"]
// 	fmt.Println("the age of bond",age1)

// 	if v,ok := m["faizal"]; !ok{
// 		fmt.Println("there is a bond  lookup entry, and the bond age is ",v)
// 	}

// 	age2 := m["Q"]
// 	fmt.Println("the  age of bond", age2)

// 	if v,ok := m["Q"]; !ok{
// 		fmt.Println("there i no Q, and here is the zero value of an int",v)
// 	}
// }

package main
import (
    "fmt"
	"math/rand"
)

func main(){
	c :=0
	for i:=0;i<100;i++{

		if x:= rand.Intn(5); x==3{
			fmt.Printf("Iteration %v\t total count %v\t x is %v\n",i,c,x)
			c++
		}
			
	}
}