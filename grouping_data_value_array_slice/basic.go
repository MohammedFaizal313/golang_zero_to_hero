package main
import "fmt"

func main(){

	as := [...]string{"faizz","sam","sathish","gaja","rash"}
	fmt.Println(as)
	fmt.Println(len(as))
	fmt.Printf("%T\n",as)

	for i, v := range as{
		fmt.Printf("index %v\t - value %v\n",i,v)
	}

	fmt.Println(as[0])

	for i:=0;i<len(as);i++{
		fmt.Println(as[i])
	}
// 	a := [3]int{42,43,44}
// 	fmt.Println(a)
//     // ... it will check the length automatically
// 	b := [...]string{"hello","faizzz"}
// 	fmt.Println(b)

// 	fmt.Printf("%T\n",a)
// 	fmt.Printf("%T\n",b)
// }


}