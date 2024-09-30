package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/kipitix/picture_spawn/internal/domain/pictureinfo"
	"github.com/kipitix/picture_spawn/internal/domain/sourceparser"
	"github.com/rs/zerolog/log"
)

type ParsePictureInfoFromSourceAndPutInRepo interface {
	Do(ctx context.Context) error
}

type parsePictureInfoFromSourceAndPutInRepo struct {
	sourceParser sourceparser.SourceParser
	repo         pictureinfo.PictureInfoRepo
	putTimeout   time.Duration
}

func NewParsePictureInfoFromSourceAndPutInRepo(
	sourceParser sourceparser.SourceParser,
	repo pictureinfo.PictureInfoRepo,
	putTimeout time.Duration,
) *parsePictureInfoFromSourceAndPutInRepo {
	return &parsePictureInfoFromSourceAndPutInRepo{
		sourceParser: sourceParser,
		repo:         repo,
		putTimeout:   putTimeout,
	}
}

var _ ParsePictureInfoFromSourceAndPutInRepo = (*parsePictureInfoFromSourceAndPutInRepo)(nil)

func (p *parsePictureInfoFromSourceAndPutInRepo) Do(ctx context.Context) error {
	const (
		errT = "failed to parse and store picture info from source: %w"
	)

	parserCtx, cancelParse := context.WithCancel(ctx)
	defer cancelParse() // cancel the context when the function returns

	go p.sourceParser.Parse(parserCtx)

	for {
		select {
		case picInfo, ok := <-p.sourceParser.PictureInfoChan():
			if !ok {
				log.Warn().Msg("Parsing finished")
				return nil // parsing has finished
			}

			putCtx, cancelStore := context.WithTimeout(ctx, p.putTimeout)
			defer cancelStore()

			// store the picture info
			if err := p.repo.Put(putCtx, picInfo); err != nil {
				log.Warn().Msg("Parsing error")
				return fmt.Errorf(errT, err)
			}

		case err := <-p.sourceParser.ErrorChan():
			return fmt.Errorf(errT, err)

		case <-ctx.Done():
			log.Warn().Msg("Parsing canceled")
			return nil
		}
	}
}
