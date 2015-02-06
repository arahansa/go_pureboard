package main

import (
	"fmt"
	"github.com/gorilla/mux"
	c "go_pureboard/controller"
	"net/http"
)

func main() {
	fmt.Println("Hello world Board Server")

	fs := http.FileServer(http.Dir("view/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	r := mux.NewRouter()
	r.HandleFunc("/board/{id}", handleRead)
	r.HandleFunc("/board", c.BoardIndex)
	r.HandleFunc("/", handler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello I will be a Board. %s", r.URL.Path[1:])
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Id := vars["id"]
	fmt.Println("id값은 ,", Id)
	fmt.Fprintf(w, "Hello I am Board.  게시글 넘버 : %s", Id)
}
