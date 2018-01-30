package main

import (
	"testing"
	"net/http"
)

func BenchmarkNethttp(b *testing.B){

	// タイマーリセット
	b.ResetTimer()

	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/", nil)
	}
}