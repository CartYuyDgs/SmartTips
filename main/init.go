package main

import (
	"SmartTips/logic"
	"log"
)

func Init() {
	/*option := render.Options{
		Directory:  "views",
		Extensions: []string{".html", ".tmpl"},
	}
	renderHtml = render.New(option)
	*/
	err := logic.InitLogic("./data/school.dat")
	if err != nil {
		return
	}
	log.Println("init success!")
}
