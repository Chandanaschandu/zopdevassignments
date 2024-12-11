package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	res, err := Stack(90, []int{10, 20, 30})
	assert.NoError(t, err)
	assert.Equal(t, []int{10, 20, 30}, res)

}
func BenchmarkPop(b *testing.B) {
	for range b.N {
		Stack(90, []int{10, 20, 30})
	}
}
func BenchmarkPush(b *testing.B) {
	for range b.N {
		Stack(90, []int{10, 20, 30})
	}
}
