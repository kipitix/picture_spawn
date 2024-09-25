package pictureinfo

// Base implementation of PictureInfo
type pictureInfo struct {
	url        string
	name       string
	tags       []string
	resolution string
}

// NewPictureInfo creates a new PictureInfo object.
func NewPictureInfo(url string, name string, tags []string, resolution string) *pictureInfo {
	return &pictureInfo{
		url:        url,
		name:       name,
		tags:       tags,
		resolution: resolution,
	}
}

// Ensure what pictureInfo struct implements PictureInfo interface
var _ PictureInfo = (*pictureInfo)(nil)

func (p *pictureInfo) URL() string {
	return p.url
}

func (p *pictureInfo) Name() string {
	return p.name
}

func (p *pictureInfo) Tags() []string {
	return p.tags
}

func (p *pictureInfo) Resolution() string {
	return p.resolution
}
