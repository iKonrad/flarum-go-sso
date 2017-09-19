package main

import (
	"log"
	sso "github.com/iKonrad/flarum-go-sso/lib"
)

func main() {

	client := sso.NewClient("http://localhost:1234", "12345678", 14)

	pass := "Test.123";

	token, userId, err := client.LogIn("test", pass)

	if err != nil {
		panic(err)
	}

	//client.UpdateBio(token, userId, "lol")
	client.UpdateEmail(token, userId, "jarssssson@icloud.com", pass)

	log.Println(token, err)



}


