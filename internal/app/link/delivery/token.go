package delivery

import (
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"
	pkgToken "github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/token"
)

type token struct {
	HashMinLength int
	HashSalt      string
}

func (t token) Generate() string {
	return pkgToken.GenerateToken(t.HashMinLength, t.HashSalt)
}

func NewToken(minL int, salt string) link.Token {
	return token{
		HashMinLength: minL,
		HashSalt:      salt,
	}
}
