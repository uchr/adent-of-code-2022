package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func PrintResult(step int, r *Rope, tailPositions map[Point]bool) {
	width := 26
	height := 26
	centerx := 13
	centery := 13

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	for p := range tailPositions {
		img.Set(centerx+p.x, centery+p.y, color.RGBA{121, 199, 197, 0xff})
	}

	for i, part := range r.parts {
		img.Set(centerx+part.tail.x, centery+part.tail.y, color.RGBA{uint8(255.0 * float32(i) / 10.0), uint8(255.0 * float32(i) / 10.0), uint8(255.0 * float32(i) / 10.0), 0xff})
	}

	img.Set(centerx+r.parts[0].head.x, centery+r.parts[0].head.y, color.RGBA{183, 79, 111, 0xff})
	img.Set(centerx+r.parts[len(r.parts)-1].tail.x, centery+r.parts[len(r.parts)-1].tail.y, color.RGBA{115, 171, 132, 0xff})

	newImage := resize.Resize(uint(width*20), uint(height*20), img, resize.NearestNeighbor)

	os.MkdirAll("images", os.ModePerm)
	f, err := os.Create(fmt.Sprintf("images/image-%d.png", step))
	if err != nil {
		log.Panic(err)
	}
	png.Encode(f, newImage)
}
