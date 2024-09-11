package myrouter

import (
	"firm.com/function"
	"firm.com/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetFirms(c *gin.Context) {
	fmt.Print("ddd")
	var firms []models.VWFirm
	firms, err := function.GetFirm()
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, firms)
}

func GetFirmById(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	var firm models.VWFirm
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return //finish function
	}

	firm, err := function.GetSingleFirm(intInt)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(http.StatusOK, firm)
}

func AddFirm(c *gin.Context) {
	var firm models.VWFirm //create variable firm
	// Bind dữ liệu JSON từ request body vào struct firm
	//get data from body then mapping to firm by reference
	if err := c.BindJSON(&firm); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := function.InsertFirm(&firm)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, firm)
}

func EditFirm(c *gin.Context) {
	var firm models.VWFirm
	// Bind dữ liệu JSON từ request body vào struct firm
	if err := c.BindJSON(&firm); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(&firm.Name)
	_, err := function.UpdateFirm(&firm)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, firm)
}

func RemoveFirm(c *gin.Context) {
	id := c.Param("id") //get path variable
	var intInt int
	if _, err := fmt.Sscan(id, &intInt); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	err := function.DeleteFirm(intInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "REMOVE " + id + " SUCCESSFULLY"})
}
