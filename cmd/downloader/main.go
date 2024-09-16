package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	log "github.com/sirupsen/logrus"
)

const (
	rootURL = "https://wallpaperswide.com/vintage-desktop-wallpapers.html"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.Info("Hello World")

	parseWallPages()
}

func parseWallFunc(g *geziyor.Geziyor, r *client.Response) {
	r.HTMLDoc.Find("li.wall").Each(func(i int, s *goquery.Selection) {
		if href, ok := s.Find("a").Attr("href"); ok {

			log.Info(s.Find("h1").Text())

			imagePageURL := r.JoinURL(href)
			log.Info(imagePageURL)

			// parseImagePage(imagePageURL, "1920x1080")
		}
	})

	r.HTMLDoc.Find("div.pagination").Find("a").Each(func(i int, s *goquery.Selection) {
		if s.Text() == "Next »" {
			if href, ok := s.Attr("href"); ok {
				log.Info(href)
				g.Get(r.JoinURL(href), parseWallFunc)
			}
		}
	})
}

func parseWallPages() {
	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{rootURL},
		ParseFunc: parseWallFunc,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}

func parseImagePage(imagePageULR string, resolution string) {

	parseImageFunc := func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("div.wallpaper-resolutions").Each(func(i int, s *goquery.Selection) {
			sel := s.Find("a")
			for sel.Nodes != nil {
				if sel.Nodes[0].FirstChild.Data == resolution {
					if href, ok := sel.Attr("href"); ok {
						imageURL := r.JoinURL(href)

						log.Info(imageURL)
						// grab.Get("./images/", imageURL)

						return
					}
				}
				sel = sel.NextFiltered("a")
			}
		})
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{imagePageULR},
		ParseFunc: parseImageFunc,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}
