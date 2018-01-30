package foo

import "testing"

func BenchmarkMakeSlice(b *testing.B){
	b.ResetTimer()
	makeSlice(b.N)
}
