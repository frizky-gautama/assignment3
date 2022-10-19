package controllers

import (
	"assignment3/server/models"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var min = 1
var max = 100

func GetWeathers(ctx *gin.Context) {
	r := rand.NewSource(time.Now().UnixNano())
	ctx.Header("content-type", "application/json")
	data := models.Status{
		Water: rand.New(r).Intn((max - min + 1) + min),
		Wind:  rand.New(r).Intn((max - min + 1) + min),
	}
	log.Printf("Measurement => Water: %v Wind: %v\n", data.Water, data.Wind)
	ctx.JSON(http.StatusOK, gin.H{
		"status": data,
	})
}
