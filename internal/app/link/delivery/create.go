package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoDelivery "github.com/simonnik/GB_Backend1_CW_GO/internal/app/echo/delivery"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (d delivery) Create(ectx echo.Context) error {
	ectx.Logger().Info("Create")
	newLink := &models.Link{}
	if err := ectx.Bind(newLink); err != nil {
		return err
	}
	if err := d.links.Create(ectx.Request().Context(), newLink); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ectx.JSON(http.StatusOK, echoDelivery.Map{"link": "http://localhost:8083/" + newLink.Token})
}
