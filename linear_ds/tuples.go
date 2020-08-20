package main

import(
	"fmt"
)

func h(x int, y int) int{
	return x*y
}

func main(){
	//should print 6
	fmt.Println(h(3,2))
}
