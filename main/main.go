package main

import (
	"github.com/unrolled/render"
	"log"
	"net/http"
)

var (
	renderHtml *render.Render
)

func init() {
	option := render.Options{
		Directory:  "views",
		Extensions: []string{".html", ".tmpl"},
	}
	renderHtml = render.New(option)
}

func main() {

	//加载学校数据
	http.HandleFunc("/index", handlerIndex)
	//接口
	http.HandleFunc("/search", handlerSearch)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	renderHtml.HTML(w, http.StatusOK, "index", nil)
	//w.Write([]byte("hello world!"))
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {

}
