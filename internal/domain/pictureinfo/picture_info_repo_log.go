package pictureinfo

import (
	"context"
	"strings"

	"github.com/rs/zerolog/log"
)

type PictureInfoRepoLog struct {
}

func NewPictureInfoRepoLog() *PictureInfoRepoLog {
	return &PictureInfoRepoLog{}
}

var _ PictureInfoRepo = (*PictureInfoRepoLog)(nil)

func (p *PictureInfoRepoLog) Put(ctx context.Context, picInfo PictureInfo) error {
	log.Debug().
		Str("name", picInfo.Name()).
		Str("URL", picInfo.URL()).
		Str("tags", strings.Join(picInfo.Tags(), ",")).
		Str("resolution", picInfo.Resolution()).
		Send()

	return nil
}

func (p *PictureInfoRepoLog) GetRandom(ctx context.Context) (PictureInfo, error) {
	return nil, nil
}

// func (p *PictureInfoRepoLog) SearchPictureInfo(ctx context.Context, picInfoRequest PictureInfo) ([]PictureInfo, error) {
// 	return nil, nil
// }
