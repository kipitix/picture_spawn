package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
	"github.com/kipitix/picture_spawn/internal/domain/sourceparser"
)

const (
	storeMaxDuration = 5 * time.Second
)

const (
	errTemplate = "failed to parse and store picture info from source: %w"
)

func ParsePictureInfoFromSource(ctx context.Context, sourceParser sourceparser.SourceParser, repo pictureinfo.PictureInfoRepo) error {

	parserCtx, cancelParse := context.WithCancel(ctx)
	defer cancelParse() // cancel the context when the function returns

	go sourceParser.Parse(parserCtx)

	for {
		select {
		case picInfo, ok := <-sourceParser.PictureInfoChan():
			if !ok {
				return nil // parsing has finished
			}

			storeCtx, cancelStore := context.WithTimeout(ctx, storeMaxDuration)
			defer cancelStore()

			// store the picture info
			if err := repo.StorePictureInfo(storeCtx, picInfo); err != nil {
				return fmt.Errorf(errTemplate, err)
			}

		case err := <-sourceParser.ErrorChan():
			return fmt.Errorf(errTemplate, err)
		}
	}
}
