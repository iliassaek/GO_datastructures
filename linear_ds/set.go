package main

import(
	"fmt"
)

type Set struct{
	integerMap map[int]bool
}

func (set *Set) New(){
	set.integerMap = make(map[int]bool)
}

func (set *Set) ContainsElement(element int) bool{
	var exists bool
	_, exists = set.integerMap[element]
	return exists
}

func main(){

	var set = &Set{}
	set.New()
	set.integerMap[2]=true
	
	//should return true
	fmt.Println(set.ContainsElement(2))

	//sould return false
	fmt.Println(set.ContainsElement(1))

}
