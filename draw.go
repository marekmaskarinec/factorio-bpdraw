package main

import (
	"image"
	"image/png"
	"image/draw"
	"os"
	"fmt"
)

type drawer struct {
	items []string
	drawfn func(name string) (image.Image, error)
}

func NormalDraw(name string) (image.Image, error) {
	return LoadImage(name)
}

func Init(ents []Entity, offx, offy float64) (*image.RGBA, []drawer) {
	var mw, mh float64

	for i:=0; i < len(ents); i++ {
		ents[i].Position.X -= offx
		ents[i].Position.Y -= offy
		if ents[i].Position.X + 9*64 > mw {
			mw = ents[i].Position.X + 9*64
		}
		if ents[i].Position.Y + 9 > mh {
			mh = ents[i].Position.Y + 9*64
		}
	}

	return image.NewRGBA(image.Rect(0, 0, int(mw), int(mh))), []drawer{drawer{[]string{"solar-panel"}, NormalDraw}}
}

func Draw(ents []Entity, dst *image.RGBA, drs []drawer) {
	
	for i:=0; i < len(ents); i++ {
		fmt.Printf("Drawing %s\n", ents[i].Name)

		var cdrawer drawer
		inited := false
		for j:=0; j < len(drs) && !inited; j++ {
			for k:=0; k < len(drs[j].items) && !inited; k++ {
				if drs[j].items[k] == ents[i].Name {
					cdrawer = drs[j]
					inited = true
					break
				}
			}
		}

		if !inited {
			fmt.Printf("Can't find proper drawer for %s. Please file an issue on github.\n", ents[i].Name)
			continue
		}

		img, err := cdrawer.drawfn(ents[i].Name)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		pos := image.Point{int(ents[i].Position.X*64), int(ents[i].Position.Y*64)}
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
