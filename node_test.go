package rope

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode_Append(t *testing.T) {
	tests := []struct {
		name    string
		start   *Node
		appends []string
		want    string
	}{
		{
			name:    "append to empty",
			start:   newNode(FromString(""), FromString("")).(*Node),
			appends: []string{"g", "h", "i"},
			want:    "ghi",
		},
		{
			name:    "append to double-leaf",
			start:   newNode(FromString("abc"), FromString("def")).(*Node),
			appends: []string{"g", "h", "i"},
			want:    "abcdefghi",
		},
		{
			name:    "append to one-leaf, one-branch",
			start:   newNode(FromString("abc"), FromString("def").Append(FromString("ghi"))).(*Node),
			appends: []string{"j", "k", "l"},
			want:    "abcdefghijkl",
		},
		{
			name:    "append to one-branch, one-leaf",
			start:   newNode(FromString("abc").Append(FromString("def")), FromString("ghi")).(*Node),
			appends: []string{"j", "k", "l"},
			want:    "abcdefghijkl",
		},
		{
			name:    "append to double-branch",
			start:   newNode(FromString("abc").Append(FromString("def")), FromString("ghi").Append(FromString("jkl"))).(*Node),
			appends: []string{"m", "n", "o"},
			want:    "abcdefghijklmno",
		},
		{
			name:    "append to oversize",
			start:   newNode(FromString(strings.Repeat("a", maxLeafSize)), FromString(strings.Repeat("b", maxLeafSize))).(*Node),
			appends: []string{"c", "d", "e"},
			want:    strings.Repeat("a", maxLeafSize) + strings.Repeat("b", maxLeafSize) + "cde",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n Rope = tt.start
			for _, a := range tt.appends {
				n = n.Append(newLeaf([]rune(a)))
			}
			assert.Equalf(t, tt.want, n.String(), "Append(%v)", tt.appends)
		})
	}
}

func TestNode_At(t *testing.T) {
	tests := []struct {
		name  string
		left  string
		right string
		at    int
		want  rune
	}{
		{
			name:  "empty",
			left:  "",
			right: "",
			at:    0,
			want:  -1,
		},
		{
			name:  "first",
			left:  "abc",
			right: "def",
			at:    0,
			want:  'a',
		},
		{
			name:  "middle left",
			left:  "abc",
			right: "def",
			at:    1,
			want:  'b',
		},
		{
			name:  "first right",
			left:  "abc",
			right: "def",
			at:    3,
			want:  'd',
		},
		{
			name:  "last",
			left:  "abc",
			right: "def",
			at:    5,
			want:  'f',
		},
		{
			name:  "out of range -1",
			left:  "abc",
			right: "def",
			at:    -1,
			want:  -1,
		},
		{
			name:  "out of range 6",
			left:  "abc",
			right: "def",
			at:    6,
			want:  -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNode(FromString(tt.left), FromString(tt.right))
			assert.Equalf(t, tt.want, n.At(tt.at), "At(%v)", tt.at)
		})
	}
}

func TestNode_Balance(t *testing.T) {
	type fields struct {
		left       Rope
		right      Rope
		weight     int
		lineWeight int
	}
	tests := []struct {
		name   string
		fields fields
		want   Rope
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				left:       tt.fields.left,
				right:      tt.fields.right,
				weight:     tt.fields.weight,
				lineWeight: tt.fields.lineWeight,
			}
			assert.Equalf(t, tt.want, n.Balance(), "Balance()")
		})
	}
}

func TestNode_Data(t *testing.T) {
	tests := []struct {
		name  string
		left  []rune
		right []rune
		want  []rune
	}{
		{
			name:  "empty",
			left:  []rune{},
			right: []rune{},
			want:  []rune{},
		},
		{
			name:  "left",
			left:  []rune("abc"),
			right: []rune{},
			want:  []rune("abc"),
		},
		{
			name:  "right",
			left:  []rune{},
			right: []rune("def"),
			want:  []rune("def"),
		},
		{
			name:  "both",
			left:  []rune("abc"),
			right: []rune("def"),
			want:  []rune("abcdef"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := newNode(FromString(string(tt.left)), FromString(string(tt.right)))
			assert.Equalf(t, tt.want, n.Data(), "Data()")
		})
	}
}

