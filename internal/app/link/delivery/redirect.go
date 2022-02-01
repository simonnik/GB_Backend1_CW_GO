package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (d delivery) Redirect(ectx echo.Context) error {
	ectx.Logger().Info("Redirect")
	request := struct {
		Link
		LinkFilter
	}{}
	if err := ectx.Bind(&request); err != nil {
		return err
	}

	if request.LinkFilter.Token == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "link id can't be empty")
	}
	request.Link.Token = *request.LinkFilter.Token

	link, err := d.links.FindByToken(ectx.Request().Context(), models.Link(request.Link))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	ectx.Response().Header().Set("Cache-Control", "no-cache")
	return ectx.Redirect(http.StatusMovedPermanently, link.Link)
}
