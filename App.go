package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"gopkg.in/macaron.v1"
	"gopkg.in/mgo.v2"
)

func main() {
	m := macaron.Classic()
	//insert()
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		key := "stuff"
		val := req.FormValue(key)
		fmt.Println("Value: " + val)
		res.Header().Set("Content-Type", "text/html")
		io.WriteString(res, `<form method="POST"><input type="text" name="stuff"><input type="submit"></form>`)

	})

	m.Run()
}

type (
	User struct {
		Username string
		Password string
	}
)

//adapted from https://stevenwhite.com/building-a-rest-service-with-golang-3/ used to make connection to mongoDB database

func insert() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://test:test@ds035006.mlab.com:35006/heroku_lzbj5rj0")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}
	c := s.DB("heroku_lzbj5rj0").C("Users")
	err = c.Insert(&User{"Ale", "+55 53 8116 9639"},
		&User{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	return s
}
