package business

import (
	"testing"
)

func TestSellAskCount(t *testing.T) {
	SellAskCount()
}

func BenchmarkSellAskCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SellAskCount()
	}
}
