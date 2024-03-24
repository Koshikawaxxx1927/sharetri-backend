package seed

import (
	"log"
	"fmt"
	// "time"
	"encoding/json"
	"io/ioutil"
	"github.com/Koshikawaxxx1927/sharetri-backend/app/models"
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

func SeedUser() {
	// JSONファイルの読み込み
	projectRoot := utils.ProjectRoot
	prefecturePath := fmt.Sprintf("%s/database/seeddata/users.json", projectRoot)
	file, err := ioutil.ReadFile(prefecturePath)
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var users []models.User

	// JSONデータのデコード
	err = json.Unmarshal(file, &users)

	if err != nil {
		log.Fatal("Error decoding JSON:", err)
	}
	for _, user := range users {
		// user.LastLoginTime = time.Now()
		user.CreateUser()
	}
}