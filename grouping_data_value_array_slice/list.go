package main
import "fmt"

func main(){
	a := []int{23,24,25}
	fmt.Println(a)
	fmt.Println("-------------------")

	a = append(a,26,27,28)
	fmt.Println(a)
    fmt.Println("--------------------")
    
	// [inclusive:exclusive]
	fmt.Printf("a - %v\n",a[0:5])
    
	// [:exclusive]
	fmt.Printf("a - %#v\n",a[:6])

	// [inclusive:]
	fmt.Printf("a - %#v\n",a[2:])

	fmt.Printf("a - %#v\n",a[:])
}