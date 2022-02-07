package inmemory

import (
	"context"

	linkapp "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/models"
)

type inmemory struct {
	iterator int64 // race unsafe
	links    []models.Link
}

func (in *inmemory) Create(_ context.Context, link *models.Link) error {
	link.ID = in.iterator
	in.iterator++

	in.links = append(in.links, *link)

	return nil
}

func (in *inmemory) FindByToken(_ context.Context, token string) (*models.Link, error) {
	for _, lnk := range in.links {
		if lnk.Token == token {
			return &lnk, nil
		}
	}

	return nil, linkapp.ErrLinkNotFound
}

func (in *inmemory) FindAllByToken(ctx context.Context, token string) (models.StatList, error) {
	return nil, nil
}

func (in *inmemory) SaveStat(ctx context.Context, id int64, ip string) error {
	return nil
}

func New(initItems []models.Link) linkapp.Repository {
	return &inmemory{
		iterator: 1,
		links:    initItems,
	}
}
