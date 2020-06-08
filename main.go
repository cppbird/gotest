package main

import (
	"encoding/json"
	"fmt"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	// fs := Fs{
	// 	List: []int64{12, 4},
	// }
	// fss := FsStr{
	// 	List: []string{"fuck"},
	// }
	// r := Rules{
	// 	ArchiveType: &ArchiveType{
	// 		TID:    []Fs{fs},
	// 		SubTID: []Fs{fs},
	// 	},
	// 	UPRule: &UPRule{
	// 		UName: []string{"fuck"},
	// 		MID:   []Fs{fs},
	// 		Fans:  123,
	// 	},
	// 	ArchiveProperty: &ArchiveProperty{
	// 		Titles: []FsStr{fss},
	// 		Ratio: []*Ratio{
	// 			&Ratio{
	// 				SubTID: []int64{1, 23},
	// 				MinYX:  0.23,
	// 				MaxYX:  0.89,
	// 			},
	// 		},
	// 		Tags: []FsStr{fss},
	// 	},
	// }
	// str, _ := json.Marshal(&r)
	// a := strings.Split("a", ",")
	fmt.Println(json.Marshal(&AppMatchAllReq{
		ExpContext: &ExpContext{},
	}))
}

type AppMatchAllReq struct {
	ExpContext           ExpContext `protobuf:"bytes,2,opt,name=exp_context,json=expContext,proto3" json:"exp_context,omitempty"`
	TerminalVersion      string     `protobuf:"bytes,1,opt,name=terminal_version,json=terminalVersion,proto3" json:"terminal_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

type ExpContext struct {
	Device               *Device  `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	Mid                  uint64   `protobuf:"varint,2,opt,name=mid,proto3" json:"mid,omitempty"`
	TraceId              string   `protobuf:"bytes,3,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	QueryId              string   `protobuf:"bytes,4,opt,name=query_id,json=queryId,proto3" json:"query_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}
type Device struct {
	// 产品编号
	// 数据平台分配：粉=1、白=2、蓝=3、直播姬=4、HD=5、海外=6、OTT=7、漫画=8、TV野版=9、小视频=10、网易漫画=11、网易漫画lite=12、网易漫画HD=13、国际版=14
	AppId int32 `protobuf:"varint,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// 版本号
	Build int32 `protobuf:"varint,2,opt,name=build,proto3" json:"build,omitempty"`
	// 设备id
	Buvid string `protobuf:"bytes,3,opt,name=buvid,proto3" json:"buvid,omitempty"`
	// 包类型
	MobiApp string `protobuf:"bytes,4,opt,name=mobi_app,json=mobiApp,proto3" json:"mobi_app,omitempty"`
	// 平台：ios/android
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty"`
	// 运行设备
	Device string `protobuf:"bytes,6,opt,name=device,proto3" json:"device,omitempty"`
	// 渠道
	Channel string `protobuf:"bytes,7,opt,name=channel,proto3" json:"channel,omitempty"`
	// 手机品牌
	Brand string `protobuf:"bytes,8,opt,name=brand,proto3" json:"brand,omitempty"`
	// 手机型号
	Model string `protobuf:"bytes,9,opt,name=model,proto3" json:"model,omitempty"`
	// 系统版本
	Osver                string   `protobuf:"bytes,10,opt,name=osver,proto3" json:"osver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

const (
	mutexLocked = 1 << iota // mutex is locked
	mutexWoken
	mutexStarving
	mutexWaiterShift = iota
)

type T struct {
	R Rules `json:"rules"`
	sync.Mutex
}

// Rules 稿件导入规则
type Rules struct {
	ArchiveType     *ArchiveType     `json:"archive_type"`
	UPRule          *UPRule          `json:"up_rule"`
	ArchiveProperty *ArchiveProperty `json:"archive_property"`
}

//ArchiveType archive tid/subtid
type ArchiveType struct {
	TID    []Fs `json:"tid"`
	SubTID []Fs `json:"subtid"`
}

// Fs .
type Fs struct {
	List       []int64 `json:"list"`
	ImportFlag bool    `json:"import_flag"`
}

// FsStr .
type FsStr struct {
	List       []string `json:"list"`
	ImportFlag bool     `json:"import_flag"`
}

// UPRule .
type UPRule struct {
	UName []string `json:"uname"`
	MID   []Fs     `json:"mids"`
	Fans  uint32   `json:"fans"`
}

//ArchiveProperty archive property
type ArchiveProperty struct {
	Titles    []FsStr `json:"titles"`
	CopyRight int64   `json:"copy_right"`
	//NOTE :画幅尺寸
	Resolution  int64    `json:"resolution"`
	Ratio       []*Ratio `json:"ratio"`
	MaxDuration int64    `json:"max_duration"`
	MinDuration int64    `json:"min_duration"`
	// 视频标签
	Tags []FsStr `json:"tags"`
}

//Ratio page ratio
type Ratio struct {
	SubTID []int64 `json:"sub_tid"`
	MinYX  float32 `json:"min_ratio_yx"`
	MaxYX  float32 `json:"max_ratio_yx"`
}

//UnMarshal ..
func (r *Rules) UnMarshal(s string) (err error) {
	return jsoniter.UnmarshalFromString(s, r)
}
