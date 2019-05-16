package main

import (
	"database/sql"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = customHTTPErrorHandler

	e.POST("/cards", CreateCard)
	e.POST("/user", CreateUser)
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
