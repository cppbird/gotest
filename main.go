package main

import (
	"github.com/spaolacci/murmur3"
)

func main() {

}

// Comp .
func Comp(s string) {
	a := murmur3.Sum32([]byte(s))
	if a == b {
	}
}

func Fuck(s []int) {
	s = append(s, 10)
}

// Rules 稿件导入规则
type Rules struct {
	ArchiveType     *ArchiveType     `json:"archive_type"`
	UPRule          *UPRule          `json:"up_rule"`
	ArchiveProperty *ArchiveProperty `json:"archive_property"`
}

//ArchiveType archive tid/subtid
type ArchiveType struct {
	TID []struct {
		List       []int32 `json:"tid"`
		ImportFlag bool    `json:"import_flag"`
	}
	SubTID []struct {
		List       []int32 `json:"sub_tid"`
		ImportFlag bool    `json:"import_flag"`
	}
}

// UPRule .
type UPRule struct {
	UName []string `json:"uname"`
	MID   []struct {
		List       []int64 `json:"mid"`
		ImportFlag bool    `json:"import_flag"`
	}
	Fans uint32 `json:"fans"`
}

//ArchiveProperty archive property
type ArchiveProperty struct {
	Titles []struct {
		List       []string `json:"title"`
		ImportFlag bool     `json:"import_flag"`
	}
	CopyRight int64 `json:"copy_right"`
	//NOTE :画幅尺寸
	Resolution  int64    `json:"resolution"`
	Ratio       []*Ratio `json:"ratio"`
	MaxDuration int64    `json:"max_duration"`
	MinDuration int64    `json:"min_duration"`
	// 视频标签
	Tags []struct {
		List       []string `json:"title"`
		ImportFlag bool     `json:"import_flag"`
	}
}

//Ratio page ratio
type Ratio struct {
	SubTID []int64
	MinYX  float32 `json:"min_ratio_yx"`
	MaxYX  float32 `json:"max_ratio_yx"`
}
