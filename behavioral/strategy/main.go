package main

import "github.com/fwidjaya20/go-designpatterns/behavioral/strategy/login"

func main() {
	loginService := login.New()

	loginService.Login(login.PWA)
	loginService.Login(login.WEB)
}
