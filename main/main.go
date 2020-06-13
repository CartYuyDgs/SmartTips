package main

import (
	"SmartTips/logic"
	"encoding/json"
	"github.com/unrolled/render"
	"log"
	"net/http"
	"time"
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

	Init()

	//加载学校数据
	http.HandleFunc("/index", handlerIndex)
	//接口
	http.HandleFunc("/school/search", handlerSearch)

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
	r.ParseForm()
	keyword := r.FormValue("keyword")
	start := time.Now().UnixNano()
	//schools := logic.SearchV2(keyword, 16)
	schools := logic.Search(keyword, 16)
	end := time.Now().UnixNano()
	log.Printf("keyword : %s , cost : %d us\n", keyword, (end-start)/1000)
	responseSuccess(w, schools)
}

func responseSuccess(w http.ResponseWriter, data interface{}) {
	m := make(map[string]interface{}, 16)
	m["code"] = 0
	m["message"] = "success"
	m["data"] = data

	res, err := json.Marshal(m)
	if err != nil {
		return
	}

	w.Write(res)
}
