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


// returns the last node of the linked list
func (linkedList LinkedList) LastNode() *Node{
	var node *Node
	var lastNode *Node
	
	for node = linkedList.headNode; node != nil; node= node.nextNode{
		if node.nextNode == nil {
			lastNode= node
		}
	}
	return lastNode
}

//returns the node with the input value
func (linkedList *LinkedList) NodeWithValue(property int) *Node{
	var node *Node 
	var nodeWithValue *Node
	for node= linkedList.headNode; node != nil; node= node.nextNode{
		if node.property == property {
			nodeWithValue= node
			break;
		}
	}
	return nodeWithValue
}

func main() {

	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)
	//should print 1
	
	var nodeWithValue *Node
	nodeWithValue= linkedList.NodeWithValue(2)
	if nodeWithValue == nil {
		fmt.Println("not found")
	}else{
		//should return 2
		fmt.Println(nodeWithValue.property)
	}
}

