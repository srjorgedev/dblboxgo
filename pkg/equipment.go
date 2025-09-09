package pkg

import (
	"os"
	"strconv"

	"github.com/srjorgedev/dblboxgo/internal/domain/equipment"
)

func GetEquipmentImages(numID int, rarity int, awaken bool, awakenFrom int) equipment.EquipmentImages {
	equipIMG := os.Getenv("EQUIP_IMG_URL")
	equipICO := os.Getenv("EQUIP_ICO_URL")

	images := equipment.EquipmentImages{}

	images.RarityImage = equipICO + strconv.Itoa(rarity) + ".webp"
	images.IconImage = equipIMG + strconv.Itoa(numID) + ".webp"

	if awaken {
		images.RarityImage = equipICO + strconv.Itoa(rarity) + "A.webp"
		images.IconImage = equipIMG + strconv.Itoa(awakenFrom) + ".webp"
	}

	return images
}
