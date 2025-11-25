
package main
import "fmt"
func main(){
	y := 0
	for i :=0; i<=5; i++{
		fmt.Printf("counting %v\n",i)
	}

	for y<10{
		fmt.Printf("y i %v \t\t\t second for loop\n",y)
		y++
	}

	for {
		
        if y>20{
			
			break
		}
		fmt.Printf("y is %v \t\t third for loop\n",y)
		y++

	}

	for i:=0; i<20; i++{
		if i%2!=0{
			continue
		}
		fmt.Println("counting even numvers:",i)
	}

}

