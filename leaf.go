package rope

import "strings"

var _ Rope = (*Leaf)(nil)

const maxLeafSize = 256

type Leaf struct {
	data []rune
}

func newLeaf(data []rune) Rope {
	return &Leaf{
		data: data,
	}
}

func (l Leaf) String() string {
	return string(l.data)
}

func (l Leaf) Length() int {
	return len(l.data)
}

func (l Leaf) Append(n Rope) Rope {
	if l.Length()+n.Length() <= maxLeafSize {
		return newLeaf(append(l.data, n.Data()...))
	}
	return newNode(l, n)
}

func (l Leaf) Prepend(n Rope) Rope {
	if l.Length()+n.Length() <= maxLeafSize {
		return newLeaf(append(n.Data(), l.data...))
	}
	return newNode(n, l)
}

func (l Leaf) Split(at int) (Rope, Rope) {
	if at < 0 {
		at = 0
	}
	if at > len(l.data) {
		at = len(l.data)
	}
	return newLeaf(l.data[:at]), newLeaf(l.data[at:])
}

func (l Leaf) Sub(start, end int) Rope {
	if start < 0 {
		start = 0
	}
	if end > len(l.data) {
		end = len(l.data)
	}
	return newLeaf(l.data[start:end])
}

func (l Leaf) Index(r rune) int {
	return strings.IndexRune(l.String(), r)
}

func (l Leaf) LastIndex(r rune) int {
	return strings.LastIndex(l.String(), string(r))
}

func (l Leaf) At(i int) rune {
	if i < 0 || i >= len(l.data) {
		return -1
	}
	return l.data[i]
}

func (l Leaf) Line(line int) Rope {
	if line < 0 || line > l.NewLineCount() {
		return newLeaf(nil)
	}
	vl := 0
	var start, end int
	for i, r := range l.data {
		if r == '\n' {
			vl++
			if vl == line {
				start = i + 1
			} else if vl == line+1 {
				end = i
				return newLeaf(l.data[start:end])
			}
		}
	}
	if start == 0 && line == 1 {
		return &l
	}
	return newLeaf(l.data[start:])
}

func (l Leaf) Balance() Rope {
	return l
}

func (l Leaf) NewLineCount() int {
	if len(l.data) == 0 {
		return 0
	}
	return strings.Count(l.String(), "\n")

}

func (l Leaf) Depth() int {
	return 1
}

func (l Leaf) leaves() []Rope {
	return []Rope{l}
}

func (l Leaf) Data() []rune {
	return l.data
}
