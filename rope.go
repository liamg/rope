package rope

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Rope is a data structure for efficiently storing and manipulating UTF-8 text.
// It is a binary tree where each leaf node contains a string of runes.
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
// NewLineCount returns the number of new lines in the tree.
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
	NewLineCount() int
	Balance() Rope
	Depth() int

	leaves() []Rope
}

// FromFile reads a file into a rope.
func FromFile(path string) (Rope, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()
	r, err := FromReader(f)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return r, nil
}

// FromReader reads a reader into a rope.
func FromReader(r io.Reader) (Rope, error) {
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
		data: data,
	}, nil
}

// FromString creates a rope from a string.
func FromString(s string) Rope {
	return (&Leaf{
		data: []rune(s),
	}).Balance()
}

// FromRunes creates a rope from a slice of runes.
func FromRune(r rune) Rope {
	return &Leaf{
		data: []rune{r},
	}
}
