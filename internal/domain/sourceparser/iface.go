package sourceparser

import (
	"context"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
)

// SourceParser is an interface for parsing information about picture from a source.
type SourceParser interface {
	Parse(ctx context.Context)
	ImageChan() <-chan imginfo.Image
	ErrorChan() <-chan error
}
