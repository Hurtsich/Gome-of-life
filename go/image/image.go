package image

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func GetImageFromFilePath(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return png.Decode(f)
}

func StatusByColor(pixel color.Color) bool {
	blanc := color.NRGBA{uint8(255), uint8(255), uint8(255), uint8(255)}
	result := pixel != blanc
	return result
}

func Upscale(img image.PalettedImage, times int) *image.Paletted {
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y
	resized := resize.Resize(uint(width*times), uint(height*times), img, resize.NearestNeighbor)

	var palette = []color.Color{
		color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
		color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
	}
	photo := image.NewPaletted(resized.Bounds(), palette)

	for y := photo.Bounds().Min.Y; y < photo.Bounds().Max.Y; y++ {
		for x := photo.Bounds().Min.X; x < photo.Bounds().Max.X; x++ {
			photo.Set(x, y, resized.At(x, y))
		}
	}

	return photo
}
