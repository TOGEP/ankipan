package main

import (
	"database/sql"
	"net/http"

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
	card := new(models.Card)
	if err = c.Bind(card); err != nil {
		panic(err.Error())
	}

	//fixme user_idは仮置き
	query := "INSERT INTO cards(user_id, problem_statement, answer_text, memo) values(0, ?, ?, ?)"
	_, err = db.Exec(query, card.Problem, card.Anser, card.Memo)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, card)
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
	query := "INSERT INTO users(name, email, token, uid) values(?, ?, ?, ?)"
	result, err := db.Exec(query, user.Name, user.Email, getUUID(), user.UID)
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
