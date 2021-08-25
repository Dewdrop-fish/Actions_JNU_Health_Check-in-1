package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("username")
	password := os.Getenv("password")
	password = Encrypt(password)
	jnuID := Login(username, password)
	firstID := List(jnuID)
	mainTable, secondTable := Review(firstID, jnuID)
	_, msg := Checkin(jnuID, mainTable, secondTable)
	fmt.Println(msg)
	PostNotice(msg)
}
