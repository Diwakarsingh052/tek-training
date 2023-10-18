package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	router.GET("/home/:name", Home)
	router.Run(":8080")
}

// func(*Context)
func Home(c *gin.Context) {

	name := c.Param("name")
	fmt.Println(name)
	//c.String(http.StatusOK, "this is my home page")
	//using the map to send the json response
	c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})
}
