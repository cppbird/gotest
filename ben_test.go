package main

import "testing"

func BenchmarkLocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Local()
	}
}

func BenchmarkLocal1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Local1()
	}
}

func BenchmarkLoop(b *testing.B) {
	a := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		Loop(a, 1)
	}
}

func BenchmarkLoop2(b *testing.B) {
	a := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		Loop(a, 2)
	}
}

func BenchmarkLoop3(b *testing.B) {
	a := make([]int, 10000)
	for i := 0; i < b.N; i++ {
		Loop(a, 3)
	}
}
