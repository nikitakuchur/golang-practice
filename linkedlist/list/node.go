package list

type Node[T any] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T
}

func (n *Node[T]) AddPrev(node *Node[T]) {
	head := n.Prev

	n.Prev = node
	node.Next = n

	node.Prev = head

	if head != nil {
		head.Next = node
	}
}

func (n *Node[T]) AddNext(node *Node[T]) {
	tail := n.Next

	n.Next = node
	node.Prev = n

	node.Next = tail

	if tail != nil {
		tail.Prev = node
	}
}

func (n *Node[T]) Remove() {
	prev := n.Prev
	next := n.Next

	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
}
