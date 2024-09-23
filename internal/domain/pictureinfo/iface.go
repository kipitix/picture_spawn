package pictureinfo

type PictureInfo interface {
	Name() string       // Really cool picture
	URL() string        // http://www.example.com/image.jpg
	Tags() []string     // tags for this picture, e.g. "cat", "dog"
	Resolution() string // original resolution of the picture in pixels, e.g. "1024x768", "640x480"
}
