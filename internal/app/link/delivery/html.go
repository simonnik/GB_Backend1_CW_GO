package delivery

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type Payload struct {
	jwt.StandardClaims

	Name string
}

func (d delivery) HTML(ectx echo.Context) error {
	ectx.Logger().Info("HTML")

	token, err := d.getJWTToken()
	if err != nil {
		return err
	}

	return ectx.Render(http.StatusOK, "form.html", map[string]interface{}{"jwtToken": token})
}

func (d delivery) getJWTToken() (*string, error) {
	payload := Payload{
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &payload)

	signedToken, err := token.SignedString(d.secretKey)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
