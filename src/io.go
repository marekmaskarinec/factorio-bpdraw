package bpdraw

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"path/filepath"
)

var ImgCache map[string]image.Image
var FactorioPath string

// Reads entity texture
// If the entity is animated, the entire texture atlas is returned
func LoadImage(name string) (image.Image, error) {
	if val, ok := ImgCache[name]; ok {
		return val, nil
	}

	f, err := os.Open(FactorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+name+".png"))
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	ImgCache[name] = img

	return img, nil
}
