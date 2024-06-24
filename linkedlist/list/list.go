package list

import "fmt"

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func NewList[T any]() List[T] {
	return List[T]{Size: 0}
}

func (l *List[T]) AddFirst(value T) {
	if l.Size == 0 {
		node := Node[T]{Value: value}
		l.Head = &node
		l.Tail = &node
		l.Size = 1
		return
	}

	newNode := Node[T]{Value: value}
	l.Head.AddPrev(&newNode)
	l.Head = &newNode
	l.Size++
}

func (l *List[T]) AddLast(value T) {
	if l.Size == 0 {
		node := Node[T]{Value: value}
		l.Head = &node
		l.Tail = &node
		l.Size = 1
		return
	}

	newNode := Node[T]{Value: value}
	l.Tail.AddNext(&newNode)
	l.Tail = &newNode
	l.Size++
}

func (l *List[T]) RemoveFirst() {
	if l.Head == nil {
		return
	}
	newHead := l.Head.Next
	l.Head.Remove()
	l.Head = newHead
	l.Size--
}

func (l *List[T]) RemoveLast() {
	if l.Tail == nil {
		return
	}
	newTail := l.Tail.Prev
	l.Tail.Remove()
	l.Tail = newTail
	l.Size--
}

func (l *List[T]) FirstValue() T {
	if l.Size == 0 {
		var zero T
		return zero
	}
	return l.Head.Value
}

func (l *List[T]) LastValue() T {
	if l.Size == 0 {
		var zero T
		return zero
	}
	return l.Tail.Value
}

func (l *List[T]) Print() {
	if l.Size == 0 {
		fmt.Println("empty list")
		return
	}

	for n := l.Head; n != nil; n = n.Next {
		fmt.Printf("%v -> ", n.Value)
	}
	fmt.Print("nil\n")
}

func (l *List[T]) PrintReversed() {
	if l.Size == 0 {
		fmt.Println("empty list")
		return
	}

	for n := l.Tail; n != nil; n = n.Prev {
		fmt.Printf("%v -> ", n.Value)
	}
	fmt.Print("nil\n")
}
