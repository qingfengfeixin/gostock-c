package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"net/http"
	"strings"
)

func getinfo(stockcode string) (data []string) {
	url := "http://hq.sinajs.cn/list=" + stockcode

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get err=", err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read body err=", err)
	}
	bodystr := mahonia.NewDecoder("gbk").ConvertString(string(body))
	str := strings.Split(bodystr, "\"")
	data = strings.Split(str[1], ",")
	return data
}
