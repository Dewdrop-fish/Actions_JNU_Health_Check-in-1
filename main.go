package main

import (
	"os"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	sckey := os.Getenv("SCKEY")

	password = Encrypt(password)
	jnuID := Login(username, password)
	firstID := List(jnuID)
	mainTable, secondTable := Review(firstID, jnuID)
	_, msg := Checkin(jnuID, mainTable, secondTable)

	// fmt.Println(msg)
	if len(sckey) > 1 {
		PostNotice(msg, sckey)
	}

}
