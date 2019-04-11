package models

type Card struct {
  Id int `json:id`
  Problem string `json:problem`
  Anser string `json:anser`
  Note string `json:note`
}
/*
type InputObjectConfig struct {
    Name        string      `json:"name"`
    Fields      interface{} `json:"fields"`
    Description string      `json:"description"`
}*/

