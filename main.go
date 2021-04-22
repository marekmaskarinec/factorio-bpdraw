package main

import (
	"fmt"
	"os"
	"image"
)

var factorioPath string

func main() {
	// train station, rails and 2 assemblers crafting burner-mining-drill with one input belt line and inserters
	//toParse := "0eNqV1etugyAUB/B3OZ9hES+z9VWWZUF71p4E0QAuaxrffaBJ26w3+CYYfv+DXDxBqyYcDWkHzQmoG7SF5uMElvZaqtDnjiNCA+SwBwZa9qElrcW+VaT3vJfdgTTyHGYGpHf4C42Y2UvCSFJXQ/L5kwFqR45wLWFpHL/01LdovPk8nME4WD940CHRgzx/qxgc1wefY7CjpYx2MhoN70kHYWdIhTpu8vLEvOxBHCrsnKHuZWBxDiRt0Tjf92RS2ZKyI5+zvi/vkGUEmSWJ1Vn8ltZxZ6S242Acb1G5J/Xe4Pkd/D0RL1LwOhHPUvBNIi5S8G0anmSLLA1P+ipCpOFJm0VcDqj1MO0Pji93yp19sqhFhFlEmyLaLGPNeLKKJeNnfjl7QdTcumF85P1fHebLkOszQLjHl5u+ufq3MFDSr3noC+oPGrsO3Yiy3uZ1UdeVyMp5/gPQTCjf"

	// single solar panel
	//toParse := "0eNptjsEOwiAQRP9lztjY2oryK8YYqhtDQhcC1Ng0/LtQLx68zUx23uyK0c7kg+EEtcLcHUeoy4ponqxtzdLiCQom0QQB1lN10Vkddl4zWWQBww96Q7X5KkCcTDL0xWxmufE8jRTKwV+AgHexdBzXvcLZdc0gsBRxbIZcodu6+nlW4EUhbpXu1Pby3MmDlEO773P+ABhkRWI="

	// 3 solar panels. 2 up one down in the middle
	toParse := "0eNqN0F0KwjAMAOC75Lkb7qdUexUR6TRIoUtL24lj9O62E0FQcG9JSL6QLDCYCZ3XFEEuoC+WAsjjAkHfSJlSi7NDkKAjjsCA1FiyYI3ylVOEBhIDTVd8gGzSiQFS1FHji1mT+UzTOKDPDT8BBs6GPGOp7MtOJWrOYM4Br3lK7MtpNzrdH6fb6PC30xYnH7l+Q348j8EdfVhH2n3Ti0MrOiF4s+tTegIAP3Ni"

	imgChache = map[string]image.Image{}

	if len(os.Args) > 1 {
		factorioPath = os.Args[1] + "/"
	}

	bp, err := parseBPString(toParse)
	if err != nil {
		panic(err)
	}

	_, err = LoadImage("assembling-machine-2")
	if err != nil {
		fmt.Println(err)
	}

	offx, offy := bp.FindZero()
	dst, drws := Init(bp.Entities, offx, offy)
	Draw(bp.Entities, dst, drws)
}
