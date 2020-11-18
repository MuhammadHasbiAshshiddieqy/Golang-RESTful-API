package service

import (
	"fmt"
  "log"
	"net/http"
  "encoding/json"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/db"
	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/model"
)

type resBook struct {
	Title   			string   `json:"title"`
	Description   string   `json:"description"`
	Content   		string   `json:"content"`
}

func mapping(o model.Book) resBook {
	res := resBook{}
	res.Title = o.Title.String
	res.Description = o.Description.String
	res.Content = o.Content.String

	return res
}

func getBookByID(w http.ResponseWriter, r *http.Request){
	d := db.Connect()
	defer d.Close()

  fmt.Println("Endpoint Hit: getBookID")
  b, err := unmarshal(r)
	if err != nil {
		fmt.Println(err)
	}

	o, err := model.Read(b.BookID, d)
	if err != nil {
		log.Printf("Error read model: %s", err)
	}

	res := mapping(o)

  json.NewEncoder(w).Encode(res)
}
