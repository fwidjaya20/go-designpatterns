package strategy

import "fmt"

// ILoginStrategy .
type ILoginStrategy interface {
	Login()
}

// LoginStrategy .
type LoginStrategy struct {
	Strategy ILoginStrategy
}

// PWALoginStrategy .
type PWALoginStrategy struct{}

// WEBLoginStrategy .
type WEBLoginStrategy struct{}

// Login is a function for login with web logic
func (s WEBLoginStrategy) Login() {
	fmt.Println("Login with WEB")
}

// Login is a function for login with pwa logic
func (s PWALoginStrategy) Login() {
	fmt.Println("Login with PWA")
}
