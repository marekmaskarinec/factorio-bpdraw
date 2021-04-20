package main

import (
	"fmt"
	"os"
)

func main() {
	toParse := "0eNqV1etugyAUB/B3OZ9hES+z9VWWZUF71p4E0QAuaxrffaBJ26w3+CYYfv+DXDxBqyYcDWkHzQmoG7SF5uMElvZaqtDnjiNCA+SwBwZa9qElrcW+VaT3vJfdgTTyHGYGpHf4C42Y2UvCSFJXQ/L5kwFqR45wLWFpHL/01LdovPk8nME4WD940CHRgzx/qxgc1wefY7CjpYx2MhoN70kHYWdIhTpu8vLEvOxBHCrsnKHuZWBxDiRt0Tjf92RS2ZKyI5+zvi/vkGUEmSWJ1Vn8ltZxZ6S242Acb1G5J/Xe4Pkd/D0RL1LwOhHPUvBNIi5S8G0anmSLLA1P+ipCpOFJm0VcDqj1MO0Pji93yp19sqhFhFlEmyLaLGPNeLKKJeNnfjl7QdTcumF85P1fHebLkOszQLjHl5u+ufq3MFDSr3noC+oPGrsO3Yiy3uZ1UdeVyMp5/gPQTCjf"

	if len(os.Args) > 1 {
		toParse = os.Args[1]
	}

	bp, err := parseBPString(toParse)
	if err != nil {
		panic(err)
	}

	fmt.Println(bp)
}
