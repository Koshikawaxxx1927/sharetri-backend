package config

import (
	"os"
	"fmt"
	"log"
	"gorm.io/gorm/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	
	"github.com/Koshikawaxxx1927/sharetri-backend/utils"
)

var NotFound = gorm.ErrRecordNotFound

var db *gorm.DB

func InitDB(env string, isReset bool, models ...interface{}) {

	dsn := setDSN(env)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
        panic("Could not connect with database!")
    }

	// テーブルの消去
	if isReset {
        db.Migrator().DropTable(models...)
    }

	db.Logger = db.Logger.LogMode(logger.Info)
    db.AutoMigrate(models...)
}

func GetDB() *gorm.DB {
	return db
}

func setDSN(env string) (dsn string) {
	// MySQLサーバに接続するためのデータソースネームを設定する
	// 設定変更は .env/.env.development または .env/.env.production を変更してください
	// 環境変数の読み込み
	projectRoot := utils.ProjectRoot
	envFile := fmt.Sprintf("%s/env/.env.%s", projectRoot, env)
	err := godotenv.Load(envFile) 
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	dbname := os.Getenv("DB_NAME")

	dsn = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	return dsn
}