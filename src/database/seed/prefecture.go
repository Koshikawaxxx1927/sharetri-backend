package seed

import (
	"log"
	"fmt"
	"encoding/json"
	"io/ioutil"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/utils"
)

func SeedPrefecture() {
	// JSONファイルの読み込み
	projectRoot := utils.ProjectRoot
	prefecturePath := fmt.Sprintf("%s/database/seeddata/prefectures.json", projectRoot)
	file, err := ioutil.ReadFile(prefecturePath)
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