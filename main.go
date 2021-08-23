package main

import (
	"flag"
	"fmt"
)

func main() {
	username := flag.String("username", "", "学号")
	password := flag.String("password", "", "密码")
	flag.Parse()
	if *username == "" || *password == "" {
		flag.Usage()
		return
	}

	*password = Encrypt(*password)
	jnuID := Login(*username, *password)
	firstID := List(jnuID)
	mainTable, secondTable := Review(firstID, jnuID)
	_, msg := Checkin(jnuID, mainTable, secondTable)
	fmt.Println(msg)
	PostNotice()
}

// func PostNotice() {
// 	os.Setenv("sckey", "SCT66541T7f9O8TnGeu9m9nlZndrVFo7h")
// 	var SendKey = os.Getenv("sckey")
// 	// fmt.Println(SendKey)

// 	msg := fmt.Sprintf("%s %s", SendKey, ".send?title=checkin_result&desp=success")
// 	// fmt.Println(msg)
// 	resp, err := http.Get("https://sctapi.ftqq.com/" + msg)
// 	if err != nil {
// 		// handle error
// 	}
// 	defer resp.Body.Close()
// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		// handle error
// 	}
// 	fmt.Println(string(body))
// }
