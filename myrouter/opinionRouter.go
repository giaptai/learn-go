package myrouter

import (
	"firm.com/function"
	"firm.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetOpinions(c *gin.Context) {
	var opinions []models.Opinion
	opinions, err := function.GetOpinions()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, opinions)
}

func GetFirmOpinions(c *gin.Context) {
	id := c.Param("id")
	var opinions []models.Opinion
	opinions, err := function.GetFirmOpinions(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, opinions)
}

func GetOpinionById(c *gin.Context) {
	id := c.Param("id") //get path variable
	// var intInt int
	fmt.Println(id)
	var opinion models.Opinion
	// if _, err := fmt.Sscan(id, &intInt); err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	opinion, err := function.GetSingleOpinion(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, opinion)
}

func AddOpinion(c *gin.Context) {
	var opinion models.Opinion
	if err := c.Bind(&opinion); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	_, err := function.InsertOpinion(&opinion)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, opinion)
}

func EditOpinion(c *gin.Context) {
	var opinion models.Opinion
	// Bind dữ liệu JSON từ request body vào struct firm
	if err := c.BindJSON(&opinion); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.UpdateOpinion(&opinion)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, opinion)
}

func RemoveOpinion(c *gin.Context) {
	id := c.Param("id") //get path variable
	err := function.DeleteOpinion(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "REMOVE " + id + " SUCCESSFULLY"})
}
