package types

import (
	"testing"
)

const n = 16

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		randomString(n)
	}
}
