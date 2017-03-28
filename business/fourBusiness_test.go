package business

import "testing"

func TestSellAskCount(t *testing.T) {
	SellAskCount(1)
}

func BenchmarkSellAskCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SellAskCount(1)
	}
}

func TestYangxiuCount(t *testing.T) {
	YangxiuCount(1)
}

func BenchmarkYangxiuCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		YangxiuCount(1)
	}
}

func TestSellPromiseCount(t *testing.T) {
	SellPromiseCount(1)
}

func BenchmarkSellPromiseCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SellPromiseCount(1)
	}
}

func TestXubaoCount(t *testing.T) {
	XubaoCount(1)
}

func BenchmarkXubaoCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		XubaoCount(1)
	}
}

func TestSupportCount(t *testing.T) {
	SupportCount(1)
}

func BenchmarkSupportCount(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		SupportCount(1)
	}
}
