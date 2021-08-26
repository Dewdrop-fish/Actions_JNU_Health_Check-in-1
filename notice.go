package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func PostNotice(msg string, sckey string) {

	data := url.Values{}
	data.Add("title", "check-in_results")
	data.Add("desp", msg)

	resp, err := http.PostForm("https://sctapi.ftqq.com/"+sckey+".send", data)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
