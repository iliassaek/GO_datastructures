package main

import(
	"fmt"
)

func h(x int, y int) int{
	return x*y
}

func g(l int, m int) (x int, y int){
	x= l*2
	y= m*8
	return
}

func main(){
	//should print 6,16
	fmt.Println(g(3,2))
}
