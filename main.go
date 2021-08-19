package main

import (
	"encoding/json"
	"log"
)

type Book struct {
	Author string `json:"author_name"`
	Year   int    `json:"year_edition"`
}

// Campos não exportaveis
type User struct {
	fullname string
	age      int
}

func JSONToUser(j string) User {
	var val struct {
		Fullname string `json:"full_name"`
		Age      int    `json:"user_age"`
	}

	err := json.Unmarshal([]byte(j), &val)
	if err != nil {
		log.Fatal(err)
	}

	return User{
		fullname: val.Fullname,
		age:      val.Age,
	}
}

// Permite exportar os campos
func (s *User) ToJSON() string {
	dynamicJSON := struct {
		Fullname string `json:"full_name"`
		Age      int    `json:"user_age"`
	}{
		Fullname: s.fullname,
		Age:      s.age,
	}

	byte, err := json.Marshal(dynamicJSON)
	if err != nil {
		log.Fatal(err)
	}

	return string(byte)
}

func main() {
	// ----------------------------------------------------------------
	// Transformar um JSON em STRUCT
	// ----------------------------------------------------------------

	strJson := `{"full_name":"vinicius mazon","user_age":27}`

	userStruct := JSONToUser(strJson)
	log.Print(userStruct)

	// ----------------------------------------------------------------
	// Transformar uma STRUCT em JSON
	// ----------------------------------------------------------------

	book := Book{
		Author: "vinicius mazon",
		Year:   2021,
	}

	byte, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(byte))

	// ----------------------------------------------------------------
	// Campos não exportaveis
	// ----------------------------------------------------------------

	exportableUserJson := User{
		fullname: "vinicius mazon",
		age:      26,
	}

	log.Print(exportableUserJson.ToJSON())

	// ----------------------------------------------------------------
	// Tags de json
	// ----------------------------------------------------------------

}
