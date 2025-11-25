package main
import (
  "math/rand" 
  "fmt"
)
func main(){
	x := rand.Intn(10)
  y := rand.Intn(10)
  fmt.Printf("the value of x & y is %v and %v\t",x,y)

//   if x < 4 && y < 4{
//     fmt.Println("both are less then 4")
//   }else if x > 6 && y > 6{
//     fmt.Println("both are greater then 6")
//   }else if x >=4 && x<=6{
//     fmt.Println("x is from 4 to 6 inclusive of both numbers")
//   }else if   y!=5{
//     fmt.Println("y is not euqal to 5")
//   }else {
//     fmt.Println("none of the previous were met")
//   } 
// }
  switch{
  case x < 4 && y < 4:
    fmt.Println("both are less then 4")
  case x > 6 && y > 6:
    fmt.Println("both are greater then 6")
  case x >=4 && x<=6:
    fmt.Println("x is from 4 to 6 inclusive of both numbers")
  case y!=5:
    fmt.Println("y is not equal to 5")
  default:
    fmt.Println("none of the previous were met")
  }
}