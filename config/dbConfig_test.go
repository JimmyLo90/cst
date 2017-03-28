package config

import (
	"testing"
)

//TestgetDbConfig 测试
func TestGetDbConfig(t *testing.T) {
	getDbConfig()
}

//BenchmarkgetDbConfig benchmark
func BenchmarkGetDbConfig(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		getDbConfig()
	}
}

//TestgetDbConfig 测试
func TestGetDbDSN(t *testing.T) {
	GetDbDSN()
}

//BenchmarkgetDbConfig benchmark
func BenchmarkGetDbDSN(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		GetDbDSN()
	}
}
