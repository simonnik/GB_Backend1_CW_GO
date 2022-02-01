package postgres

import (
	"context"
	"fmt"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (r repository) FindByToken(ctx context.Context, link models.Link) (*models.Link, error) {
	query := "SELECT * FROM links WHERE token = $1"
	err := r.db.GetContext(ctx, &link, query, link.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to select item from db: %w", err)
	}

	return &link, nil
}
