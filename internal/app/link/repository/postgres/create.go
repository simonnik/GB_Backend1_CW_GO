package postgres

import (
	"context"
	"fmt"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (r repository) Create(ctx context.Context, link *models.Link) error {
	query := "INSERT INTO links (link, token) VALUES " +
		"(:link, :token) RETURNING id"
	res, err := r.db.NamedQueryContext(ctx,
		query,
		link,
	)
	if err != nil {
		return fmt.Errorf("failed to insert link to db: %w", err)
	}

	var id int64

	next := res.Next()
	if err = res.Scan(&id); err != nil || !next {
		return fmt.Errorf("failed to get last inserted id from db: %w (next: %t)", err, next)
	}

	link.ID = id
	return nil
}
