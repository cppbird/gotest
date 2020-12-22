package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
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

var img = flag.String("url", "", "图片路径")

func main() {

	// http.HandleFunc("/", myHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))
	flag.Parse()
	img := readImage()
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	var newh, neww, stx, sty, endx, endy int
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
			neww = height * 9 / 16
			newh = height
			stx = width/2 - neww/2
			endx = width/2 + neww/2
			sty = 0
			endy = height
		} else {
			newh = width * 16 / 9
			neww = width
			stx = 0
			sty = height/2 - newh/2
			endx = width
			sty = height/2 + newh/2
		}
	}
	rgba := cut(img, stx, sty, endx, endy)
	if newh/neww >= 1 {
		hstep := newh / 3
		wstep := neww / 3
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				subImage := cut(rgba, i*wstep, j*hstep, (i+1)*wstep, (j+1)*hstep)
				export(subImage, fmt.Sprintf("%d.png", 9-j*3-i))
			}
		}
	} else {
		wstep := neww / 3
		for i := 0; i < 3; i++ {
			subImage := cut(rgba, i*wstep, 0, (i+1)*wstep, newh)
			export(subImage, fmt.Sprintf("%d.png", 3-i))
		}
	}

}

// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	ct := make([]byte, 512)
// 	r.Body.Read(ct)
// 	fmt.Println(http.DetectContentType(ct))
// 	io.Copy(w, r.Body)
// }

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
