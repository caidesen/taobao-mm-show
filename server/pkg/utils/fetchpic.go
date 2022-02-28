package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PicRes struct {
	Url  string `json:"url"`
	Text string `json:"text"`
}
type UomgData struct {
	Code   int    `json:"code"`
	Imgurl string `json:"imgurl"`
	Msg    string `json:"msg"`
}

func FindOneByUomg() (picRes *PicRes, err error) {
	resp, err := http.Get("https://api.uomg.com/api/rand.img3?format=json")
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var data UomgData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &PicRes{Url: data.Imgurl, Text: ""}, nil
}

func FindOneByXiaoApi() (picRes *PicRes, err error) {
	resp, err := http.Get("http://jiuli.xiaoapi.cn/i/mjx.php")
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	arr := strings.Split(string(body), "±")
	if len(arr) < 3 {
		return nil, fmt.Errorf("参数不对, %s", string(body))
	}
	return &PicRes{Url: arr[1][4:], Text: arr[2][1 : len(arr[2])-1]}, nil
}
