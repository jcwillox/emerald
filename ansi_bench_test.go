package emerald

import (
	"testing"
)

func BenchmarkColorCodeBasic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCode("cyan")
	}
}

func BenchmarkColorCodeComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorCode("cyan:bu+red:i")
	}
}

func BenchmarkColorIndexFg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ColorIndexFg(6)
	}
}
