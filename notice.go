package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func PostNotice(msg string) {
	// os.Setenv("sckey", "SCT66541T7f9O8TnGeu9m9nlZndrVFo7h")
	var SendKey = os.Getenv("sckey")
	// fmt.Println(SendKey)

	append := fmt.Sprintf("%s%s%s", SendKey, ".send?title=checkin_result&desp=", msg)
	fmt.Println(append)
	resp, err := http.Get("https://sctapi.ftqq.com/" + append)
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}
