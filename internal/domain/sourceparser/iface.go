package sourceparser

import (
	"context"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
)

// SourceParser is an interface for parsing information about picture from a source.
type SourceParser interface {
	Parse(ctx context.Context) error
	PictureInfoChan() <-chan pictureinfo.PictureInfo
}
