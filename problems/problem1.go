// Find closest Value in BST

/*
	time complexity: O(log(n))
	Space complexity: O(log(n))
*/
package main

import(
	"math"
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}


func findClosest(bst *BST, closest *int, target int){
	
	
	if bst == nil {
		return
	}
	
	if absolute(*closest - target) > absolute(bst.Value - target){
		*closest= bst.Value
	}
	
	if target > bst.Value{
		findClosest(bst.Right, closest, target)
	}else if target < bst.Value{
		findClosest(bst.Left, closest, target)
	}else{
		*closest= bst.Value
		return
	}
	
	
}



func (tree *BST) FindClosestValue(target int) int {
	// Write your code here.
	var infinity= math.MaxInt64/2
	var closest =&infinity
	fmt.Println(tree.Value)
	findClosest(tree, closest, target)
	return *closest
}

func absolute(x int) int{
	if x > 0 {
		return x
	}else{
		return -x
		
	}
}
