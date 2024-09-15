package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/cavaliergopher/grab/v3"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"github.com/geziyor/geziyor/export"
	log "github.com/sirupsen/logrus"
)

const (
	rootURL = "https://wallpaperswide.com"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.Info("Hello World")

	parseRoot()
}

func parseRoot() {

	parseWallFunc := func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("li.wall").Each(func(i int, s *goquery.Selection) {
			if href, isExist := s.Find("a").Attr("href"); isExist {
				// log.Info(s.Find("h1").Text())
				imagePageURL := fmt.Sprintf("%s/%s", rootURL, href)
				// log.Info(imageURL)
				parseImage(imagePageURL, "1920x1080")
			}
		})
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{rootURL},
		ParseFunc: parseWallFunc,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()

}

func parseImage(imageULR string, resolution string) {

	parseImageFunc := func(g *geziyor.Geziyor, r *client.Response) {
		r.HTMLDoc.Find("div.wallpaper-resolutions").Each(func(i int, s *goquery.Selection) {
			sel := s.Find("a")
			for sel.Nodes != nil {
				if sel.Nodes[0].FirstChild.Data == resolution {
					if href, isExist := sel.Attr("href"); isExist {
						imageURL := fmt.Sprintf("%s/%s", rootURL, href)
						grab.Get("./images/", imageURL)
						return
					}
				}
				sel = sel.NextFiltered("a")
			}
		})
	}

	geziyor.NewGeziyor(&geziyor.Options{
		StartURLs: []string{imageULR},
		ParseFunc: parseImageFunc,
		Exporters: []export.Exporter{&export.JSON{}},
	}).Start()
}
