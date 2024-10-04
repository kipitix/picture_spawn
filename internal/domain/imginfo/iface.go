package imginfo

import (
	"context"
)

// Picture is an interface for a picture information.
type Picture interface {
	ID() string         // UUID of this picture info
	URL() string        // "http://www.example.com/image800x600.jpg"
	Resolution() string // original resolution of the picture in pixels, e.g. "1024x768", "640x480"

	Image() Image // Parent image for this picture info
	SetImage(img Image) error
}

// Group of PictureInfo with the same image.
// For example it can be same image with different resolution.
type Image interface {
	ID() string     // UUID of this picture info group
	Name() string   // "Really cool picture 800x600"
	Tags() []string // tags for this picture, e.g. "cat", "dog"

	AddPicture(pic Picture) error // Append new picture to image
	Pictures() []Picture          // Slice of of PictureInfo with the same image
}

// // PictureRepo is an interface for a repository of Pic.
// type PictureRepo interface {
// 	// Put stores a PictureInfo into repository.
// 	Put(ctx context.Context, pic Picture) error

// 	Get(ctx context.Context, reqPic Picture) (Picture, error)

// 	// GetRandom(ctx context.Context) (Pic, error)

// 	// SearchPictureInfo searches a PictureInfo from repository.
// 	// Search based on particular information in imginfo.
// 	// Get(ctx context.Context, imginfoRequest PictureInfo) ([]PictureInfo, error)
// }

// ImageRepo is an interface for a repository of Image.
type ImageRepo interface {
	// Put stores a Image into repository.
	Put(ctx context.Context, img Image) error

	Get(ctx context.Context, reqImg Image) (Image, error)

	// GetRandom(ctx context.Context) (Image, error)

	// SearchImage searches a Image from repository.
	// Search based on particular information in Image.
	// Get(ctx context.Context, imageInfoRequest Image) ([]Image, error)

}
