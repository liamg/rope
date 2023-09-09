package rope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Rope_Basics(t *testing.T) {
	r1 := FromString("Hello")
	r2 := r1.Append(FromString(" World"))
	r3 := r2.Append(FromString("!"))
	r4 := r3.Balance()
	assert.Equal(t, r1.String(), "Hello")
	assert.Equal(t, r2.String(), "Hello World")
	assert.Equal(t, r3.String(), "Hello World!")
	assert.Equal(t, r4.String(), "Hello World!")
}

func Test_Rope_Split(t *testing.T) {

	tests := []struct {
		name      string
		appends   []string
		split     int
		wantLeft  string
		wantRight string
	}{
		{
			name:      "split at 0",
			appends:   []string{"Hello", " World!"},
			split:     0,
			wantLeft:  "",
			wantRight: "Hello World!",
		},
		{
			name:      "split at 1",
			appends:   []string{"Hello", " World!"},
			split:     1,
			wantLeft:  "H",
			wantRight: "ello World!",
		},
		{
			name:      "split at 5",
			appends:   []string{"Hello", " World!"},
			split:     5,
			wantLeft:  "Hello",
			wantRight: " World!",
		},
		{
			name:      "split at 6",
			appends:   []string{"Hello", " World!"},
			split:     6,
			wantLeft:  "Hello ",
			wantRight: "World!",
		},
		{
			name:      "split at 12",
			appends:   []string{"Hello", " World!"},
			split:     12,
			wantLeft:  "Hello World!",
			wantRight: "",
		},
		{
			name:      "split at 13",
			appends:   []string{"Hello", " World!"},
			split:     13,
			wantLeft:  "Hello World!",
			wantRight: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := FromString(tt.appends[0])
			for _, s := range tt.appends[1:] {
				r = r.Append(FromString(s))
			}
			l, r := r.Split(tt.split)
			assert.Equal(t, tt.wantLeft, l.String())
			assert.Equal(t, tt.wantRight, r.String())
		})
	}
}
