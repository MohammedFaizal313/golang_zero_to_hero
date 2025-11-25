package main
import "fmt"
func main(){

	xi := []int{42,43,44,45,46}
	for i,v := range xi{
		fmt.Println("ranging over a slice ",i,v)

	}

	m:= map[string]int{
		"faizz" :26,
		"satish" : 27,
	}

	for k,v := range m{
		fmt.Println("ranging over a map", k, v)
	}

}