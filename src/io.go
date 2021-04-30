package bpdraw

import (
	"bufio"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"fmt"
)

var ImgCache map[string]image.Image
var FactorioPath string

// Reads entity texture
// If the entity is animated, the entire texture atlas is returned
func LoadImage(name, fname string) (image.Image, error) {
	_, fname = filepath.Split(fname)
	if val, ok := ImgCache[fname]; ok {
		return val, nil
	}

	fmt.Println(fname)

	f, err := os.Open(FactorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+fname))
	if err != nil {
		return nil, err
	}

	img, err := png.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	ImgCache[fname] = img

	return img, nil
}
