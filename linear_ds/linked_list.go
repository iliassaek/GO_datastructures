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

//AddAfter method adds a node with nodeProperty after node with property

func (linkedList *LinkedList) AddAfter(nodeProperty int,property int){
	var nodeToAdd *Node
	nodeToAdd = &Node{property: nodeProperty}
	
	var targetNode *Node
	targetNode= linkedList.NodeWithValue(property)
	
	if targetNode != nil {
		nodeToAdd.nextNode= targetNode.nextNode
		targetNode.nextNode= nodeToAdd
	}
}

func main() {

	var linkedList LinkedList
	linkedList.AddToHead(1)
	linkedList.AddToHead(2)
	linkedList.AddToHead(3)

	linkedList.AddAfter(20,2)
	linkedList.AddAfter(50,20)	

	//should return 3 2 20 50 1
	linkedList.IterateList()
}
