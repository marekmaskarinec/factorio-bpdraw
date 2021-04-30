package bpdraw

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

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
			mw = ents[i].Position.X + 18*64
		}
		if ents[i].Position.Y+9 > mh {
			mh = ents[i].Position.Y + 18*64
		}
	}

	return image.NewRGBA(image.Rect(0, 0, int(mw), int(mh)))
}

// Main drawing function. Calls the corresponding drawing functions for every entity in the blueprint
// ents - entity array from the blueprint string
// dst  - the image used as a canvas for drawing
func Draw(ents []Entity, dst *image.RGBA, info map[string]EntityInfo) {
	size := image.Rect(0, 0, 0, 0)

	for i := 0; i < len(ents); i++ {
		fmt.Printf("Drawing %s\n", ents[i].Name)

		var layers []Layer
		switch ents[i].Name {
		default:
			if _, ok := info[ents[i].Name]; !ok {
				fmt.Printf("Can't find proper drawer for %s. Please file an issue on github.\n", ents[i].Name)
				continue
			}
			if info[ents[i].Name].Picture.Layers == nil {
				fmt.Println(ents[i].Direction)
				switch ents[i].Direction {
				case 0:
					layers = info[ents[i].Name].Picture.North.Layers
					fmt.Println("her0")
				case 2:
					layers = info[ents[i].Name].Picture.West.Layers
					fmt.Println("her1")
				case 4:
					layers = info[ents[i].Name].Picture.East.Layers
					fmt.Println("her2")
				case 6:
					layers = info[ents[i].Name].Picture.South.Layers
					fmt.Println("her3")
				}
				break
			}
			layers = info[ents[i].Name].Picture.Layers	
		}

		for j:=len(layers)-1; j >= 0; j-- { 
			img, err := LoadImage(ents[i].Name, layers[j].Path)
			if err != nil {
				fmt.Printf("Can't load %s. Make sure you provided correct factorio path.\n%s\n", ents[i].Name, err.Error())
				continue
			}

			pos := image.Point{int(ents[i].Position.X * 64), int(ents[i].Position.Y * 64)}
			pos.X += int(layers[j].Shift[0] * 64) + 4.5 * 64
			pos.Y += int(layers[j].Shift[1] * 64) + 4.5 * 64

			layer := layers[j].HrVersion
			r := image.Rectangle{pos.Sub(image.Point{layer.Width/2, layer.Height/2}), pos.Add(img.Bounds().Max).Sub(image.Point{layer.Width/2, layer.Height/2})}
			r.Max.X = layer.Width + r.Min.X
			r.Max.Y = layer.Height + r.Min.Y
			//r.Add(off)
			draw.Draw(dst, r, img, image.Point{0, ents[i].Direction * layer.Height}, draw.Over)
      
			if r.Min.X < size.Min.X || i == 0 {
				size.Min.X = r.Min.X
			}
			if r.Min.Y < size.Min.Y || i == 0 {
				size.Min.Y = r.Min.Y
			}
			if r.Max.X > size.Max.X {
				size.Max.X = r.Max.X
			}
			if r.Max.Y > size.Max.Y {
				size.Max.Y = r.Max.Y
			}
		}
	}

	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(f, dst.SubImage(size))
	if err != nil {
		panic(err)
	}
	f.Close()
}
