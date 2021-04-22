package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

func LoadImage(name string) (image.Image, error) {
	f, err := os.Open(factorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+name+".png"))
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	return img, nil
}
