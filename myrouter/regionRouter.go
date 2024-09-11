package myrouter

import (
	"firm.com/function"
	"firm.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRegions(c *gin.Context) {
	var regions []models.Region
	regions, err := function.GetRegions()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, regions)
}

func GetRegionById(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	var region models.Region
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	region, err := function.GetSingleRegion(intInt)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, region)
}

func AddRegion(c *gin.Context) {
	var region models.Region
	if err := c.BindJSON(&region); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.InsertRegion(&region)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, region)
}

func EditRegion(c *gin.Context) {
	var region models.Region
	// Bind dữ liệu JSON từ request body vào struct firm
	if err := c.BindJSON(&region); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.UpdateRegion(&region)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, region)
}

func RemoveRegion(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err := function.DeleteRegion(intInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "REMOVE " + id + " SUCCESSFULLY"})
}
