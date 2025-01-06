package main

import (
	"fmt"
	_ "image/png"
	"math"

	"github.com/palSagnik/hash-visualization.git/models"
	"github.com/palSagnik/hash-visualization.git/utils"
)

const (
	WIDTH = 400
	HEIGHT = 400
)


var pixels [WIDTH][HEIGHT] models.RGBA32

func renderPixels(f func(x float32, y float32) models.Color) {

	for y := 0; y < HEIGHT; y++ {
		normalizedY := float32(y)/float32(HEIGHT)*2.0 - 1.0
		for x := 0; x < WIDTH; x++ {
			normalizedX := float32(x)/float32(WIDTH)*2.0 - 1.0
			c := f(normalizedX, normalizedY)

			// Map color values from [-1, 1] to [0, 255]
			pixels[x][y].R = uint8((c.R + 1)/2 * 255)
			pixels[x][y].G = uint8((c.G + 1)/2 * 255)
			pixels[x][y].B = uint8((c.B + 1)/2 * 255)
			pixels[x][y].A = 255
		}
	}
}

func grayGradient(x float32, _ float32) models.Color {
	return models.Color {
		R: x,
		G: x,
		B: x,
	}
}

func coolGradient(x float32, y float32) models.Color {
	if x * y >= 0 {
		return models.Color {
			R: x,
			G: y,
			B: 1,
		}
	} else {
		r := math.Mod(float64(x), float64(y))

		return models.Color {
			R: float32(r),
			G: float32(r),
			B: float32(r),
		}
	}
}

func main() {
	renderPixels(coolGradient)
	filepath := "output.png"
	if err := utils.WriteRGBAPNG(WIDTH, HEIGHT, pixels, filepath); err != nil {
		fmt.Printf("Unable to write image: %s", err.Error())
	}
}