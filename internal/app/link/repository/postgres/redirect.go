package postgres

import (
	"context"
	"fmt"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (r repository) FindByToken(ctx context.Context, token string) (*models.Link, error) {
	link := models.Link{}
	query := "SELECT * FROM links WHERE token = $1"
	err := r.db.GetContext(ctx, &link, query, token)
	if err != nil {
		return nil, fmt.Errorf("failed to select link from db: %w", err)
	}

	return &link, nil
}
