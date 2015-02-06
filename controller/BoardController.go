package controller

import (
	"fmt"
	"html/template"
	"net/http"
)

func BoardIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("보드에서 인덱스 페이지 접속")
	t, err := template.ParseFiles("view/board/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(w, nil)
}
