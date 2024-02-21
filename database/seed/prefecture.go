package seed

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
)

func SeedPrefecture() {
	// JSONファイルの読み込み
	file, err := ioutil.ReadFile("database/seeddata/prefectures.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	// Prefecture構造体のスライスを定義
	var prefectures []models.Prefecture

	// JSONデータのデコード
	err = json.Unmarshal(file, &prefectures)

	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	models.CreatePrefecturesBatches(prefectures)
}