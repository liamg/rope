package rope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeaf_Append(t *testing.T) {
	type args struct {
		n Rope
	}
	tests := []struct {
		name    string
		start   []rune
		appends []string
		want    string
	}{
		{
			name:    "append to empty",
			start:   []rune{},
			appends: []string{"a", "b", "c"},
			want:    "abc",
		},
		{
			name:    "append to non-empty",
			start:   []rune("abc"),
			appends: []string{"d", "e", "f"},
			want:    "abcdef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := newLeaf(tt.start)
			for _, a := range tt.appends {
				l = l.Append(newLeaf([]rune(a)))
			}
			assert.Equalf(t, tt.want, l.String(), "Append %s", tt.appends)
		})
	}
}

func TestLeaf_At(t *testing.T) {
	tests := []struct {
		name string
		leaf func() Rope
		at   int
		want rune
	}{
		{
			name: "empty",
			leaf: func() Rope {
				return newLeaf(nil)
			},
			at:   0,
			want: 0,
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.leaf()
			assert.Equalf(t, tt.want, l.At(tt.at), "At(%v)", tt.at)
		})
	}
}

func TestLeaf_Balance(t *testing.T) {
	type fields struct {
		data []rune
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
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Balance(), "Balance()")
		})
	}
}

func TestLeaf_Depth(t *testing.T) {
	type fields struct {
		data []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Depth(), "Depth()")
		})
	}
}

func TestLeaf_Index(t *testing.T) {
	type fields struct {
		data []rune
	}
	type args struct {
		r rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Index(tt.args.r), "Index(%v)", tt.args.r)
		})
	}
}

func TestLeaf_LastIndex(t *testing.T) {
	type fields struct {
		data []rune
	}
	type args struct {
		r rune
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.LastIndex(tt.args.r), "LastIndex(%v)", tt.args.r)
		})
	}
}

func TestLeaf_Length(t *testing.T) {
	type fields struct {
		data []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
			assert.Equalf(t, tt.want, l.Length(), "Length()")
		})
	}
}

func TestLeaf_Line(t *testing.T) {
	type fields struct {
		data []rune
	}
	type args struct {
		line int
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
			assert.Equalf(t, tt.want, l.Line(tt.args.line), "Line(%v)", tt.args.line)
		})
	}
}

func TestLeaf_NewLineCount(t *testing.T) {
	type fields struct {
		data []rune
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Leaf{
				data: tt.fields.data,
			}
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
