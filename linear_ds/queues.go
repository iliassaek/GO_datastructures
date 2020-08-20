package main

import(
	"fmt"
)
// array of Order addresses
type Queue []*Order

type Order struct{
	priority int
	quantity int
	product string
	customerName string
}

func (order *Order) New(priority int, quantity int, product string, customerName string ){
	order.priority= 	priority
	order.quantity=		quantity
	order.product=		product
	order.customerName=	customerName	
}

func main(){
	var order= &Order{}
	order.New(1,2,"pc","iliass")
	
	//should print &{1 2 pc iliass}
	fmt.Println(order)

}
