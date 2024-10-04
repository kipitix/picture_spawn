package sourceparser

import (
	"context"

	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/google/uuid"
	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
	"github.com/kipitix/picture_spawn/internal/domain/sourceparser"
	"github.com/rs/zerolog/log"
)

type WallpapersWideSourceParser struct {
	imageChan chan imginfo.Image
	errorChan chan error
}

func NewWallpapersWideSourceParser() *WallpapersWideSourceParser {
	return &WallpapersWideSourceParser{
		imageChan: make(chan imginfo.Image),
		errorChan: make(chan error),
	}
}

var _ sourceparser.SourceParser = (*WallpapersWideSourceParser)(nil)

func (p *WallpapersWideSourceParser) Parse(ctx context.Context) {
	p.parseWallPages()

	close(p.imageChan)
	close(p.errorChan)
}

func (p *WallpapersWideSourceParser) ImageChan() <-chan imginfo.Image {
	return p.imageChan
}

func (p *WallpapersWideSourceParser) ErrorChan() <-chan error {
	return p.errorChan
}

const (
	// rootURL string = "https://wallpaperswide.com"
	rootURL string = "https://wallpaperswide.com/biking-desktop-wallpapers.html"
)

func (p *WallpapersWideSourceParser) parseWallPages() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs:   []string{rootURL},
		ParseFunc:   p.parseWallFunc,
		LogDisabled: true,
	}).Start()
}

func (p *WallpapersWideSourceParser) parseWallFunc(g *geziyor.Geziyor, r *client.Response) {
	// parse image page
	r.HTMLDoc.Find("li.wall").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Find("a").Attr("href"); ok {
			imagePageURL := r.JoinURL(href)
			name := s.Find("h1").Text()
			p.parseImagePage(imagePageURL, name)
		}
	})
	// parse next page
	r.HTMLDoc.Find("div.pagination").Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Next »" {
			if href, ok := s.Attr("href"); ok {
				log.Debug().Msg(href)
				g.Get(r.JoinURL(href), p.parseWallFunc)
			}
		}
	})
}

func (p *WallpapersWideSourceParser) parseImagePage(imagePageULR string, name string) {
	image := imginfo.NewImage(uuid.NewString(), name, nil)

	parseImageFunc := func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("div.wallpaper-resolutions").Each(func(i int, s *goquery.Selection) {
			sel := s.Find("a")
			for sel.Nodes != nil {
				resolution := sel.Nodes[0].FirstChild.Data
				log.Debug().Msg(resolution)
				if href, ok := sel.Attr("href"); ok {
					imageURL := r.JoinURL(href)
					log.Debug().Msg(imageURL)

					picture := imginfo.NewPicture(uuid.New().String(), imageURL, resolution)

					image.AddPicture(picture)

					// p.imageChan <- imginfo.NewPicture(uuid.New().String(), imageURL, name, nil, resolution)
					// return
				}

				// if sel.Nodes[0].FirstChild.Data == resolution {
				// }
				sel = sel.NextFiltered("a")
			}
		})
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs:   []string{imagePageULR},
		ParseFunc:   parseImageFunc,
		LogDisabled: true,
	}).Start()

	p.imageChan <- image
}
