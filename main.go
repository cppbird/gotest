package main

import (
	"bytes"
	"flag"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
)

var img = flag.String("imgpath", "", "input img path")
var outpath = flag.String("outpath", "", "output path")

type Cnt struct {
	Max int `json:"max"`
	Min int `json:"min"`
	M   interface{}
}

type AsrResp struct {
	TaskID      string `json:"task_id"`
	Result      string `json:"result"`
	Status      int    `json:"status"`
	Message     string `json:"message"`
	FlashResult struct {
		Duration  int  `json:"duration"`
		Completed bool `json:"completed"`
		Sentences []struct {
			Text      string `json:"text"`
			BeginTime int    `json:"begin_time"`
			EndTime   int    `json:"end_time"`
			ChannelID int    `json:"channel_id"`
		} `json:"sentences"`
		Latency int `json:"latency"`
	} `json:"flash_result"`
}

type MsgStruct struct {
	TaskID  string `json:"task_id"`
	BizID   string `json:"biz_id"`
	ModelID string `json:"model_id"`
	State   int    `json:"state"`
	Result  string `json:"result"`
}

type AudioAnalyserResult struct {
	Filename string `json:"fileName"`
	Segments []struct {
		Type     string  `json:"type"`
		StartSec float32 `json:"startSec"`
		EndSec   float32 `json:"endSec"`
	} `json:"segments"`
	VocalsPath string `json:"vocals_path"`
}

func main() {

	fmt.Println(4 & 4s)

	// f, err := os.Open("/Users/jdq/Desktop/1.m4a")
	// if err != nil {
	// 	panic(err)
	// }
	// hc := &http.Client{}
	// var asrReq *http.Request
	// var asrResp *http.Response
	// u := url.Values{}
	// u.Set("appkey", "JGu2gRVdb1cLAV11")
	// u.Set("token", "4bee07b129cb420d8f85c817d421da5f")
	// u.Set("format", "mp3")
	// if asrReq, err = http.NewRequest(http.MethodPost, "http://nls-gateway.cn-shanghai.aliyuncs.com/stream/v1/FlashRecognizer?"+u.Encode(), f); err != nil {
	// 	return
	// }
	// asrReq.Header.Set("Content-type", "application/octet-stream")
	// asrReq.Header.Set("Host", "nls-gateway.cn-shanghai.aliyuncs.com")
	// if asrResp, err = hc.Do(asrReq); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer asrResp.Body.Close()
	// bs, err := readAll(asrResp.Body, 16<<10)
	// r := &AsrResp{}
	// fmt.Println(jsoniter.UnmarshalFromString(string(bs), r))
	// fmt.Println(r)
	os.Exit(0)
	// c := Cnt{
	// 	Max: 12,
	// 	Min: 1,
	// 	M: struct {
	// 		MMax int
	// 		MMin int
	// 	}{
	// 		MMax: 10,
	// 		MMin: 2,
	// 	},
	// }
	// s, _ := jsoniter.MarshalToString(c)
	// cc := &Cnt{}
	// ccc := &Cnt{
	// 	Max: 123,
	// 	Min: 23,
	// }
	// jsoniter.UnmarshalFromString(s, cc)
	// ccc.M = cc.M
	// fmt.Println(jsoniter.MarshalToString(ccc))
	// os.Exit(0)
	// flag.Parse()

	// http.HandleFunc("/", myHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	img := readImage()
	stx, sty, endx, endy, neww, newh := adjuImg(img.Bounds().Dx(), img.Bounds().Dy())
	rgba := cut(img, stx, sty, endx, endy)

	if newh/neww >= 1 {
		hstep := newh / 3
		wstep := neww / 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				subImage := cut(rgba, i*wstep, j*hstep, (i+1)*wstep, (j+1)*hstep)
				export(subImage, fmt.Sprintf("%s/%d.png", *outpath, 9-j*3-i))
			}
		}
	} else {
		wstep := neww / 3
		for i := 0; i < 3; i++ {
			subImage := cut(rgba, i*wstep, 0, (i+1)*wstep, newh)
			export(subImage, fmt.Sprintf("%s/%d.png", *outpath, 3-i))
		}
	}

}

func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	buf := bytes.NewBuffer(make([]byte, 0, capacity))
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	ct := make([]byte, 512)
	r.Body.Read(ct)
	fmt.Println(http.DetectContentType(ct))
	io.Copy(w, r.Body)
}

func adjuImg(width, height int) (stx, sty, endx, endy, neww, newh int) {
	if width/height < 1 {
		if float64(width)/float64(height) > 0.5625 {
			newh = width * 16 / 9
			neww = width
			stx = 0
			sty = height/2 - newh/2
			endx = width
			sty = height/2 + newh/2
		} else {
			neww = height * 9 / 16
			newh = height
			stx = width/2 - neww/2
			endx = width/2 + neww/2
			sty = 0
			endy = height
		}
	} else {
		if float64(width)/float64(height) > 1.77777 {
			neww = height * 16 / 9
			newh = height
			stx = width/2 - neww/2
			endx = width/2 + neww/2
			sty = 0
			endy = height
		} else {
			newh = width * 9 / 16
			neww = width
			stx = 0
			sty = height/2 - newh/2
			endx = width
			endy = height/2 + newh/2
		}
	}
	return
}

func cut(img image.Image, stx, sty, endx, endy int) (res image.Image) {
	switch v := img.(type) {
	case *image.RGBA:
		res = v.SubImage(image.Rect(stx, sty, endx, endy)).(*image.RGBA)
	case *image.NRGBA:
		res = v.SubImage(image.Rect(stx, sty, endx, endy)).(*image.NRGBA)
	}
	return
}

func export(img image.Image, filename string) {
	create, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	err = png.Encode(create, img)
	if err != nil {
		panic(err)
	}
}
func readImage() image.Image {
	if *img == "" {
		panic("img url empty")
	}

	f, err := os.Open(*img)
	if err != nil {
		panic(err)
	}
	m, err := png.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("read img success...")
	return m
}
