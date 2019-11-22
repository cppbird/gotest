package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

func main() {
	bp := 80007234055
	for i := 0; i < 1000; i++ {
		bp++
		Curl(fmt.Sprintf("%012d", bp))
		time.Sleep(1 * time.Second)
	}
}

type Resp struct {
	Code int64 `json:"code,string"`
	Data struct {
		FunShoot struct {
			PlayURL string `json:"playUrl"`
		} `json:"funShoot"`
	} `json:"data"`
	Msg string `json:"msg"`
}

func Curl(str string) error {

	req, err := http.NewRequest(http.MethodPost, "https://m.migudm.cn/client/funShoot/queryFunShootDetail.html", strings.NewReader(fmt.Sprintf("hwOpusId=%s", str)))
	if err != nil {
		return err
	}
	req.Header.Set("authority", "m.migudm.cn")
	req.Header.Set("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Set("origin", "https://m.migudm.cn")
	req.Header.Set("x-requested-with", "XMLHttpRequest")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.70 Safari/537.36")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("referer", fmt.Sprintf("https://m.migudm.cn/client/funShoot/%s.html?fromShareFlag=1&traceId=ef8180e499fe4101abf32d1019bb71bb:04", str))
	req.Header.Set("cookie", "cookieid=2019103016420600011812604828c91d5fa1f794898b08e3028cf2601bb34271; SESSION=63b143ac-e296-477e-a898-a5396f98eba0")

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	iowriter := bytes.NewBuffer([]byte{})
	if _, err = iowriter.ReadFrom(resp.Body); err != nil {
		return err
	}
	ret := &Resp{}
	if err = jsoniter.UnmarshalFromString(iowriter.String(), ret); err != nil {
		return err
	}
	if ret.Data.FunShoot.PlayURL != "" {
		fmt.Println(ret.Data.FunShoot.PlayURL)
	}
	return nil
}
