package makename

import "testing"

func TestGenerate(t *testing.T) {
	_ = Generate()
}

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Generate()
	}
}
