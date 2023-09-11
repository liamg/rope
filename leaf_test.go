package rope

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaf_Append(t *testing.T) {
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
	tests := []struct {
		name      string
		start     []rune
		prepends  []string
		want      string
		wantDepth int
	}{
		{
			name:      "prepend to empty",
			start:     []rune{},
			prepends:  []string{"a", "b", "c"},
			want:      "cba",
			wantDepth: 1,
		},
		{
			name:      "prepend to non-empty",
			start:     []rune("abc"),
			prepends:  []string{"d", "e", "f"},
			want:      "fedabc",
			wantDepth: 1,
		},
		{
			name:      "prepend to non-empty with max length",
			start:     []rune(strings.Repeat("a", maxLeafSize)),
			prepends:  []string{"b", "c", "d"},
			want:      "dcb" + strings.Repeat("a", maxLeafSize),
			wantDepth: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.start)
			for _, a := range tt.prepends {
				l = l.Prepend(newLeaf([]rune(a)))
			}
			assert.Equalf(t, tt.want, l.String(), "Prepend %s", tt.prepends)
			assert.Equalf(t, tt.wantDepth, l.Depth(), "Depth() after append")
		})
	}
}

func TestLeaf_Split(t *testing.T) {
	tests := []struct {
		name  string
		input string
		at    int
		want1 string
		want2 string
	}{
		{
			name:  "empty",
			input: "",
			at:    0,
			want1: "",
			want2: "",
		},
		{
			name:  "split at -1",
			input: "abc",
			at:    -1,
			want1: "",
			want2: "abc",
		},
		{
			name:  "split at 0",
			input: "abc",
			at:    0,
			want1: "",
			want2: "abc",
		},
		{
			name:  "split at 1",
			input: "abc",
			at:    1,
			want1: "a",
			want2: "bc",
		},
		{
			name:  "split at 2",
			input: "abc",
			at:    2,
			want1: "ab",
			want2: "c",
		},
		{
			name:  "split past end",
			input: "abc",
			at:    5,
			want1: "abc",
			want2: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf([]rune(tt.input))
			got, got1 := l.Split(tt.at)
			assert.Equal(t, tt.want1, got.String())
			assert.Equal(t, tt.want2, got1.String())
		})
	}
}

func TestLeaf_String(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{
			name:  "empty",
			value: "",
		},
		{
			name:  "1",
			value: "a",
		},
		{
			name:  "multi lines",
			value: "a\nb\nc",
		},
		{
			name:  "emoji",
			value: "this is a thumb-up em oji: 👍 - did it work?",
		},
		{
			name:  "null bytes",
			value: "this is a null byte: \x00 - did it work?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf([]rune(tt.value))
			assert.Equal(t, tt.value, l.String())
		})
	}
}

func TestLeaf_Sub(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		start, end int
		want       string
	}{
		{
			name:  "empty",
			input: "",
			start: 0,
			end:   0,
			want:  "",
		},
		{
			name:  "start past end",
			input: "abc",
			start: 5,
			end:   10,
			want:  "",
		},
		{
			name:  "end past end",
			input: "abc",
			start: 0,
			end:   5,
			want:  "abc",
		},
		{
			name:  "start after beginning",
			input: "abc",
			start: 1,
			end:   3,
			want:  "bc",
		},
		{
			name:  "start at beginning",
			input: "abc",
			start: 0,
			end:   1,
			want:  "a",
		},
		{
			name:  "start at beginning, end at end",
			input: "abc",
			start: 0,
			end:   3,
			want:  "abc",
		},
		{
			name:  "start == end",
			input: "abc",
			start: 1,
			end:   1,
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf([]rune(tt.input))
			assert.Equal(t, tt.want, l.Sub(tt.start, tt.end).String())
		})
	}
}
