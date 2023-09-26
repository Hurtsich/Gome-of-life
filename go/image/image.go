package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
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
