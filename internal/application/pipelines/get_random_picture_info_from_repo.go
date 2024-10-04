package pipelines

import (
	"context"
	"time"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
)

type GetRandomImageFromRepo interface {
	Do(ctx context.Context) (imginfo.Picture, error)
}

type getRandomImageFromRepo struct {
	repo       imginfo.ImageRepo
	getTimeout time.Duration
}

func NewGetRandomPictureInfoFromRepo(repo imginfo.ImageRepo, getTimeout time.Duration) *getRandomImageFromRepo {
	return &getRandomImageFromRepo{
		repo:       repo,
		getTimeout: getTimeout,
	}
}

var _ GetRandomImageFromRepo = (*getRandomImageFromRepo)(nil)

func (p *getRandomImageFromRepo) Do(ctx context.Context) (imginfo.Picture, error) {
	// const errTemplate = "failed to get random picture info from repo: %w"

	// getCtx, cancel := context.WithTimeout(ctx, p.getTimeout)
	// defer cancel()

	// pi, err := p.repo.GetRandom(getCtx)

	// if err != nil {
	// 	return nil, fmt.Errorf(errTemplate, err)
	// }

	// return pi, nil
	return nil, nil
}
