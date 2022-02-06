package link

import (
	"context"
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

var (
	ErrLinkNotFound = errors.New("link not found")
)

type Delivery interface {
	Create(ectx echo.Context) error
	Redirect(ectx echo.Context) error
	Stat(ectx echo.Context) error
	HTML(ectx echo.Context) error
}

type Usecase interface {
	Create(ctx context.Context, link *models.Link) error
	FindByToken(ctx context.Context, token string) (*models.Link, error)
	FindAllByToken(ctx context.Context, token string) (models.StatList, error)
	SaveStat(ctx context.Context, id int64, ip string) error
}

type Repository interface {
	Create(ctx context.Context, link *models.Link) error
	FindByToken(ctx context.Context, token string) (*models.Link, error)
	FindAllByToken(ctx context.Context, token string) (models.StatList, error)
	SaveStat(ctx context.Context, id int64, ip string) error
}
