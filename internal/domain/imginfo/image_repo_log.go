package imginfo

import (
	"context"

	"github.com/rs/zerolog/log"
)

type ImageRepoLog struct {
}

func NewImageRepoLog() *ImageRepoLog {
	return &ImageRepoLog{}
}

var _ ImageRepo = (*ImageRepoLog)(nil)

func (p *ImageRepoLog) Put(ctx context.Context, img Image) error {
	log.Debug().Interface("put image", NewImageJSON(img)).Send()
	return nil
}

func (p *ImageRepoLog) Get(ctx context.Context, reqImg Image) (Image, error) {
	log.Debug().Interface("get image", NewImageJSON(reqImg)).Send()
	return nil, nil
}
