package sougou

import (
	"io/ioutil"
	"net/http"
	"time"
)

var url = "https://weixin.sogou.com/weixin?type=2&s_from=input&ie=utf8&query="

func Act(q string) []byte {
	client := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}
	res, err := client.Get(url + q)
	if err != nil {
		panic("http.Get error")
	}

	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("ioutil.ReadAll error")
	}
	return result
}
