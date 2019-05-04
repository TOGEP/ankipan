package main

import (
	"database/sql"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

var db *sql.DB
var err error

func main() {
	e := echo.New()

	e.POST("/create", CreateCard)
	e.POST("/user", CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func getDB() {
	db, err = sql.Open("mysql", "root:@/ankipan")
	if err != nil {
		panic(err.Error())
	}
	return
}

func CreateCard(c echo.Context) error {
	getDB()
	card := new(models.Card)
	if err = c.Bind(card); err != nil {
		panic(err.Error())
	}

	//fixme user_idは仮置き
	query := "INSERT INTO cards(user_id, problem_statement, answer_text, memo, question_time) values(0,?,?,?,NOW())"
	_, err = db.Exec(query, card.Id, card.Problem, card.Anser)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, card)
}

func CreateUser(c echo.Context) error {
	getDB()
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		panic(err.Error())
	}

	query := "INSERT INTO users(name, email, token, uid, created_at) values(?, ?, ?, ?, NOW())"
	_, err = db.Exec(query, user.Name, user.Email, user.Token, user.Uid)
	if err != nil {
		panic(err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
