package main

import (
	"net/http"
	"net/url"
)

func PostNotice(msg string, sckey string) {

	var title string
	if msg == "重复提交问卷" {
		title = "健康状态申报重复"
	} else if msg == "插入问卷数据成功" {
		title = "您已完成今天的健康状态申报"
	}

	data := url.Values{}
	data.Add("title", title)
	data.Add("desp", msg)

	resp, err := http.PostForm("https://sctapi.ftqq.com/"+sckey+".send", data)
	if err != nil {
		println(err)
	}
	// defer
	resp.Body.Close()
	// ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

}
