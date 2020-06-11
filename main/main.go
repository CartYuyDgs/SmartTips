package main

import (
	"log"
	"net/http"
)

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
	w.Write([]byte("hello world!"))
}

func handlerSearch(w http.ResponseWriter, r *http.Request) {

}
