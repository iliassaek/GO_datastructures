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

func (queue *Queue) Add(order *Order){

	//case 1: queue empty
	//case 2.1: queue priority conflict
	//cade 2.2: non queue priority conflict

	//case1
	if len(*queue)==0 {
		*queue= append(*queue, order)
	}else{
		var added bool = false
		var iterativeOrder *Order
		var i int

		for i,iterativeOrder= range *queue{
			if order.priority > iterativeOrder.priority{
				*queue= append((*queue)[:i], append(Queue{order}, (*queue)[i:]...)...)
				added= true
				break
			}
		}
		if !added {
			*queue= append(*queue, order)
		}
	}
}

func (queue *Queue)QueuePrinter(){
	var order *Order
	for _, order= range *queue{
		fmt.Println(*order)
	}
}

func main(){
	var order1= &Order{}
	order1.New(5,20,"pc","iliass1")
	
	var order2= &Order{}
	order2.New(4,10,"pc","iliass2")
	
	var order3= &Order{}
	order3.New(2,67,"pc","iliass3")
	
	var order4= &Order{}
	order4.New(1,9,"pc","iliass4")
	
	var orderToAdd= &Order{}
	orderToAdd.New(3,30,"pc","iliass3")
	
	var queue Queue
	queue= Queue{order1, order2, order3, order4}
	
	queue.Add(orderToAdd)
		
	queue.QueuePrinter()

}
