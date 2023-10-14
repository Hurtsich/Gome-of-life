package image

import (
	"image"
	"image/color"
	"image/png"
	"os"
	"sync"
	"fmt"
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

func Upscale(img image.PalettedImage, times int) image.PalettedImage {
	pxColor := make(chan color.Color)
	wg := sync.WaitGroup{}
	width := img.Bounds().Max.X
	height := img.Bounds().Max.Y

	fmt.Println("BEGINNING UPSCALE :")
	fmt.Printf("Image size : w = %d, h = %d", width, height)
	fmt.Println()
	fmt.Printf("Image target size : w = %d, h = %d", width*times, height*times)

	var palette = []color.Color{
		color.RGBA{uint8(255), uint8(255), uint8(255), uint8(255)},
		color.RGBA{uint8(0), uint8(0), uint8(0), uint8(255)},
	}
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width*times, height*times}
	result := image.NewPaletted(image.Rectangle{topLeft, bottomRight}, palette)
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				pxColor <- img.At(j,i)
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < height*times; i+=times {
			for j := 0; j < width*times; j+=times {
				color := <- pxColor
				for k := i; k < i+times; k++ {
					for l := j; l < j+times; l++ {
						result.Set(l, k, color)
					}
				}
			}
		}
	}()

	wg.Done()

	return result
}
