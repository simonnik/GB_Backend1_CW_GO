package delivery

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	echoDelivery "github.com/simonnik/GB_Backend1_CW_GO/internal/app/echo/delivery"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
	contextUtils "github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/context"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/pkg/token"
)

func (d delivery) Create(ectx echo.Context) error {
	ectx.Logger().Info("Create")
	newLink := &Link{}
	if err := ectx.Bind(newLink); err != nil {
		return err
	}
	cfg := contextUtils.GetConfig(ectx.Request().Context())
	newLink.Token = token.GenerateToken(cfg.HashMinLength, cfg.HashSalt)

	if err := d.links.Create(ectx.Request().Context(), (*models.Link)(newLink)); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	link := strings.TrimRight(cfg.Host, "/") + "/" + newLink.Token
	return ectx.JSON(http.StatusOK, echoDelivery.Map{"link": link})
}
