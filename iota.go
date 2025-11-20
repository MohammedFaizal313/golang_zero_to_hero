/*


First Part — iota and constants
const(
    _ = iota
    a
    b
    c
    d
    e
    f
)

Name	Value
_	0 (ignored)
a	1
b	2
c	3
d	4
e	5
f	6

So:

a = 1
b = 2
c = 3
d = 4
e = 5
f = 6


We skip 0 using _ because we don't need it.

🔹 Second Part — Bit shifting (1 << a)

1 << a means shift the binary value of 1 to the left a times

Binary of 1:

1 = 00000001


Shifting left once (1 << 1):

00000001 → 00000010 = 2


Shifting left twice (1 << 2):

00000001 → 00000100 = 4


Shifting left 3 times (1 << 3):

00000001 → 00001000 = 8


So you get powers of 2:

Expression	Decimal	Binary
1	1	1
1 << a → 1 << 1	2	10
1 << b → 1 << 2	4	100
1 << c → 1 << 3	8	1000
1 << d → 1 << 4	16	10000
1 << e → 1 << 5	32	100000
1 << f → 1 << 6	64	1000000
*/


// package main
// import "fmt"

// const(
// 	c0 = iota // c0 == 0
// 	c1 = iota
// 	c2 = iota
// )

// const(
// 	c3 = iota
// 	c4
// 	c5
// 	c6
// )

// func main(){
// 	fmt.Println(c0,c1,c2)
// 	fmt.Println(c3,c4,c5,c6)
// 	fmt.Printf("%d \t %b\n",1<<1,1<<1)
// 	fmt.Printf("%d \t %b\n",1<<2,1<<2)
// 	fmt.Printf("%d \t %b\n",1<<3,1<<3)
// 	fmt.Printf("%d \t %b\n",1<<4,1<<4)
// 	fmt.Printf("%d \t %b\n",1<<5,1<<5)
// 	fmt.Printf("%d \t %b\n",1<<6,1<<6)
// }


// package main
// import "fmt"

// const(
// 	_ = iota
// 	a
// 	b
// 	c
// 	d
// 	e
// 	f
// )

// func main(){
// 	fmt.Printf("%d \t %b \n",1,1)
// 	fmt.Printf("%d \t %b\n",1<<a,1<<a)
// 	fmt.Printf("%d \t %b\n",1<<b,1<<b)
// 	fmt.Printf("%d \t %b\n",1<<c,1<<c)
// 	fmt.Printf("%d \t %b\n",1<<d,1<<d)
// 	fmt.Printf("%d \t %b\n",1<<e,1<<e)
// 	fmt.Printf("%d \t %b\n",1<<f,1<<f)
// }



// package main
// import "fmt"

// type ByteSize int

// const(
// 	_ = iota
// 	KB  ByteSize = 1 << (10 * iota)
// 	MB
// 	GB
// 	TB
// 	PB
// 	EB
// )

// func main(){
// 	fmt.Printf("%d \t\t\t %b\n",KB,KB)
// 	fmt.Printf("%d \t\t %b\n",MB,MB)
// 	fmt.Printf("%d \t\t %b\n",GB,GB)
// 	fmt.Printf("%d \t\t %b\n",TB,TB)
// 	fmt.Printf("%d \t %b\n",PB,PB)
//     fmt.Printf("%d \t %b\n",EB,EB)
// }























