/*

Keyword	Meaning
range	Iterates over items of a slice
i	Index of current element
v	Value of current element

*/


package main
import "fmt"

func main(){
	xi := []int{42,43,44,45,46,47}
	for i, v:= range xi{
		fmt.Printf("index %v \t value %v\n",i,v)
	}
    // string means key and int mean value
	// map in go and dict in python 
	m := map[string]int{
		"faizal" : 26,
		"sathish" : 29,
	}

	for k,v := range m{
		fmt.Printf("key %v\t value %v\n",k,v)
	}
}