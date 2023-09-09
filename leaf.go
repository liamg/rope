package rope

import "strings"

var _ Rope = (*Leaf)(nil)

type Leaf struct {
	data     []rune
	leafSize int
}

func (l Leaf) String() string {
	return string(l.data)
}

func (l Leaf) Length() int {
	return len(l.data)
}

func (l Leaf) Append(n Rope) Rope {
	return newNode(l, n)
}

func (l Leaf) Prepend(n Rope) Rope {
	return newNode(n, l)
}

func (l Leaf) Split(at int) (Rope, Rope) {
	return &Leaf{
			data: l.data[:at],
		}, &Leaf{
			data: l.data[at:],
		}
}

func (l Leaf) Sub(start, end int) Rope {
	if start < 0 {
		start = 0
	}
	if end > len(l.data) {
		end = len(l.data)
	}
	return &Leaf{
		data: l.data[start:end],
	}
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
	vl := 1
	var start, end int
	for i, r := range l.data {
		if r == '\n' {
			vl++
			if vl == line {
				start = i + 1
			} else if vl == line+1 {
				end = i
				return &Leaf{
					data: l.data[start:end],
				}
			}
		}
	}
	if start == 0 && line == 1 {
		return &l
	}
	return &Leaf{
		data: l.data[start:],
	}
}

func (l Leaf) Balance() Rope {
	if l.leafSize == 0 || len(l.data) <= l.leafSize {
		return l
	}
	return newNode(
		(&Leaf{data: l.data[:l.leafSize]}).Balance(),
		(&Leaf{data: l.data[l.leafSize:]}).Balance(),
	)
}
