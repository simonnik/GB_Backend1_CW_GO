package delivery

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (d delivery) Stat(ectx echo.Context) error {
	ectx.Logger().Info("Stat")
	f := LinkFilter{}
	if err := ectx.Bind(&f); err != nil {
		return err
	}

	if f.Token == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "link id can't be empty")
	}

	links, err := d.links.FindAllByToken(ectx.Request().Context(), *f.Token)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return ectx.Render(http.StatusOK, "stat.html", links)
}
