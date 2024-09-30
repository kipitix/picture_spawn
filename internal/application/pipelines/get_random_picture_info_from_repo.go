package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
)

type GetRandomPictureInfoFromRepo interface {
	Do(ctx context.Context) (pictureinfo.PictureInfo, error)
}

type getRandomPictureInfoFromRepo struct {
	repo       pictureinfo.PictureInfoRepo
	getTimeout time.Duration
}

func NewGetRandomPictureInfoFromRepo(repo pictureinfo.PictureInfoRepo, getTimeout time.Duration) *getRandomPictureInfoFromRepo {
	return &getRandomPictureInfoFromRepo{
		repo:       repo,
		getTimeout: getTimeout,
	}
}

var _ GetRandomPictureInfoFromRepo = (*getRandomPictureInfoFromRepo)(nil)

func (p *getRandomPictureInfoFromRepo) Do(ctx context.Context) (pictureinfo.PictureInfo, error) {
	const errTemplate = "failed to get random picture info from repo: %w"

	getCtx, cancel := context.WithTimeout(ctx, p.getTimeout)
	defer cancel()

	pi, err := p.repo.GetRandom(getCtx)

	if err != nil {
		return nil, fmt.Errorf(errTemplate, err)
	}

	return pi, nil
}
