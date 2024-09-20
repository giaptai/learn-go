package main

import (
	// "fmt"
	// "io"
	"firm.com/connectDB"
	"firm.com/myrouter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "net/http"
)

func main() {
	// 
	connectdb.ConnMySql()
	connectdb.ConnMongoDB()
	// 
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5500"}
	r.Use(cors.New(config))
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
	//OPINION
	r.GET("/opinion", myrouter.GetOpinions)
	r.GET("/:id/opinion", myrouter.GetFirmOpinions)
	r.GET("/opinion/:id", myrouter.GetOpinionById)
	r.POST("/opinion", myrouter.AddOpinion)
	r.PUT("/opinion", myrouter.EditOpinion)
	r.DELETE("/opinion/:id", myrouter.RemoveOpinion)
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
