package main

import "golang-practice/linkedlist/list"

func main() {
	l := list.NewList[int]()
	l.AddFirst(11)
	l.AddFirst(12)
	l.AddFirst(13)
	l.AddFirst(14)

	l.AddLast(15)
	l.AddLast(16)
	l.AddLast(17)
	l.AddLast(18)

	l.RemoveFirst()
	l.RemoveLast()

	l.Print()
	l.PrintReversed()
}
