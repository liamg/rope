package rope

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaf_Append(t *testing.T) {
	type args struct {
		n Rope
	}
	tests := []struct {
		name      string
		start     []rune
		appends   []string
		want      string
		wantDepth int
	}{
		{
			name:      "append to empty",
			start:     []rune{},
			appends:   []string{"a", "b", "c"},
			want:      "abc",
			wantDepth: 1,
		},
		{
			name:      "append to non-empty",
			start:     []rune("abc"),
			appends:   []string{"d", "e", "f"},
			want:      "abcdef",
			wantDepth: 1,
		},
		{
			name:      "append to non-empty with max length",
			start:     []rune(strings.Repeat("a", maxLeafSize)),
			appends:   []string{"b", "c", "d"},
			want:      strings.Repeat("a", maxLeafSize) + "bcd",
			wantDepth: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.start)
			for _, a := range tt.appends {
				l = l.Append(newLeaf([]rune(a)))
			}
			assert.Equalf(t, tt.want, l.String(), "Append %s", tt.appends)
			assert.Equalf(t, tt.wantDepth, l.Depth(), "Depth() after append")
		})
	}
}

func TestLeaf_At(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		at   int
		want rune
	}{
		{
			name: "empty",
			data: nil,
			at:   0,
			want: -1,
		},
		{
			name: "first",
			data: []rune("abc"),
			at:   0,
			want: 'a',
		},
		{
			name: "middle",
			data: []rune("abc"),
			at:   1,
			want: 'b',
		},
		{
			name: "last",
			data: []rune("abc"),
			at:   2,
			want: 'c',
		},
		{
			name: "out of range",
			data: []rune("abc"),
			at:   3,
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.At(tt.at), "At(%v)", tt.at)
		})
	}
}

func TestLeaf_Balance(t *testing.T) {
	l := newLeaf([]rune("hello"))
	balanced := l.Balance()
	assert.Equal(t, *(l.(*Leaf)), balanced)
}

func TestLeaf_Depth(t *testing.T) {
	tests := []struct {
		name    string
		appends int
		want    int
	}{
		{
			name:    "1 deep",
			appends: 0,
			want:    1,
		},
		{
			name:    "2 deep",
			appends: 1,
			want:    2,
		},
		{
			name:    "3 deep",
			appends: 1000,
			want:    1001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf([]rune("a"))
			for i := 0; i < tt.appends; i++ {
				l = l.Append(newLeaf([]rune(strings.Repeat("a", maxLeafSize))))
			}
			assert.Equalf(t, tt.want, l.Depth(), "Depth()")
		})
	}
}

func TestLeaf_Index(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		r    rune
		want int
	}{
		{
			name: "empty",
			data: nil,
			r:    'a',
			want: -1,
		},
		{
			name: "not found",
			data: []rune("abc"),
			r:    'd',
			want: -1,
		},
		{
			name: "first",
			data: []rune("abc"),
			r:    'a',
			want: 0,
		},
		{
			name: "middle",
			data: []rune("abc"),
			r:    'b',
			want: 1,
		},
		{
			name: "last",
			data: []rune("abc"),
			r:    'c',
			want: 2,
		},
		{
			name: "multiple",
			data: []rune("abbc"),
			r:    'b',
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.Index(tt.r), "Index(%v)", tt.r)
		})
	}
}

func TestLeaf_LastIndex(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		r    rune
		want int
	}{
		{
			name: "empty",
			data: nil,
			r:    'a',
			want: -1,
		},
		{
			name: "not found",
			data: []rune("abc"),
			r:    'd',
			want: -1,
		},
		{
			name: "first",
			data: []rune("abc"),
			r:    'a',
			want: 0,
		},
		{
			name: "middle",
			data: []rune("abc"),
			r:    'b',
			want: 1,
		},
		{
			name: "last",
			data: []rune("abc"),
			r:    'c',
			want: 2,
		},
		{
			name: "multiple",
			data: []rune("abbc"),
			r:    'b',
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.LastIndex(tt.r), "Index(%v)", tt.r)
		})
	}
}

func TestLeaf_Length(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		want int
	}{
		{
			name: "empty",
			data: nil,
			want: 0,
		},
		{
			name: "1",
			data: []rune("a"),
			want: 1,
		},
		{
			name: "2",
			data: []rune("ab"),
			want: 2,
		},
		{
			name: "3",
			data: []rune("abc"),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.Length(), "Length()")
		})
	}
}

func TestLeaf_Line(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		line int
		want string
	}{
		{
			name: "empty",
			data: nil,
			line: 0,
			want: "",
		},
		{
			name: "-1",
			data: []rune("a"),
			line: -1,
			want: "",
		},
		{
			name: "0 of 1",
			data: []rune("a"),
			line: 0,
			want: "a",
		},
		{
			name: "1 of 3",
			data: []rune("abc\ndef\nghi"),
			line: 1,
			want: "def",
		},
		{
			name: "2 of 3",
			data: []rune("abc\ndef\nghi"),
			line: 2,
			want: "ghi",
		},
		{
			name: "3 of 3",
			data: []rune("abc\ndef\nghi"),
			line: 3,
			want: "",
		},
		{
			name: "4 of 3",
			data: []rune("abc\ndef\nghi"),
			line: 4,
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.Line(tt.line).String(), "Line(%v)", tt.line)
		})
	}
}

func TestLeaf_NewLineCount(t *testing.T) {
	tests := []struct {
		name string
		data []rune
		want int
	}{
		{
			name: "empty",
			data: nil,
			want: 0,
		},
		{
			name: "0",
			data: []rune("hello world"),
			want: 0,
		},
		{
			name: "1",
			data: []rune("a\nb"),
			want: 1,
		},
		{
			name: "3",
			data: []rune("a\nb\nc"),
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.data)
			assert.Equalf(t, tt.want, l.NewLineCount(), "NewLineCount()")
		})
	}
}

func TestLeaf_Prepend(t *testing.T) {
	type fields struct {
		data []rune
	}
	type args struct {
		n Rope
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
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Prepend(tt.args.n), "Prepend(%v)", tt.args.n)
		})
	}
}

func TestLeaf_Split(t *testing.T) {
	type fields struct {
		data []rune
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
			l := Leaf{
				data: tt.fields.data,
			}
			got, got1 := l.Split(tt.args.at)
			assert.Equalf(t, tt.want, got, "Split(%v)", tt.args.at)
			assert.Equalf(t, tt.want1, got1, "Split(%v)", tt.args.at)
		})
	}
}

func TestLeaf_String(t *testing.T) {
	type fields struct {
		data []rune
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
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.String(), "String()")
		})
	}
}

func TestLeaf_Sub(t *testing.T) {
	type fields struct {
		data []rune
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
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Sub(tt.args.start, tt.args.end), "Sub(%v, %v)", tt.args.start, tt.args.end)
		})
	}
}

func TestLeaf_leaves(t *testing.T) {
	type fields struct {
		data []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   []Rope
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.leaves(), "leaves()")
		})
	}
}
