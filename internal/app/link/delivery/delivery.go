package delivery

import (
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"
)

type delivery struct {
	links     link.Usecase
	secretKey []byte
	token     link.Token
}

func New(links link.Usecase, secret string, token link.Token) link.Delivery {
	return delivery{
		links:     links,
		secretKey: []byte(secret),
		token:     token,
	}
}
