package delivery

import "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"

type delivery struct {
	links     link.Usecase
	secretKey []byte
}

func New(links link.Usecase, secret string) link.Delivery {
	return delivery{
		links:     links,
		secretKey: []byte(secret),
	}
}
