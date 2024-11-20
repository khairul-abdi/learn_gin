package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	r := router()
	r.Run(":9000")
}

func router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLFiles("index.html")

	{
		test := r.Group("/test")
		test.GET("/ping", GetPong)
	}

	{
		v1 := r.Group("/v1")
		v1.POST("/login", UserLogin)
		v1.GET("/user", GetUser)
		v1.GET("/index", LoadHtml)
	}

	return r
}

func LoadHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func GetUser(c *gin.Context) {
	var user = User{
		ID:   1,
		Name: "John Doe",
		Age:  30,
		Address: Address{
			Street:  "Jl. ABC",
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}

	c.JSON(http.StatusOK, user)
}

func UserLogin(c *gin.Context) {
	var login Login

	err := c.ShouldBindJSON(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// manual validation
	if login.Username == "" || login.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password required"})
		return
	}

	// anggap ada logic untuk mengecek data user login

	// dummy response berhasil login
	c.JSON(http.StatusOK, gin.H{
		"message": "success login",
	})
}
