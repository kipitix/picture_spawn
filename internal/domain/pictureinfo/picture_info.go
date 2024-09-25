package pictureinfo

// Base implementation of PictureInfo
type pictureInfo struct {
	name       string
	url        string
	tags       []string
	resolution string
}

// NewPictureInfo creates a new PictureInfo object.
func NewPictureInfo(name string, url string, tags []string, resolution string) *pictureInfo {
	return &pictureInfo{
		name:       name,
		url:        url,
		tags:       tags,
		resolution: resolution,
	}
}

// Ensure what pictureInfo struct implements PictureInfo interface
var _ PictureInfo = (*pictureInfo)(nil)

func (p *pictureInfo) Name() string {
	return p.name
}

func (p *pictureInfo) URL() string {
	return p.url
}

func (p *pictureInfo) Tags() []string {
	return p.tags
}

func (p *pictureInfo) Resolution() string {
	return p.resolution
}
