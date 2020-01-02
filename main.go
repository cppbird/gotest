package main

import (
	"net/http"
	"time"
)

var mt = make(map[int]string)

var ts []int64

var tl = 1 << 10 * 3

func main() {
	http.HandleFunc("/test", Test)
	http.ListenAndServe("127.0.0.1:8000", nil)
	http.TimeoutHandler
}

func Test(resp http.ResponseWriter, r *http.Request) {
	time.Sleep(5)
}
