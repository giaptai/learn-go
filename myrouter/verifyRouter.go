package myrouter

import (
	// "fmt"
	"html"
	"time"
	"strings"
	"firm.com/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var a models.RegisterForm
	c.Bind(&a)
	// s, _ := Hash(a.PASS)
	claims:=jwt.MapClaims{
		"username": a.FULLNAME,
		"exp":time.Now().Add(time.Minute * 5).Unix(),
	}
	// claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("sadasdasdasdasdasdasdasdas343432423423sfs"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(202, gin.H{"token": tokenString})

	// err := CheckPasswordHash("$2a$14$U3tNFYINBByYA8LLlr646eZUKyYzq5nGsfiPHcVC1widim0hSYeY.", a.PASS)
	// if err != nil {
	// 	c.JSON(400, gin.H{"error": err.Error()})
	// 	return
	// }
	// c.JSON(400, gin.H{"message": "MATCH PASSWORD"})
}

func Register(c *gin.Context) {
	var a models.RegisterForm
	c.Bind(&a)
	s, _ := Hash(a.PASS)
	c.JSON(200, gin.H{
		"email":     a.EMAIL,
		"fulllname": a.FULLNAME,
		"pass":      s,
	})
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
