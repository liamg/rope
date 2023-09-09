package rope

import (
	"bufio"
	"io"
)

// Rope is a data structure for efficiently storing and manipulating very long strings.
// It is a binary tree where each leaf node contains a string of runes.
// The data type used is rune, meaning that it can store any Unicode code point.
// The length of the string is the sum of the lengths of the strings in the leaf nodes.
// The tree is balanced so that the length of the string in each leaf node is roughly the same.
// This allows for fast indexing and slicing operations.
// The tree is immutable, so any modification operations will return a new tree.
//
// Split the tree into two trees at a given index.
//
// Sub reads a substring into the tree using the provided start and end indexes.
//
// Append the provided tree to the end of the current tree.
//
// Prepend the provided tree to the start of the current tree.
//
// Index returns the first index of a given rune.
//
// LastIndex returns the last index of a given rune.
//
// At returns the rune at a given index.
//
// String returns the string representation of the tree.
//
// Length returns the length of the string.
//
// Line returns the line at the given (zero-based) index.
//
// Balance the tree, using the leaf size as the maximum length of a leaf node.
type Rope interface {
	String() string
	Length() int
	At(int) rune
	Append(Rope) Rope
	Prepend(Rope) Rope
	Split(int) (Rope, Rope)
	Sub(int, int) Rope
	Index(rune) int
	LastIndex(rune) int
	Line(int) Rope
	Balance() Rope
}

func New(r io.Reader, leafSize int) (Rope, error) {
	// pessimistically allocate enough room for 1:1 ratio of byte:rune
	data := make([]rune, 0, bufio.NewReader(r).Size())
	for {
		s, err := bufio.NewReader(r).ReadString(0)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		data = append(data, []rune(s)...)
	}
	return &Leaf{
		leafSize: leafSize,
		data:     data,
	}, nil
}

func NewFromString(s string, leafSize int) Rope {
	return (&Leaf{
		data:     []rune(s),
		leafSize: leafSize,
	}).Balance()
}

func NewFromRune(r rune, leafSize int) Rope {
	return &Leaf{
		data:     []rune{r},
		leafSize: leafSize,
	}
}
