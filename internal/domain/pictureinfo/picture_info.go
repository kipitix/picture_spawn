package pictureinfo

type pictureInfo struct {
	name       string
	url        string
	tags       []string
	resolution string
}

func NewPictureInfo(name string, url string, tags []string, resolution string) *pictureInfo {
	return &pictureInfo{
		name:       name,
		url:        url,
		tags:       tags,
		resolution: resolution,
	}
}

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
