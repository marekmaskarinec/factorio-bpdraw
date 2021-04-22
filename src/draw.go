package bpdraw

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

// For primitive, not animated, without alt-mode support, without circuit connection support entities
// such as solar panel
func NormalDraw(name string) (image.Image, error) {
	return LoadImage(name)
}

// Initializes drawing canvas with the right size
// ents - entity array from blueprint string
// offx - by how much change the X axis so that the smallest X coordinate in blueprint is on 0
// offx -                        Y                           Y
// returns the image canvas pointer and all drawers that are needed for this blueprint render
func Init(ents []Entity, offx, offy float64) *image.RGBA {
	var mw, mh float64

	for i := 0; i < len(ents); i++ {
		ents[i].Position.X -= offx
		ents[i].Position.Y -= offy
		if ents[i].Position.X+9*64 > mw {
			mw = ents[i].Position.X + 9*64
		}
		if ents[i].Position.Y+9 > mh {
			mh = ents[i].Position.Y + 9*64
		}
	}

	return image.NewRGBA(image.Rect(0, 0, int(mw), int(mh)))
}

// Main drawing function. Calls the corresponding drawing functions for every entity in the blueprint
// ents - entity array from the blueprint string
// dst  - the image used as a canvas for drawing
func Draw(ents []Entity, dst *image.RGBA) {
	for i := 0; i < len(ents); i++ {
		fmt.Printf("Drawing %s\n", ents[i].Name)

		var img image.Image
		var err error
		switch ents[i].Name {
		case "solar-panel", "accumulator":
			img, err = NormalDraw(ents[i].Name)
			if err != nil {
				fmt.Printf("Can't load %s. Make sure you provided correct factorio path.\n%s\n", ents[i].Name, err.Error())
				continue
			}
		default:
			fmt.Printf("Can't find proper drawer for %s. Please file an issue on github.\n", ents[i].Name)
			continue
		}

		pos := image.Point{int(ents[i].Position.X * 64), int(ents[i].Position.Y * 64)}
		r := image.Rectangle{pos, pos.Add(img.Bounds().Max)}
		draw.Draw(dst, r, img, image.Point{0, 0}, draw.Over)
	}

	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, dst)
	f.Close()
}
