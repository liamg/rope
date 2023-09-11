package rope

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fibonacci(t *testing.T) {
	assert.Equal(t, []int{
		0, 1, 1, 2, 3, 5, 8, 13, 21, 34,
	}, fibonacci[:10])
}
