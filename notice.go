package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	// "net/http"
	"os"
)

func PostNotice() {
	os.Setenv("sckey", "SCT66541T7f9O8TnGeu9m9nlZndrVFo7h")
	var SendKey = os.Getenv("sckey")
	// fmt.Println(SendKey)

	msg := fmt.Sprintf("%s %s", SendKey, ".send?title=checkin_result&desp=success")
	// fmt.Println(msg)
	resp, err := http.Get("https://sctapi.ftqq.com/" + msg)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
