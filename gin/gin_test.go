package main

import (
	"testing"
	"net/http"
)

func BenchmarkGin(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/", nil)
	}
}

func BenchmarkGinParam(b *testing.B){
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/user/:1", nil)
	}
}
