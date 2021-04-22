package main

import (
	"os"
	"image"
	"github.com/marekmaskarinec/factorio-bpdraw/src"
	"fmt"
)


func main() {
	// train station, rails and 2 assemblers crafting burner-mining-drill with one input belt line and inserters
	//toParse := "0eNqV1etugyAUB/B3OZ9hES+z9VWWZUF71p4E0QAuaxrffaBJ26w3+CYYfv+DXDxBqyYcDWkHzQmoG7SF5uMElvZaqtDnjiNCA+SwBwZa9qElrcW+VaT3vJfdgTTyHGYGpHf4C42Y2UvCSFJXQ/L5kwFqR45wLWFpHL/01LdovPk8nME4WD940CHRgzx/qxgc1wefY7CjpYx2MhoN70kHYWdIhTpu8vLEvOxBHCrsnKHuZWBxDiRt0Tjf92RS2ZKyI5+zvi/vkGUEmSWJ1Vn8ltZxZ6S242Acb1G5J/Xe4Pkd/D0RL1LwOhHPUvBNIi5S8G0anmSLLA1P+ipCpOFJm0VcDqj1MO0Pji93yp19sqhFhFlEmyLaLGPNeLKKJeNnfjl7QdTcumF85P1fHebLkOszQLjHl5u+ufq3MFDSr3noC+oPGrsO3Yiy3uZ1UdeVyMp5/gPQTCjf"

	// single solar panel
	//toParse := "0eNptjsEOwiAQRP9lztjY2oryK8YYqhtDQhcC1Ng0/LtQLx68zUx23uyK0c7kg+EEtcLcHUeoy4ponqxtzdLiCQom0QQB1lN10Vkddl4zWWQBww96Q7X5KkCcTDL0xWxmufE8jRTKwV+AgHexdBzXvcLZdc0gsBRxbIZcodu6+nlW4EUhbpXu1Pby3MmDlEO773P+ABhkRWI="

	// 3 solar panels. 2 up one down in the middle
	//toParse := "0eNqN0F0KwjAMAOC75Lkb7qdUexUR6TRIoUtL24lj9O62E0FQcG9JSL6QLDCYCZ3XFEEuoC+WAsjjAkHfSJlSi7NDkKAjjsCA1FiyYI3ylVOEBhIDTVd8gGzSiQFS1FHji1mT+UzTOKDPDT8BBs6GPGOp7MtOJWrOYM4Br3lK7MtpNzrdH6fb6PC30xYnH7l+Q348j8EdfVhH2n3Ti0MrOiF4s+tTegIAP3Ni"

	// same as above, but with two accumulators added
	toParse := "0eNqNkdsKgzAMQP8lz51oVdz6K2OM6sIo1FR6GRPpv6+6C4MJ8y0JyckhmaDVAQeryIOYQHWGHIjjBE5dSeq55scBQYDy2AMDkv2cOaOl3Q2SUENkoOiCdxBFPDFA8sorfGKWZDxT6Fu0qWEVwGAwLs0Ymvclzq7IagZjCvKsjpH9cPg2Dv+DKT8Y2XWhD1p6Y1d0+ELha4hqm0n+MuHrIvU2kfLtkc68/EN8vY/BDa1bmvm+qJoDb8qmqYu8ivEBIVWfMA=="

	bpdraw.ImgCache = map[string]image.Image{}

	if len(os.Args) > 1 {
		bpdraw.FactorioPath = os.Args[1] + "/"
	}

	bp, err := bpdraw.ParseBPString(toParse)
	if err != nil {
		panic(err)
	}

	fmt.Println(bpdraw.ParseEntityInfo())

	offx, offy := bp.FindZero()
	dst := bpdraw.Init(bp.Entities, offx, offy)
	bpdraw.Draw(bp.Entities, dst)
}
