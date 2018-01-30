package main

import (
	"net/http"
	"testing"
)

func BenchmarkEcho(b *testing.B) {

	// タイマーリセット
	b.ResetTimer()

	// Nはコマンド引数から与えられたベンチマーク時間から自動で計算される
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/", nil)
	}
}

func BenchmarkEchoParam(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		http.NewRequest("GET", "/user/:1", nil)
	}
}