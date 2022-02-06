package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) Redirect(ectx echo.Context) error {
	ectx.Logger().Info("Redirect")
	f := LinkFilter{}
	if err := ectx.Bind(&f); err != nil {
		return err
	}

	if f.Token == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "link id can't be empty")
	}

	link, err := d.links.FindByToken(ectx.Request().Context(), *f.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = d.links.SaveStat(ectx.Request().Context(), link.ID, ectx.RealIP())
	if err != nil {
		ectx.Echo().Logger.Error(err)
	}

	ectx.Response().Header().Set("Cache-Control", "no-cache")
	return ectx.Redirect(http.StatusMovedPermanently, link.Link)
}
