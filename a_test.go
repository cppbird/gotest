package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"testing"
)

var length = 1000000
var maps map[string]string

// var slices []string
// var arrays [1000]string

func init() {
	maps = make(map[string]string, length)
	slices = make([]string, length)
	for i := 0; i < length; i++ {
		maps[strconv.Itoa(i)] = fmt.Sprintf("%x", md5.Sum([]byte(strconv.Itoa(i))))
		// slices[i] = "abc"
		// arrays[i] = "abc"
	}

}

func BenchmarkIterateMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// maps[strconv.Itoa(i)] = "abcd"
		_ = maps[strconv.Itoa(i)]
	}
	b.RunParallel(func(b *testing.PB) {})
}
