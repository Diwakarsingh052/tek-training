package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Id uint64

	Name string
}

var students = map[uint64]Student{

	101: Student{
		Id:   101,
		Name: "jeevan",
	},

	102: Student{
		Id:   102,
		Name: "afthab",
	},
}

func main() {

	router := gin.Default()

	router.GET("/home/:user_id", Home)
	router.Run(":8080")

}

// func(*Context)

func Home(c *gin.Context) {

	Stringid := c.Param("user_id")
	fmt.Println(Stringid)
	uid, err := strconv.ParseUint(Stringid, 10, 64)

	if err != nil {

		log.Println("conversion string to int error", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "error found at conversion.."})
		return

	}

	val, err := fetchStudent(uid)

	if err != nil {
		log.Println("student not found in map", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "students not found"})
		return
	}

	c.JSON(http.StatusOK, val)

	// fmt.Println(name)

	//c.String(http.StatusOK, "this is my home page")

	//using the map to send the json response

	// c.JSON(http.StatusOK, gin.H{"msg": "this is my home page"})

	// c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{})

}

func fetchStudent(uid uint64) (Student, error) {

	sData, ok := students[uid]

	if !ok {
		return Student{}, errors.New("data not there")
	}

	return sData, nil

}
