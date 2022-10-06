package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	//http.HandleFunc("/note", noteHandlr)
	log.Print("Listening port: 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", http.FileServer(http.Dir("notes/"))))
}

//func noteHandlr(w http.ResponseWriter, r *http.Request) {
//	t, err := template.ParseFiles("notes/Robot Design and implementation.html")
//	if err != nil {
//		log.Print(err)
//	}
//	t.Execute(w, nil)
//}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.RequestURI() == "/favicon.ico" {
		return
	}
	fmt.Fprintf(w, "Welcome\n")
	fmt.Printf("%s %s %s\n", r.Method, r.URL, r.Proto)
	fmt.Printf("RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Printf("Form[%q] = %q\n", k, v)
	}
}