func TestNode_Depth(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		want int
	}{
		{
			name: "split leaves",
			node: newNode(FromString("a"), FromString("b")),
			want: 2,
		},
		{
			name: "split branches",
			node: newNode(newNode(FromString("a"), FromString("b")), newNode(FromString("c"), FromString("d"))),
			want: 3,
		},
		{
			name: "split branches and leaves",
			node: newNode(newNode(FromString("a"), FromString("b")), FromString("cd")),
			want: 3,
		},
		{
			name: "deeper",
			node: newNode(
				newNode(
					FromString("a"),
					newNode(
						FromString("a"),
						newNode(
							FromString("a"),
							newNode(
								FromString("a"),
								FromString("b"),
							),
						),
					),
				),
				FromString("cd"),
			),
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.Depth(), "Depth()")
		})
	}
}

func TestNode_Index(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		r    rune
		want int
	}{
		{
			name: "empty",
			node: newNode(FromString(""), FromString("")),
			r:    'a',
			want: -1,
		},
		{
			name: "not found",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'g',
			want: -1,
		},
		{
			name: "first left",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'a',
			want: 0,
		},
		{
			name: "last left",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'c',
			want: 2,
		},
		{
			name: "first right",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'd',
			want: 3,
		},
		{
			name: "last right",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'f',
			want: 5,
		},
		{
			name: "multiple",
			node: newNode(FromString("abbc"), FromString("abbc")),
			r:    'b',
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.Index(tt.r), "Index(%v)", tt.r)
		})
	}
}

func TestNode_LastIndex(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		r    rune
		want int
	}{
		{
			name: "empty",
			node: newNode(FromString(""), FromString("")),
			r:    'a',
			want: -1,
		},
		{
			name: "not found",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'g',
			want: -1,
		},
		{
			name: "first left",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'a',
			want: 0,
		},
		{
			name: "last left",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'c',
			want: 2,
		},
		{
			name: "first right",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'd',
			want: 3,
		},
		{
			name: "last right",
			node: newNode(FromString("abc"), FromString("def")),
			r:    'f',
			want: 5,
		},
		{
			name: "multiple",
			node: newNode(FromString("abbc"), FromString("abbc")),
			r:    'b',
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.LastIndex(tt.r), "Index(%v)", tt.r)
		})
	}
}

func TestNode_Length(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		want int
	}{
		{
			name: "empty",
			node: newNode(FromString(""), FromString("")),
			want: 0,
		},
		{
			name: "left=1 right=0",
			node: newNode(FromString("a"), FromString("")),
			want: 1,
		},
		{
			name: "left=0 right=1",
			node: newNode(FromString(""), FromString("a")),
			want: 1,
		},
		{
			name: "left=3 right=5",
			node: newNode(FromString("abc"), FromString("defgh")),
			want: 8,
		},
		{
			name: "left=1,3,5 right=7,11",
			node: newNode(FromString("a").Append(FromString("bcd")).Append(FromString("efgh")), FromString("ijklmno").Append(FromString("pqrstuvwxyz"))),
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.Length(), "Length()")
		})
	}
}

func TestNode_Line(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		line int
		want string
	}{
		{
			name: "one line",
			node: newNode(FromString("abc"), FromString("def")),
			line: 0,
			want: "abcdef",
		},
		{
			name: "-1",
			node: newNode(FromString("abc"), FromString("def")),
			line: -1,
			want: "",
		},
		{
			name: "after end",
			node: newNode(FromString("abc"), FromString("def")),
			line: 1,
			want: "",
		},
		{
			name: "first of many",
			node: newNode(FromString("abc\ndef\nghi"), FromString("\njkl\nmno")),
			line: 0,
			want: "abc",
		},
		{
			name: "middle of many",
			node: newNode(FromString("abc\ndef\nghi"), FromString("\njkl\nmno")),
			line: 1,
			want: "def",
		},
		{
			name: "last of many",
			node: newNode(FromString("abc\ndef\nghi"), FromString("\njkl\nmno")),
			line: 4,
			want: "mno",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.Line(tt.line).String(), "Line(%v)", tt.line)
		})
	}
}

