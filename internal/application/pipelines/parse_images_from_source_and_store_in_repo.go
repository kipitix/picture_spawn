package pipelines

import (
	"context"
	"fmt"
	"time"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
	"github.com/kipitix/picture_spawn/internal/domain/sourceparser"
	"github.com/rs/zerolog/log"
)

type ParseImagesFromSourceAndPutInRepo interface {
	Do(ctx context.Context) error
}

type parseImagesFromSourceAndPutInRepo struct {
	sourceParser sourceparser.SourceParser
	repo         imginfo.ImageRepo
	putTimeout   time.Duration
}

func NewParseImagesFromSourceAndPutInRepo(
	sourceParser sourceparser.SourceParser,
	repo imginfo.ImageRepo,
	putTimeout time.Duration,
) *parseImagesFromSourceAndPutInRepo {
	return &parseImagesFromSourceAndPutInRepo{
		sourceParser: sourceParser,
		repo:         repo,
		putTimeout:   putTimeout,
	}
}

var _ ParseImagesFromSourceAndPutInRepo = (*parseImagesFromSourceAndPutInRepo)(nil)

func (p *parseImagesFromSourceAndPutInRepo) Do(ctx context.Context) error {
	const (
		errT = "failed to parse and store picture info from source: %w"
	)

	parserCtx, cancelParse := context.WithCancel(ctx)
	defer cancelParse() // cancel the context when the function returns

	// start parsing
	go p.sourceParser.Parse(parserCtx)

	for {
		select {
		case imginfo, ok := <-p.sourceParser.ImageChan():
			if !ok {
				log.Warn().Msg("Parsing finished")
				return nil // parsing has finished
			}

			putCtx, cancelStore := context.WithTimeout(ctx, p.putTimeout)
			defer cancelStore()

			// store the picture info
			if err := p.repo.Put(putCtx, imginfo); err != nil {
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
