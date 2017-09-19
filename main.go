package main

import (
	"log"
	sso "github.com/iKonrad/flarum-go-sso/lib"
)

func main() {

	client := sso.NewClient("http://localhost:1234", "12345678", 14)

	token, userId, err := client.LogIn("konrado", "test")

	if err != nil {
		panic(err)
	}

	client.UpdateBio(token, userId, "lol")

	log.Println(token, err)



}


