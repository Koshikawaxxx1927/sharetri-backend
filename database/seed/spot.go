package seed

import (
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func SeedSpot() {
	// JSONファイルの読み込み
	projectRoot := utils.ProjectRoot
	spotPath := fmt.Sprintf("%s/database/seeddata/spots.json", projectRoot)
	file, err := ioutil.ReadFile(spotPath)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var spots []models.Spot

	// JSONデータのデコード
	err = json.Unmarshal(file, &spots)

	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	for _, spot := range spots {
		spot.CreateSpot()
	}
}