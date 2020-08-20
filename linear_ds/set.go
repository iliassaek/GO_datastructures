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

func (set *Set) AddElement(element int){
	if(!set.ContainsElement(element)){
		set.integerMap[element]= true
	}
}


func (set *Set) DeleteElement(element int){
	delete(set.integerMap, element)
}

func main(){

	var set = &Set{}
	set.New()
	set.AddElement(2)
	
	//should return true
	fmt.Println(set.ContainsElement(2))

	//should return false
	set.DeleteElement(2)
	fmt.Println(set.ContainsElement(2))

}
