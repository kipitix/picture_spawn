package imginfo_test

import (
	"testing"

	"github.com/kipitix/picture_spawn/internal/domain/imginfo"
)

func TestPictureCreation(t *testing.T) {
	pic := imginfo.NewPicture("123", "http://example.com/image.png", "800x600")
	if pic == nil {
		t.Error("Picture is nil")
	}
	if pic.ID() != "123" {
		t.Error("Picture ID is incorrect")
	}
	if pic.URL() != "http://example.com/image.png" {
		t.Error("Picture URL is incorrect")
	}
	if pic.Resolution() != "800x600" {
		t.Error("Picture Resolution is incorrect")
	}
}

func TestImageCreation(t *testing.T) {
	img := imginfo.NewImage("123", "Super vehicle", []string{"bikes", "cars", "trucks"})
	if img == nil {
		t.Error("Image is nil")
	}
	if img.ID() != "123" {
		t.Error("Image ID is incorrect")
	}
	if img.Name() != "Super vehicle" {
		t.Error("Image URL is incorrect")
	}
	if len(img.Tags()) != 3 {
		t.Error("Image Tags is incorrect")
	}
	if img.Tags()[0] != "bikes" {
		t.Error("Image Tags is incorrect")
	}
	if img.Tags()[1] != "cars" {
		t.Error("Image Tags is incorrect")
	}
	if img.Tags()[2] != "trucks" {
		t.Error("Image Tags is incorrect")
	}
}

func TestSetImageIntoPicture(t *testing.T) {
	img := imginfo.NewImage("123", "Super vehicle", []string{"bikes", "cars", "trucks"})
	pic := imginfo.NewPicture("456", "http://example.com/image.png", "800x600")
	if pic.Image() != nil {
		t.Error("Picture Image is not nil")
	}

	err := pic.SetImage(img)
	if pic.Image() == nil {
		t.Error("Picture Image is nil")
	}
	if err != nil {
		t.Error("Error is not nil")
	}

	err = pic.SetImage(img)
	if err == nil {
		t.Error("Error is nil")
	}

	err = pic.SetImage(nil)
	if err == nil {
		t.Error("Error is nil")
	}
}

func TestAddPicturesToImage(t *testing.T) {
	img := imginfo.NewImage("123", "Super vehicle", []string{"bikes", "cars", "trucks"})
	pic1 := imginfo.NewPicture("456", "http://example.com/image.png", "800x600")
	pic2 := imginfo.NewPicture("789", "http://example.com/image.png", "800x600")
	pic3 := imginfo.NewPicture("101", "http://example.com/image.png", "800x600")

	if pic1.Image() != nil {
		t.Error("Image Image is not nil")
	}
	if pic2.Image() != nil {
		t.Error("Image Image is not nil")
	}
	if pic3.Image() != nil {
		t.Error("Image Image is not nil")
	}

	err := img.AddPicture(pic1)
	if err != nil {
		t.Error("Error is not nil")
	}
	err = img.AddPicture(pic2)
	if err != nil {
		t.Error("Error is not nil")
	}
	err = img.AddPicture(pic3)
	if err != nil {
		t.Error("Error is not nil")
	}

	if pic1.Image() == nil {
		t.Error("Image Image is nil")
	}
	if pic2.Image() == nil {
		t.Error("Image Image is nil")
	}
	if pic3.Image() == nil {
		t.Error("Image Image is nil")
	}

	err = img.AddPicture(pic1)
	if err == nil {
		t.Error("Error is nil")
	}
	err = img.AddPicture(pic2)
	if err == nil {
		t.Error("Error is nil")
	}
	err = img.AddPicture(pic3)
	if err == nil {
		t.Error("Error is nil")
	}
}
