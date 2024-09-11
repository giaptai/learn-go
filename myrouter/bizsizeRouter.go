package myrouter

import (
	"firm.com/function"
	"firm.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBizSizes(c *gin.Context) {
	var bizsizes []models.BizSize
	bizsizes, err := function.GetBizSizes()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, bizsizes)
}

func GetBizSizeById(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	var bizsize models.BizSize
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	bizsize, err := function.GetSingleBizSize(intInt)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bizsize)
}

func AddBizSize(c *gin.Context) {
	var bizsize models.BizSize
	if err := c.BindJSON(&bizsize); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.InsertBizSize(&bizsize) // why underscore because param is memory address
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, bizsize)
}

func EditBizSize(c *gin.Context) {
	var bizsize models.BizSize
	// Bind dữ liệu JSON từ request body vào struct firm
	if err := c.BindJSON(&bizsize); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.UpdateBizSize(&bizsize)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, bizsize)
}

func RemoveBizSize(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err := function.DeleteBizSize(intInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "REMOVE " + id + " SUCCESSFULLY"})
}
