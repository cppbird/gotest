package main

import (
	"github.com/alecthomas/jsonschema"
)

type S10 struct {
	Gs []GameScore `json:"gs"`
	Vs []Video     `json:"vs"`
}
type GameScore struct {
}
type Video struct {
	Svids     []int64 `json:"svids"`
	EndTime   []int64 `json:"end_time"`
	StartTime []int64 `json:""`
}

func main() {
	jsonschema.Reflect(&TestUser{})
}
