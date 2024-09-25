package pictureinfo

type PictureInfoJson struct {
	URL        string   `json:"url"`
	Name       string   `json:"name"`
	Tags       []string `json:"tags"`
	Resolution string   `json:"resolution"`
}

func NewPictureInfoJson(pi PictureInfo) *PictureInfoJson {
	return &PictureInfoJson{
		Name:       pi.Name(),
		URL:        pi.URL(),
		Tags:       pi.Tags(),
		Resolution: pi.Resolution(),
	}
}

func NewPictureInfoFromJson(pij *PictureInfoJson) *pictureInfo {
	return NewPictureInfo(pij.Name, pij.URL, pij.Tags, pij.Resolution)
}
