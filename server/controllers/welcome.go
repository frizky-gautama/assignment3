package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Render(ctx *gin.Context, f string, datas map[string]interface{}) {
	tpl, err := template.ParseFiles(f)
	if err != nil {
		log.Printf("Error : %v\n", err)

		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.Header("Content-Type", "text/html")
	tpl.Execute(ctx.Writer, datas)
}
