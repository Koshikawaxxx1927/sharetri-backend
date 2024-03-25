package seed

import (
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/utils"
)

func SeedTrip() {
	// JSONファイルの読み込み
	projectRoot := utils.ProjectRoot
	tripPath := fmt.Sprintf("%s/database/seeddata/trips.json", projectRoot)
	file, err := ioutil.ReadFile(tripPath)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var trips []models.Trip

	// JSONデータのデコード
	err = json.Unmarshal(file, &trips)

	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	for _, trip := range trips {
		trip.CreateTrip()
	}
}