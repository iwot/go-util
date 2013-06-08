package util

import (
	"github.com/stretchrcom/testify/assert"
	"testing"
)

func TestStringWithFirstValue(t *testing.T) {
	sb := NewStringBuffer("sam", "ple")
	assert.Equal(t, "sample", sb.String(), "they should be equal")

	sb.Append("aa")
	assert.Equal(t, "sampleaa", sb.String(), "they should be equal")

	sb.Append("BB")
	assert.Equal(t, "sampleaaBB", sb.String(), "they should be equal")
}

func TestStringNoneFirstValue(t *testing.T) {
	sb := NewStringBuffer()

	assert.Equal(t, "", sb.String(), "they should be equal")

	sb.Append("aa")
	assert.Equal(t, "aa", sb.String(), "they should be equal")

	sb.Append("BB")
	assert.Equal(t, "aaBB", sb.String(), "they should be equal")
}

func BenchmarkStringAppend(b *testing.B) {
	b.ResetTimer()
	str := ""
	for i := 0; i < b.N; i++ {
		str += "a"
	}
}

func BenchmarkStringBufferAppend(b *testing.B) {
	b.ResetTimer()
	sb := NewStringBuffer()
	for i := 0; i < b.N; i++ {
		sb.Append("a")
	}
}
