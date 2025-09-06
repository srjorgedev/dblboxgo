package pkg

import (
	"os"
	"strconv"

	"github.com/srjorgedev/dblboxgo/internal/domain/unit"
)

func GetUnitImages(numID int, Transform bool, TagSwitch bool) []unit.Images {
	bchacutURL := os.Getenv("BCHACUT_URL")
	bchaicoURL := os.Getenv("BCHAICO_URL")

	images := []unit.Images{}

	// Definimos cuántas imágenes necesitamos
	count := 1
	if Transform || TagSwitch {
		count = 2
	}
	if Transform && TagSwitch {
		count = 3
	}

	// Generamos las imágenes
	for i := 1; i <= count; i++ {
		suffix := ""
		if i > 1 {
			suffix = strconv.Itoa(i)
		}

		img := unit.Images{
			BChaCut: bchacutURL + strconv.Itoa(numID) + suffix + ".webp",
			BChaIco: bchaicoURL + strconv.Itoa(numID) + suffix + ".webp",
		}
		images = append(images, img)
	}

	return images
}
