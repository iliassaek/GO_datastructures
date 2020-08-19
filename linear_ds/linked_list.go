package main

import(
	"fmt"
)

type Node struct{
	property int
	nextNode *Node
}


func main() {

	var node Node
	node = Node{property: 1}
	fmt.Println(node.property)	
}

