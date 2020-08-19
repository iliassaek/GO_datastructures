package main

import(
	"fmt"
)

// create basic node
type Node struct{
	property int
	nextNode *Node
}

// create the head node
type LinkedList struct{
	headNode *Node
}

func (linkedList *LinkedList)AddToHead(property int){

	var node *Node
	node = &Node{property: property,nextNode: nil}
	
	if linkedList.headNode != nil {
		node.nextNode= linkedList.headNode
	}
	
	//in all cases
	linkedList.headNode= node
}

func (linkedList *LinkedList)IterateList(){
	var node *Node
	
	for node= linkedList.headNode; node != nil; node = node.nextNode{
		fmt.Println(node.property)
	}  
}

func main() {

	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)

	//should print 2 1
	linkedList.IterateList()
}

