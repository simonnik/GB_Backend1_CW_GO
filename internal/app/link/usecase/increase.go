package usecase

import (
	"context"
)

func (u usecase) SaveStat(ctx context.Context, id int64, ip string) error {
	return u.repo.SaveStat(ctx, id, ip)
}
