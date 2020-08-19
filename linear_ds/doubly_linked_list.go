package main

import(

	"fmt"
)

type Node struct{
	nextNode *Node
	property int
	previousNode *Node
}

type LinkedList struct{
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(property int){
	var node= &Node{property: property}
	
	//be carful headNode can be nil
	
	if linkedList.headNode != nil{
		node.nextNode= linkedList.headNode
		linkedList.headNode.previousNode= node
	}

	linkedList.headNode= node	
}


//func (linkedList *LinkedList) NodeBetweenValues(firstProperty int,secondProperty int) *Node{
//	return
//}

func (linkedList *LinkedList) IterateList(){
	var node *Node
	for node= linkedList.headNode; node != nil; node= node.nextNode{
		fmt.Println(node.property)
	}
}

//func (linkedList LinkedList) AddToEnd(property int){
//}

func main(){
	fmt.Println("hi")

	var linkedList = &LinkedList{}
	
	linkedList.AddToHead(5)

	linkedList.AddToHead(3)
	linkedList.AddToHead(6)

	// should print 6 3 5
	linkedList.IterateList()

}

