package rope

var _ Rope = (*Node)(nil)

type Node struct {
	left, right Rope
	weight      int
}

func newNode(l, r Rope) Rope {
	// TODO: calculate weights, lines and stuff here
	return &Node{
		left:   l,
		right:  r,
		weight: l.Length(),
	}
}

func (n Node) String() string {
	return n.left.String() + n.right.String()
}

func (n Node) Length() int {
	return n.left.Length() + n.right.Length()
}

func (n Node) Append(r Rope) Rope {
	return newNode(n, r)
}

func (n Node) Prepend(r Rope) Rope {
	return newNode(r, n)
}

func (n Node) Split(at int) (Rope, Rope) {
	if at < n.weight {
		// split left
		left, right := n.left.Split(at)
		return left, newNode(right, n.right)
	} else if at > n.weight {
		// split right
		left, right := n.right.Split(at - n.weight)
		return newNode(n.left, left), right
	} else {
		// split here
		return n.left, n.right
	}
}

func (n Node) Sub(start, end int) Rope {
	if start < 0 {
		start = 0
	}
	if end > n.Length() {
		end = n.Length()
	}
	if start >= end {
		return &Leaf{}
	}
	if start < n.weight && end < n.weight {
		// sub left
		return n.left.Sub(start, end)
	} else if start > n.weight && end > n.weight {
		// sub right
		return n.right.Sub(start-n.weight, end-n.weight)
	} else {
		// sub both
		left := n.left.Sub(start, n.weight)
		right := n.right.Sub(0, end-n.weight)
		return newNode(left, right)
	}
}

func (n Node) Index(r rune) int {
	//TODO implement me
	panic("implement me")
}

func (n Node) LastIndex(r rune) int {
	//TODO implement me
	panic("implement me")
}

func (n Node) At(i int) rune {
	//TODO implement me
	panic("implement me")
}

func (n Node) Balance() Rope {
	//TODO implement me
	panic("implement me")
}
