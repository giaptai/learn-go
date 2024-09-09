package main

import (
	"firm.com/myrouter"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/firm", myrouter.GetFirms)
	r.GET("/firm/:id", myrouter.GetFirmBMyId)
	r.POST("/firm", myrouter.AddFirm)
	r.PUT("/firm", myrouter.EditFirm)
	r.DELETE("/firm/:id", myrouter.RemoveFirm)
	r.Run("localhost:8771")
}
