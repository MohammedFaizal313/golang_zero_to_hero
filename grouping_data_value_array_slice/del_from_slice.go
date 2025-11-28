package main
import "fmt"

func main(){
	xi := []int{0,1,2,3,4,5,6,7,8,9,10}
	fmt.Printf("xi - %#v\n",xi)
	fmt.Println("-------------------------")
    
	// ...  tghis tells Go to unpack the slice, converting
	xi = append(xi[:4],xi[5:]...)
	fmt.Printf("xi - %#v\n",xi)
}