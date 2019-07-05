package login

import (
	"fmt"

	"github.com/fwidjaya20/go-designpatterns/behavioral/strategy/login/strategy"
)

// ILoginService .
type ILoginService interface {
	Login(source string)
}

type loginService struct{}

// New is a constructor function for creating LoginService
func New() ILoginService {
	return &loginService{}
}

func (ls loginService) Login(source string) {
	var s strategy.LoginStrategy

	if PWA == source {
		s = strategy.LoginStrategy{Strategy: strategy.PWALoginStrategy{}}
	} else if WEB == source {
		s = strategy.LoginStrategy{Strategy: strategy.WEBLoginStrategy{}}
	} else {
		fmt.Println("Invalid Login Option!")
		return
	}

	s.Strategy.Login()
}
