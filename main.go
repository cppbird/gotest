package main

import (
	"fmt"

	"github.com/spaolacci/murmur3"
)

func main() {
	a := make([]int, 10)
	Fuck(a)
	fmt.Println(a)
}

var b = murmur3.Sum32([]byte("http://i0.hdslb.com/bfs/bbq/video-image/userface/1558868601542006937.png"))

// Comp .
func Comp(s string) {
	a := murmur3.Sum32([]byte(s))
	if a == b {
	}
}

func Fuck(s []int) {
	s = append(s, 10)
}
