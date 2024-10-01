package pictureinfo

// Base implementation of PictureInfo
type pictureInfo struct {
	id         string
	url        string
	name       string
	tags       []string
	resolution string
}

// NewPictureInfo creates a new PictureInfo object.
func NewPictureInfo(id string, url string, name string, tags []string, resolution string) *pictureInfo {
	return &pictureInfo{
		id:         id,
		url:        url,
		name:       name,
		tags:       tags,
		resolution: resolution,
	}
}

// Ensure what pictureInfo struct implements PictureInfo interface
var _ PictureInfo = (*pictureInfo)(nil)

func (p *pictureInfo) ID() string {
	return p.id
}

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
