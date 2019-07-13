package main

import (
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/TOGEP/ankipan/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func gormDBConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	DBNAME := "ankipan"

	CONNECT := USER + "@" + "/" + DBNAME

	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler
	e.Use(middleware.CORS())

	e.POST("/cards", CreateCard)
	e.POST("/user", CreateUser)
	e.GET("/cards", GetCards)
	e.PUT("/anser/:cardid", UpdateTime)
	e.Logger.Fatal(e.Start(":8080"))
}

func getDB() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "root:@/ankipan")
	if err != nil {
		panic(err.Error())
	}
	return db, err
}

func getUUID() string {
	u, err := uuid.NewRandom()
	if err != nil {
		panic(err.Error())
	}
	uu := u.String()
	return uu
}

// CreateCard  カードを作成
func CreateCard(c echo.Context) error {
	db, err := getDB()
	defer db.Close()
	request := new(models.CreateCardRequest)
	if err = c.Bind(request); err != nil {
		panic(err.Error())
	}

	user := models.User{}
	gormDBConnect().First(&user, "token=?", request.Token)

	//fixme user_idは仮置き
	_, err = db.Exec("INSERT INTO cards(user_id, problem_statement, answer_text, memo, question_time, solved_count) values(?, ?, ?, ?, NOW(), 0)", user.ID, request.Problem, request.Anser, request.Memo)
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusOK, "success")
}

// CreateUser 新しいユーザーを登録する
func CreateUser(c echo.Context) error {
	db, err := getDB()
	defer db.Close()

	// FIXME DB connectionを2つ作るのもあれなので統一する
	gormDB := gormDBConnect()
	defer db.Close()

	user := new(models.User)
	if err = c.Bind(user); err != nil {
		panic(err.Error())
	}

	responseUser := models.User{}
	gormDB.First(&responseUser, "uid =?", user.UID)

	// FIXME ID == 0のとき見つからなかったとしている もっといいやり方がありそう
	if responseUser.ID != 0 {
		return c.JSON(http.StatusOK, responseUser)
	}

	// TODO user.Uidがfirebaseに登録されているか確認する必要がある
	// https://github.com/TOGEP/ankipan/issues/18
	result, err := db.Exec("INSERT INTO users(name, email, token, uid) values(?, ?, ?, ?)", user.Name, user.Email, getUUID(), user.UID)
	if err != nil {
		panic(err.Error())
	}

	userID, err := result.LastInsertId()
	gormDB.First(&responseUser, "id =?", userID)

	return c.JSON(http.StatusOK, responseUser)
}

// GetCards userの持ってるcardsを返す
func GetCards(c echo.Context) error {
	db := gormDBConnect()
	defer db.Close()

	token := c.QueryParam("token")

	// NOTE nilのときに文字列のnullが返ってくる
	if token == "null" {
		return c.JSON(http.StatusBadRequest, "token is null")
	}

	user := models.User{}
	db.First(&user, "token=?", token)

	if user.ID == 0 {
		// TODO return error information
		return c.JSON(http.StatusBadRequest, "bad token")
	}

	cards := []models.Card{}
	db.Find(&cards, "user_id=?", user.ID)

	return c.JSON(http.StatusOK, cards)
}

func UpdateTime(c echo.Context) error {
	//TODO 問合せしてきたユーザがこのカードを持っているか確認する
	cardid := c.Param("cardid")
	db, err := getDB()
	defer db.Close()

	// 何回このカードを復習したか取得
	var cnt int
	if err = db.QueryRow("SELECT solved_count FROM cards WHERE id=?", cardid).Scan(&cnt); err != nil {
		panic(err.Error())
	}

	// 現在時刻+(24*2^cnt)時間後の値をquestion_timeに代入
	t := time.Now()
	_, err = db.Exec("UPDATE cards SET question_time=? WHERE id=?", t.Add(time.Duration(24*math.Pow(2, float64(cnt)))*time.Hour), cardid)
	if err != nil {
		return c.String(http.StatusInternalServerError, "更新に失敗")
	}

	_, err = db.Exec("UPDATE cards SET solved_count=? WHERE id=?", cnt+1, cardid)
	return c.String(http.StatusOK, fmt.Sprintf("%v", 24*math.Pow(2, float64(cnt)))+"時間後に設定したよ")
}
