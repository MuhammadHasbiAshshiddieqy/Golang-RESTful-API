package service

import (
	"fmt"
  "log"
	"net/http"
  "encoding/json"

	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/db"
	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/model"
)

func putContent(w http.ResponseWriter, r *http.Request){
	d := db.Connect()
	defer d.Close()

  fmt.Println("Endpoint Hit: putContent")
  b, err := unmarshal(r)
	if err != nil {
		fmt.Println(err)
	}

	o, err := model.Update(b.BookID, b.Content, d)
	if err != nil {
		log.Printf("Error read model: %s", err)
	}

	res := mapping(o)

  json.NewEncoder(w).Encode(res)
}
