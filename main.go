package main

import (
	"database/sql"
	"math"
	"net/http"
	"time"

	"github.com/TOGEP/ankipan/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
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

func CreateCard(c echo.Context) error {
	db, err := getDB()
	defer db.Close()
	card := new(models.Card)
	if err = c.Bind(card); err != nil {
		panic(err.Error())
	}

	//fixme user_idは仮置き
	query := "INSERT INTO cards(user_id, problem_statement, answer_text, memo, question_time, solved_count) values(0, ?, ?, ?, NOW(), 0)"
	_, err = db.Exec(query, card.Problem, card.Anser, card.Memo)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, card)
}

func CreateUser(c echo.Context) error {
	db, err := getDB()
	defer db.Close()
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		panic(err.Error())
	}

	query := "INSERT INTO users(name, email, token, uid, created_at) values(?, ?, ?, ?, NOW())"
	_, err = db.Exec(query, user.Name, user.Email, getUUID(), user.Uid)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func GetCards(c echo.Context) error {
	db := gormDBConnect()
	defer db.Close()

	token := c.QueryParam("token")

	user := models.User{}
	db.First(&user, "token=?", token)

	if user.Id == 0 {
		// TODO return error information
		return c.JSON(http.StatusBadRequest, "bad token")
	}

	cards := []models.Card{}

	db.Find(&cards, "user_id=?", user.Id)

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
	query := "UPDATE cards SET question_time=? WHERE id=?"
	_, err = db.Exec(query, t.Add(time.Duration(24*math.Pow(2, float64(cnt)))*time.Hour), cardid)
	if err != nil {
		return c.String(http.InternalServerError, "更新に失敗")
	}

	return c.String(http.StatusOK, "success")
}
