package pictureinfo

import "context"

// PictureInfo is an interface for a picture information.
type PictureInfo interface {
	URL() string        // "http://www.example.com/image.jpg"
	Name() string       // "Really cool picture"
	Tags() []string     // tags for this picture, e.g. "cat", "dog"
	Resolution() string // original resolution of the picture in pixels, e.g. "1024x768", "640x480"
}

// PictureInfoRepo is an interface for a repository of PictureInfo.
type PictureInfoRepo interface {
	// StorePictureInfo stores a PictureInfo into repository.
	StorePictureInfo(ctx context.Context, picInfo PictureInfo) error

	// SearchPictureInfo searches a PictureInfo from repository.
	// Search based on particular information in PictureInfo.
	SearchPictureInfo(ctx context.Context, picInfoRequest PictureInfo) ([]PictureInfo, error)
}
