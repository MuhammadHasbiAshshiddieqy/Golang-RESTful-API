package service

import (
	"fmt"
	"net/http"
  "io/ioutil"
  "encoding/json"

  "github.com/gorilla/mux"
)

// Body - struct to save request
type Body struct {
	BookID    string         `json:"book_id"`
}

func HandleRequests() *mux.Router {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/bookid", getBookByID).Methods("POST")

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
