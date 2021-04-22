package main

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

var imgChache map[string]image.Image

// Reads entity texture
// If the entity is animated, the entire texture atlas is returned
func LoadImage(name string) (image.Image, error) {
	if val, ok := imgChache[name]; ok {
		return val, nil
	}

	f, err := os.Open(factorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+name+".png"))
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	imgChache[name] = img

	return img, nil
}
