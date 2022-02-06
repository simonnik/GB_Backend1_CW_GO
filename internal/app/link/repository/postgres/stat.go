package postgres

import (
	"context"
	"fmt"
	"strings"

	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

func (r repository) FindAllByToken(ctx context.Context, token string) (models.StatList, error) {
	res := models.StatList{}

	query := strings.Builder{}
	query.WriteString("SELECT l.id, l.link, ls.ip, ls.created_at FROM links_stat ls")
	query.WriteString(" JOIN links l ON l.id = ls.link_id")
	query.WriteString(" WHERE l.token = $1")

	stmt, err := r.db.PreparexContext(ctx, query.String())
	if err != nil {
		return nil, fmt.Errorf("faield to prepare statement: %w", err)
	}

	err = stmt.SelectContext(ctx, &res, token)
	if err != nil {
		return nil, fmt.Errorf("failed to select link from db: %w", err)
	}

	return res, nil
}
