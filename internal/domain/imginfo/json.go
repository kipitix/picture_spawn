package imginfo

type PictureJSON struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	Resolution string `json:"resolution"`
	ImageID    string `json:"imageId"`
}

type ImageJSON struct {
	ID       string        `json:"id"`
	Name     string        `json:"name"`
	Tags     []string      `json:"tags"`
	Pictures []PictureJSON `json:"pictures"`
}

func NewImageJSON(img Image) *ImageJSON {
	imgJ := &ImageJSON{
		ID:       img.ID(),
		Name:     img.Name(),
		Tags:     img.Tags(),
		Pictures: make([]PictureJSON, len(img.Pictures())),
	}

	for i, p := range img.Pictures() {
		imgJ.Pictures[i] = PictureJSON{
			ID:         p.ID(),
			URL:        p.URL(),
			Resolution: p.Resolution()}
	}

	return imgJ
}
