package main

import(
	"fmt"
)


type Order struct{
	priority int
	quantity int
	product string
	customerName string
}

func main(){
	var order= &Order{priority:1, quantity:2, product:"pc", customerName:"iliass"}
	
	//should print &{1 2 pc iliass}
	fmt.Println(order)

}
