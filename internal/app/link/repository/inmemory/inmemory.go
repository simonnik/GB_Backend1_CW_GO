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

func (in *inmemory) FindByToken(_ context.Context, link models.Link) (*models.Link, error) {
	for _, lnk := range in.links {
		if lnk.Token == link.Token {
			return &lnk, nil
		}
	}

	return nil, linkapp.ErrItemNotFound
}

func New() linkapp.Usecase {
	return &inmemory{
		iterator: 1,
	}
}
