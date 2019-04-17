package main

import (
	"database/sql"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.File("/", "views/page1.html")
	e.POST("/create", CreateCard)
	e.Logger.Fatal(e.Start(":8080"))
}

func CreateCard(c echo.Context) error {
	db, err := sql.Open("mysql", "root:@/ankipan")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	problem := c.FormValue("problem")
	anser := c.FormValue("anser")
	note := c.FormValue("note")

	card := models.Card{Problem: problem, Anser: anser, Note: note}

	//fixme user_idは仮置き
	query := "INSERT INTO cards(user_id, problem_statement, answer_text, memo, question_time) values(0,?,?,?,NOW())"
	_, err = db.Exec(query, problem, anser, note)
	if err != nil {
		panic(err.Error())
	}

	return c.JSON(http.StatusOK, card)
}
