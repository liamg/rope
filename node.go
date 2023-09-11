package rope

var _ Rope = (*Node)(nil)

type Node struct {
	left, right Rope
	weight      int
	lineWeight  int
}

func newNode(l, r Rope) Rope {
	return &Node{
		left:       l,
		right:      r,
		weight:     l.Length(),
		lineWeight: l.NewLineCount(),
	}
}

func (n Node) String() string {
	return n.left.String() + n.right.String()
}

func (n Node) Length() int {
	return n.weight + n.right.Length()
}

func (n Node) Append(r Rope) Rope {
	if n.Length()+r.Length() <= maxLeafSize {
		return newLeaf(append(n.Data(), r.Data()...))
	}
	return newNode(n, r)
}

func (n Node) Prepend(r Rope) Rope {
	if n.Length()+r.Length() <= maxLeafSize {
		return newLeaf(append(r.Data(), n.Data()...))
	}
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
		return newLeaf(nil)
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
	if index := n.left.Index(r); index >= 0 {
		return index
	}
	if index := n.right.Index(r); index >= 0 {
		return n.weight + index
	}
	return -1
}

func (n Node) LastIndex(r rune) int {
	if index := n.right.LastIndex(r); index >= 0 {
		return n.weight + index
	}
	if index := n.left.LastIndex(r); index >= 0 {
		return index
	}
	return -1
}

func (n Node) At(i int) rune {
	if i < n.weight {
		return n.left.At(i)
	}
	return n.right.At(i - n.weight)
}

func (n Node) Line(l int) Rope {
	if l < n.lineWeight {
		return n.left.Line(l)
	} else if l > n.lineWeight {
		return n.right.Line(l - n.lineWeight)
	}
	return n.left.Line(l).Append(n.right.Line(l - n.lineWeight))
}

func (n Node) NewLineCount() int {
	return n.lineWeight + n.right.NewLineCount()
}

func (n Node) Depth() int {
	l := n.left.Depth()
	r := n.right.Depth()
	if l >= r {
		return l + 1
	}
	return r + 1
}

func (n Node) Balance() Rope {
	d := n.Depth()
	if d < len(fibonacci) && fibonacci[d+2] < n.weight {
		return n
	}
	leaves := n.leaves()
	return merge(leaves, 0, len(leaves))
}

func (n Node) leaves() []Rope {
	return append(n.left.leaves(), n.right.leaves()...)
}

func merge(leaves []Rope, start, end int) Rope {
	rng := end - start
	if rng == 1 {
		return leaves[start]
	}
	if rng == 2 {
		return newNode(leaves[start], leaves[start+1])
	}
	mid := start + (rng / 2)
	return newNode(merge(leaves, start, mid), merge(leaves, mid, end))
}

func (n Node) Data() []rune {
	return append(n.left.Data(), n.right.Data()...)
}
