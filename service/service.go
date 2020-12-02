package service

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
	"github.com/MuhammadHasbiAshshiddieqy/Golang-RESTful-API/model"
)

// Body - struct to save request
type Body struct {
	BookID    		int 	         `json:"book_id,omitempty"`
	Description   string         `json:"description,omitempty"`
	Content   		string         `json:"content,omitempty"`
}

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

func HandleRequests() *mux.Router {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/bookid", getBookByID).Methods("POST")
	myRouter.HandleFunc("/updatecontent", putContent).Methods("PUT")

  return myRouter
}

func unmarshal(r *http.Request) (Body, error) {
	var b Body
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return b, err
	}
	err = json.Unmarshal(rBody, &b)
	if err != nil {
		return b, err
	}
	fmt.Printf("%+v \n", b)
	return b, nil
}
