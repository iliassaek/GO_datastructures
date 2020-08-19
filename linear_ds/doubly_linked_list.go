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

func (linkedList *LinkedList) LastNode() *Node {
	
	var lastNode *Node
	var node *Node

	for node= linkedList.headNode; node != nil; node= node.nextNode{
		if node.nextNode == nil{
			lastNode= node
		}		
	}
	
	return lastNode
}

func (linkedList *LinkedList) NodeBetweenValues(firstProperty int,secondProperty int) *Node{
	var betweenNode *Node
	var node *Node
	for node= linkedList.headNode; node.nextNode != nil; node= node.nextNode{
		if (node.previousNode != nil && node.previousNode.property == firstProperty && node.nextNode.property == secondProperty){
			betweenNode= node
		}
	}

	return betweenNode
}

func (linkedList *LinkedList) IterateList(){
	var node *Node
	for node= linkedList.headNode; node != nil; node= node.nextNode{
		fmt.Println(node.property)
	}
}

func (linkedList *LinkedList) AddToEnd(property int){
	var nodeToAdd = &Node{property: property}
	var lastNode= linkedList.LastNode()

	if lastNode != nil {
		lastNode.nextNode= nodeToAdd
		nodeToAdd.previousNode= lastNode
	}else{
		linkedList.headNode= nodeToAdd
		fmt.Println("last Node nil")
	}
	
}


func main(){
	fmt.Println("hi")

	var linkedList = &LinkedList{}
	
	linkedList.AddToHead(5)

	linkedList.AddToHead(3)
	linkedList.AddToHead(6)
	linkedList.AddToEnd(7)
	
	// should print 5
	fmt.Println(linkedList.NodeBetweenValues(3,7).property)
}

