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
func LoadImage(name string, info EntityInfo) (image.Image, error) {
	if val, ok := ImgCache[name]; ok {
		return val, nil
	}

	_, fname := filepath.Split(info.Picture.Layers[0].Path)

	f, err := os.Open(FactorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+fname))
	if err != nil {
		return nil, err
	}

	fmt.Println(FactorioPath + filepath.FromSlash("data/base/graphics/entity/"+name+"/hr-"+fname))

	img, err := png.Decode(bufio.NewReader(f))
	if err != nil {
		return nil, err
	}

	ImgCache[name] = img

	return img, nil
}
