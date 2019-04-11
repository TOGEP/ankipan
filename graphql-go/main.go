package main

import (
  "github.com/labstack/echo"
  "net/http"
  "fmt"
  "./models"
)

func main() {
  e := echo.New()

  e.GET("/hello", func(c echo.Context) error {
    return c.String(http.StatusOK, "Hello, World!")
  })

  e.POST("/create", CreateCard)

  e.Logger.Fatal(e.Start(":8080"))
}

func CreateCard(c echo.Context) error {
  problem := c.FormValue("problem")
  anser := c.FormValue("anser")
  note := c.FormValue("note")

  card := []models.Card{}

  fmt.Println(card)
  fmt.Println(problem)
  fmt.Println(anser)
  fmt.Println(note)
  return c.String(http.StatusOK, "Create Card")
}
