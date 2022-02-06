package usecase

import (
	"context"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (u usecase) FindAllByToken(ctx context.Context, token string) (models.StatList, error) {
	return u.repo.FindAllByToken(ctx, token)
}
