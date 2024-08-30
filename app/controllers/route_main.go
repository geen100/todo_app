package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("app/views/templets/top.html")
	if err != nil{
		log.Fatalln(err)
	}
	t.Execute(w, "Hello")
}
