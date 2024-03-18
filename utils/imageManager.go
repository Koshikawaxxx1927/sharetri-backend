package utils

import (
	"encoding/base64"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"github.com/google/uuid"
	"io/ioutil"
)

type Image struct {
	EncodedData string `json:"encodeddata" binding:"required"`
}

func SaveDecodedImage(encodedImageData string, saveDir string) (string, error) {
	// base64デコード
	decodedData, err := base64.StdEncoding.DecodeString(encodedImageData)
	if err != nil {
		return "", err
	}

	// デコードされたデータを画像として復元
	img, _, err := image.Decode(strings.NewReader(string(decodedData)))
	if err != nil {
		return "", err
	}

	// 保存先ディレクトリが存在しない場合は作成
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		return "", err
	}

	// 保存先ファイルパスの作成
	imageName := uuid.New().String()
	savePath := filepath.Join(saveDir, imageName + ".png")

	// 画像をファイルに保存
	outputFile, err := os.Create(savePath)
	if err != nil {
		return "", err
	}
	defer outputFile.Close()

	// PNG形式で画像を保存
	err = png.Encode(outputFile, img)
	if err != nil {
		return "", err
	}
	return savePath, nil
}

func DeleteFile(filePath string) error {
	// ファイルを削除
	err := os.Remove(filePath)
	return err
}

func EncodeImage(filePath string) (string, error) {
    // ファイルを読み込む
    data, err := ioutil.ReadFile(filePath)
    if err != nil {
        return "", err
    }
    // Base64エンコード
    encodedString := base64.StdEncoding.EncodeToString(data)
    return encodedString, nil
}