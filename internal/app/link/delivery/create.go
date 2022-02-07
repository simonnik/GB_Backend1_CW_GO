package delivery

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	echoDelivery "github.com/simonnik/GB_Backend1_CW_GO/internal/app/echo/delivery"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
	contextUtils "github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/context"
)

func (d delivery) Create(ectx echo.Context) error {
	ectx.Logger().Info("Create")
	newLink := &Link{}
	if err := ectx.Bind(newLink); err != nil {
		return err
	}
	cfg := contextUtils.GetConfig(ectx.Request().Context())
	newLink.Token = d.token.Generate()

	if err := d.links.Create(ectx.Request().Context(), (*models.Link)(newLink)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	host := strings.TrimRight(cfg.Host, "/")
	link := host + ectx.Echo().Reverse("redirect", newLink.Token)
	stat := host + ectx.Echo().Reverse("stat", newLink.Token)
	return ectx.JSON(http.StatusOK, echoDelivery.Map{"link": link, "stat": stat})
}
