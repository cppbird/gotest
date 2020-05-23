package main

import (
	"testing"
)

const str = "http://i0.hdslb.com/bfs/bbq/video-image/userface/245188553/59b1266932e9a3298cc134f646d4ed7a348a2379.jpg"

func BenchmarkComp(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Comp(str)
	}
}
