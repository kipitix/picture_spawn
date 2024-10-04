package imginfo

import "fmt"

// Base implementation of Picture
type picture struct {
	id         string
	url        string
	resolution string
	image      Image
}

// NewPicture creates a new Picture object
func NewPicture(id string, url string, resolution string) *picture {
	return &picture{
		id:         id,
		url:        url,
		resolution: resolution,
	}
}

// Ensure what pictureInfo struct implements PictureInfo interface
var _ Picture = (*picture)(nil)

func (p *picture) ID() string {
	return p.id
}

func (p *picture) URL() string {
	return p.url
}

func (p *picture) Resolution() string {
	return p.resolution
}

func (p *picture) SetImage(img Image) error {
	if p.image != nil {
		return fmt.Errorf("picture %v already has image %v", p.ID(), p.Image())
	}

	p.image = img

	return nil
}

func (p *picture) Image() Image {
	return p.image
}

// Base implementation of Image
type image struct {
	id          string
	name        string
	tags        []string
	pictures    []Picture
	picturesIds map[string]struct{}
}

// NewPicture creates a new Image object
func NewImage(id string, name string, tags []string) *image {
	return &image{
		id:          id,
		name:        name,
		tags:        tags,
		pictures:    make([]Picture, 0),
		picturesIds: make(map[string]struct{}),
	}
}

var _ Image = (*image)(nil)

func (i *image) ID() string {
	return i.id
}

func (i *image) Name() string {
	return i.name
}

func (i *image) Tags() []string {
	return i.tags
}

func (i *image) AddPicture(pic Picture) error {
	if _, ex := i.picturesIds[pic.ID()]; ex {
		return fmt.Errorf("picture %v already exists in image %v", pic.ID(), i.ID())
	}
	if pic.Image() != nil {
		return fmt.Errorf("picture %v already has image %v", pic.ID(), pic.Image())
	}

	i.picturesIds[pic.ID()] = struct{}{}
	i.pictures = append(i.pictures, pic)

	pic.SetImage(i)

	return nil
}

func (i *image) Pictures() []Picture {
	return i.pictures
}
