package main

import (
	"fmt"
	"log"
	"net/http"
)

func wang(w http.ResponseWriter,r *http.Request)  {
	fmt.Fprintf(w,"欢迎来到我世界")
}

func main() {
	http.HandleFunc("/",wang)
	err :=http.ListenAndServe(":5555",nil)
	if err != nil{
		log.Fatal("ListenAndServe",err)
	}
}