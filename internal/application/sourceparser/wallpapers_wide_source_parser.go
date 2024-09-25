package sourceparser

import (
	"context"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
	"github.com/kipitix/picture_spawn/internal/domain/sourceparser"
)

type WallpapersWideSourceParser struct {
	pictureInfoChan chan pictureinfo.PictureInfo
	errorChan       chan error
}

func NewWallpapersWideSourceParser() *WallpapersWideSourceParser {
	return &WallpapersWideSourceParser{
		pictureInfoChan: make(chan pictureinfo.PictureInfo),
		errorChan:       make(chan error),
	}
}

var _ sourceparser.SourceParser = (*WallpapersWideSourceParser)(nil)

func (p *WallpapersWideSourceParser) Parse(ctx context.Context) {

}

func (p *WallpapersWideSourceParser) PictureInfoChan() <-chan pictureinfo.PictureInfo {
	return p.pictureInfoChan
}

func (p *WallpapersWideSourceParser) ErrorChan() <-chan error {
	return p.errorChan
}
