package main

import (
	"log"

	sso "github.com/iKonrad/go-flarum"
	"fmt"
)

func main() {
	adminToken := "TjqfSScoHofhYzPDchaEBbnNWAmHUjodvucmDvzk"
	client := sso.NewClient("http://localhost:1234", adminToken, 14)

	login := "loleko"
	pass := "testingpass"
	email := "lolek@bolekkkk.com"
	client.SignUp(login, email, pass)

	token, userId, err := client.LogIn(login, pass)

	client.UpdateUserAttribute(userId, "password", "testingpass")

	if err != nil {
		panic(err)
	}

	client.UpdateUserAttribute(userId, "bio", "testinggggg")
	client.UpdateUserAttribute(userId, "email", "lolll@lolson.com")

	client.DeleteUser(userId)

	log.Println(token, err)
}
