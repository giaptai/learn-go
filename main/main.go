package main

import (
	// "fmt"
	// "io"
	"firm.com/myrouter"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	r := gin.Default()
	//FIRM
	r.GET("/firm", myrouter.GetFirms)
	r.GET("/firm/:id", myrouter.GetFirmById)
	r.POST("/firm", myrouter.AddFirm)
	r.PUT("/firm", myrouter.EditFirm)
	r.DELETE("/firm/:id", myrouter.RemoveFirm)
	//REGION
	r.GET("/region", myrouter.GetRegions)
	r.GET("/region/:id", myrouter.GetRegionById)
	r.POST("/region", myrouter.AddRegion)
	r.PUT("/region", myrouter.EditRegion)
	r.DELETE("/region/:id", myrouter.RemoveRegion)
	//BIZ SIZE
	r.GET("/biz-size", myrouter.GetBizSizes)
	r.GET("/biz-size/:id", myrouter.GetBizSizeById)
	r.POST("/biz-size", myrouter.AddBizSize)
	r.PUT("/biz-size", myrouter.EditBizSize)
	r.DELETE("/biz-size/:id", myrouter.RemoveBizSize)
	//VERIFY
	r.POST("/login", myrouter.Login)
	r.POST("/register", myrouter.Register)
	// resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	// if err!=nil{
	// 	return
	// }
	// defer resp.Body.Close()
	// body, _ := io.ReadAll(resp.Body)
	// fmt.Printf("%s", body)
	//
	r.Run("localhost:8771")
}
