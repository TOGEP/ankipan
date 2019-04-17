package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func main() {
	/*  db, err := sql.Open("mysql", "root:@/ankipan_card")
	      if err != nil {
	      panic(err.Error())
	    }
	    defer db.Close()*/

	e := echo.New()
	e.File("/", "views/page1.html")
	e.POST("/create", CreateCard)

	e.Logger.Fatal(e.Start(":8080"))
}

func CreateCard(c echo.Context) error {
	db, err := sql.Open("mysql", "root:@/ankipan_card")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	problem := c.FormValue("problem")
	anser := c.FormValue("anser")
	note := c.FormValue("note")

	card := models.Card{Problem: problem, Anser: anser, Note: note}

	query := "INSERT INTO user values(null,?,?,?,?)"
	result, err := db.Exec(query, "card", "problem", "anser", "note")
	if err != nil {
		panic(err.Error())
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}

	fmt.Println(card)
	fmt.Println(problem)
	fmt.Println(anser)
	fmt.Println(note)
	return c.JSON(http.StatusOK, card)
}