func TestNode_NewLineCount(t *testing.T) {
	tests := []struct {
		name string
		node Rope
		want int
	}{
		{
			name: "one line",
			node: newNode(FromString("abc"), FromString("def")),
			want: 0,
		},
		{
			name: "two lines (left)",
			node: newNode(FromString("abc\ndef"), FromString("ghi")),
			want: 1,
		},
		{
			name: "two lines (right)",
			node: newNode(FromString("abc"), FromString("def\nghi")),
			want: 1,
		},
		{
			name: "four lines",
			node: newNode(FromString("abc\ndef\nghi"), FromString("jkl\nmno")),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, tt.node.NewLineCount(), "NewLineCount()")
		})
	}
}

func TestNode_Prepend(t *testing.T) {
	tests := []struct {
		name    string
		start   *Node
		appends []string
		want    string
	}{
		{
			name:    "prepend to empty",
			start:   newNode(FromString(""), FromString("")).(*Node),
			appends: []string{"g", "h", "i"},
			want:    "ihg",
		},
		{
			name:    "prepend to double-leaf",
			start:   newNode(FromString("abc"), FromString("def")).(*Node),
			appends: []string{"g", "h", "i"},
			want:    "ihgabcdef",
		},
		{
			name:    "prepend to one-leaf, one-branch",
			start:   newNode(FromString("abc"), FromString("def").Append(FromString("ghi"))).(*Node),
			appends: []string{"j", "k", "l"},
			want:    "lkjabcdefghi",
		},
		{
			name:    "prepend to one-branch, one-leaf",
			start:   newNode(FromString("abc").Append(FromString("def")), FromString("ghi")).(*Node),
			appends: []string{"j", "k", "l"},
			want:    "lkjabcdefghi",
		},
		{
			name:    "prepend to double-branch",
			start:   newNode(FromString("abc").Append(FromString("def")), FromString("ghi").Append(FromString("jkl"))).(*Node),
			appends: []string{"m", "n", "o"},
			want:    "onmabcdefghijkl",
		},
		{
			name:    "prepend to oversize",
			start:   newNode(FromString(strings.Repeat("a", maxLeafSize)), FromString(strings.Repeat("b", maxLeafSize))).(*Node),
			appends: []string{"c", "d", "e"},
			want:    "edc" + strings.Repeat("a", maxLeafSize) + strings.Repeat("b", maxLeafSize),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n Rope = tt.start
			for _, a := range tt.appends {
				n = n.Prepend(newLeaf([]rune(a)))
			}
			assert.Equalf(t, tt.want, n.String(), "Prepend(%v)", tt.appends)
		})
	}
}

func TestNode_Split(t *testing.T) {
	type fields struct {
		left       Rope
		right      Rope
		weight     int
		lineWeight int
	}
	type args struct {
		at int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Rope
		want1  Rope
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				left:       tt.fields.left,
				right:      tt.fields.right,
				weight:     tt.fields.weight,
				lineWeight: tt.fields.lineWeight,
			}
			got, got1 := n.Split(tt.args.at)
			assert.Equalf(t, tt.want, got, "Split(%v)", tt.args.at)
			assert.Equalf(t, tt.want1, got1, "Split(%v)", tt.args.at)
		})
	}
}

func TestNode_String(t *testing.T) {
	type fields struct {
		left       Rope
		right      Rope
		weight     int
		lineWeight int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				left:       tt.fields.left,
				right:      tt.fields.right,
				weight:     tt.fields.weight,
				lineWeight: tt.fields.lineWeight,
			}
			assert.Equalf(t, tt.want, n.String(), "String()")
		})
	}
}

func TestNode_Sub(t *testing.T) {
	type fields struct {
		left       Rope
		right      Rope
		weight     int
		lineWeight int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Rope
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Node{
				left:       tt.fields.left,
				right:      tt.fields.right,
				weight:     tt.fields.weight,
				lineWeight: tt.fields.lineWeight,
			}
			assert.Equalf(t, tt.want, n.Sub(tt.args.start, tt.args.end), "Sub(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}
