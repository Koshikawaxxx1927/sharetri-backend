package controllers_test

import (
	"fmt"
    "net/http"
	"testing"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/Koshikawaxxx1927/sharetri-backend/src/app/models"
)

func TestGetPrefecture200(t *testing.T) {
	for i, prefectureJson := range prefecturesJson {
		w := httptest.NewRecorder()
		prefectureGetPath := fmt.Sprintf("/prefecture/%v", i+1)
		req, _ := http.NewRequest("GET", prefectureGetPath, nil)
		Router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		expectedJsonBody := fmt.Sprintf("{\"prefecture\":{\"ID\":%v,\"Trips\":null,\"name\":\"%s\",\"kana\":\"%s\"}}", prefectureJson.ID, prefectureJson.Name, prefectureJson.Kana)
		assert.Equal(t, w.Body.String(), expectedJsonBody)
	}
}

func TestGetPrefecture404(t *testing.T) {
	w := httptest.NewRecorder()
	prefectureGetPath := fmt.Sprintf("/prefecture/%v", 0)
	req, _ := http.NewRequest("GET", prefectureGetPath, nil)
	Router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

var prefecturesJson = [47]models.Prefecture{
	{ ID: 1, Name: "北海道", Kana: "ホッカイドウ" },
	{ ID: 2, Name: "青森県", Kana: "アオモリケン" },
	{ ID: 3, Name: "岩手県", Kana: "イワテケン" },
	{ ID: 4, Name: "宮城県", Kana: "ミヤギケン" },
	{ ID: 5, Name: "秋田県", Kana: "アキタケン" },
	{ ID: 6, Name: "山形県", Kana: "ヤマガタケン" },
	{ ID: 7, Name: "福島県", Kana: "フクシマケン" },
	{ ID: 8, Name: "茨城県", Kana: "イバラキケン" },
	{ ID: 9, Name: "栃木県", Kana: "トチギケン" },
	{ ID: 10, Name: "群馬県", Kana: "グンマケン" },
	{ ID: 11, Name: "埼玉県", Kana: "サイタマケン" },
	{ ID: 12, Name: "千葉県", Kana: "チバケン" },
	{ ID: 13, Name: "東京都", Kana: "トウキョウト" },
	{ ID: 14, Name: "神奈川県", Kana: "カナガワケン" },
	{ ID: 15, Name: "新潟県", Kana: "ニイガタケン" },
	{ ID: 16, Name: "富山県", Kana: "トヤマケン" },
	{ ID: 17, Name: "石川県", Kana: "イシカワケン" },
	{ ID: 18, Name: "福井県", Kana: "フクイケン" },
	{ ID: 19, Name: "山梨県", Kana: "ヤマナシケン" },
	{ ID: 20, Name: "長野県", Kana: "ナガノケン" },
	{ ID: 21, Name: "岐阜県", Kana: "ギフケン" },
	{ ID: 22, Name: "静岡県", Kana: "シズオカケン" },
	{ ID: 23, Name: "愛知県", Kana: "アイチケン" },
	{ ID: 24, Name: "三重県", Kana: "ミエケン" },
	{ ID: 25, Name: "滋賀県", Kana: "シガケン" },
	{ ID: 26, Name: "京都府", Kana: "キョウトフ" },
	{ ID: 27, Name: "大阪府", Kana: "オオサカフ" },
	{ ID: 28, Name: "兵庫県", Kana: "ヒョウゴケン" },
	{ ID: 29, Name: "奈良県", Kana: "ナラケン" },
	{ ID: 30, Name: "和歌山県", Kana: "ワカヤマケン" },
	{ ID: 31, Name: "鳥取県", Kana: "トットリケン" },
	{ ID: 32, Name: "島根県", Kana: "シマネケン" },
	{ ID: 33, Name: "岡山県", Kana: "オカヤマケン" },
	{ ID: 34, Name: "広島県", Kana: "ヒロシマケン" },
	{ ID: 35, Name: "山口県", Kana: "ヤマグチケン" },
	{ ID: 36, Name: "徳島県", Kana: "トクシマケン" },
	{ ID: 37, Name: "香川県", Kana: "カガワケン" },
	{ ID: 38, Name: "愛媛県", Kana: "エヒメケン" },
	{ ID: 39, Name: "高知県", Kana: "コウチケン" },
	{ ID: 40, Name: "福岡県", Kana: "フクオカケン" },
	{ ID: 41, Name: "佐賀県", Kana: "サガケン" },
	{ ID: 42, Name: "長崎県", Kana: "ナガサキケン" },
	{ ID: 43, Name: "熊本県", Kana: "クマモトケン" },
	{ ID: 44, Name: "大分県", Kana: "オオイタケン" },
	{ ID: 45, Name: "宮崎県", Kana: "ミヤザキケン" },
	{ ID: 46, Name: "鹿児島県", Kana: "カゴシマケン" },
	{ ID: 47, Name: "沖縄県", Kana: "オキナワケン" },
}