package main

import (
	"fmt"
	"container/list"
)

func main() {
	name := "John"
	lastname := "Doe"
	l := list.New()
	e4 := l.PushBack(lastname)
	e1 := l.PushFront(name)
	l.InsertBefore(3,e4)
	l.InsertAfter(2,e1)
	for e := l.Back(); e!=nil;e=e.Prev() {
		fmt.Println(e.Value)
	}
}
