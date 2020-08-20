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

func (set *Set) Intersect(anotherSet *Set) *Set{
	var intersectSet= &Set{}
	intersectSet.New()
	var value int	
	for value, _ = range set.integerMap {
		if anotherSet.ContainsElement(value){
			intersectSet.AddElement(value)
		}
	}
	return intersectSet
}

func main(){

	var set = &Set{}
	set.New()
	set.AddElement(2)
	set.AddElement(3)	

	var anotherSet= &Set{}
	anotherSet.New()
	anotherSet.AddElement(1)
	anotherSet.AddElement(2)

	var interSet = set.Intersect(anotherSet)	
	//should return &{map[2:true]}
	fmt.Println(interSet)
}
