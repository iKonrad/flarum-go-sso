package main

import (
	"log"
	sso "github.com/iKonrad/flarum-go-sso/lib"
)

func main() {

	adminToken := "TjqfSScoHofhYzPDchaEBbnNWAmHUjodvucmDvzk";
	client := sso.NewClient("http://localhost:1234", adminToken, 14)

	pass := "Test.123";

	token, userId, err := client.LogIn("test", pass)

	if err != nil {
		panic(err)
	}

	//client.UpdateBio(token, userId, "lol")
	client.UpdateEmail(adminToken, userId, "jaron@icloud.com", pass)

	log.Println(token, err)



}


