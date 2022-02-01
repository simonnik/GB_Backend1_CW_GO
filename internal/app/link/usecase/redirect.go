package usecase

import (
	"context"
	"fmt"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (u usecase) FindByToken(ctx context.Context, link models.Link) (*models.Link, error) {
	mlink, err := u.repo.FindByToken(ctx, link)
	if err != nil {
		return nil, fmt.Errorf("link not found in repo: %w", err)
	}

	return mlink, nil
}
