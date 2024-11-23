package utils

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/palSagnik/hash-visualization.git/models"
)

// WriteRGBAPNG writes RGBA pixel data to a PNG file
// width: image width
// height: image height
// pixels: slice of RGBA structs in row-major order
// filepath: output PNG file path
func WriteRGBAPNG(width, height int, pixels [400][400]models.RGBA32, filepath string) error {
    if len(pixels[0]) < width || len(pixels) < height {
        return fmt.Errorf("insufficient pixel data for %dx%d RGBA image", width, height)
    }

    // Create a new RGBA image
    img := image.NewRGBA(image.Rect(0, 0, width, height))

    // Copy pixel data
    for y := 0; y < height; y++ {
        for x := 0; x < width; x++ {
            img.Set(x, y, color.RGBA{
                R: pixels[x][y].R,
                G: pixels[x][y].G,
                B: pixels[x][y].B,
                A: pixels[x][y].A,
            })
        }
    }

    // Create output file
    f, err := os.Create(filepath)
    if err != nil {
        return err
    }
    defer f.Close()

    // Encode and write PNG
    return png.Encode(f, img)
}