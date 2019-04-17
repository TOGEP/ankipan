package main

import (
	"database/sql"
	"fmt"
	"net/http"
  "time"
	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)


func main() {
	e := echo.New()
	e.File("/", "views/page1.html")
	e.POST("/create", CreateCard)
  e.POST("/login", CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func CreateCard(c echo.Context) error {
  t := time.Now();
  db, err := sql.Open("mysql", "root:@/ankipan_card")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

	problem := c.FormValue("problem")
	anser := c.FormValue("anser")
	note := c.FormValue("note")

	card := models.Card{Problem: problem, Anser: anser, Note: note}

	//query := "INSERT INTO cards(problem_statement, answer_text, memo)values(?,?,?)"
	//result, err := db.Exec(query, "problem", "anser", "note")
  query := "INSERT INTO cards values(0,0,?,?,?,?)"
  result, err := db.Exec(query, "problem","anser","note",t)
	if err != nil {
		panic(err.Error())
	}
	if lastId, lerr := result.LastInsertId(); lerr != nil {
		fmt.Println("insert last id: %d", lastId)
	}

	fmt.Println(card)
	fmt.Println(problem)
	fmt.Println(anser)
	fmt.Println(note)
	return c.JSON(http.StatusOK, card)
}
