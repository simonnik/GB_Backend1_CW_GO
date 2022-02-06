package postgres

import (
	"context"
	"fmt"
)

func (r repository) SaveStat(ctx context.Context, id int64, ip string) error {
	query := "INSERT INTO links_stat (link_id, ip) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, id, ip)
	if err != nil {
		return fmt.Errorf("failed to insert stat to db: %w", err)
	}

	return nil
}
