package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userData struct {
	ID   string `json:"ID"`
	MAIL string `json:"MAIL"`
	PASS string `json:"PASS"`
	USER string `json:"USER"`
	PREM bool   `json:"PREM"`
}

var DataStructure = []userData{
	{ID: "1", MAIL: "1111@mail.com", PASS: "1", USER: "111", PREM: true},
	{ID: "2", MAIL: "2@mail.com", PASS: "2", USER: "22", PREM: true},
	{ID: "3", MAIL: "3@mail.com", PASS: "3", USER: "33", PREM: false},
	{ID: "4", MAIL: "4@mail.com", PASS: "4", USER: "44", PREM: false},
	{ID: "5", MAIL: "5@mail.com", PASS: "5", USER: "55", PREM: true},
	{ID: "6", MAIL: "6@mail.com", PASS: "6", USER: "66", PREM: false},
	{ID: "7", MAIL: "7@mail.com", PASS: "7", USER: "77", PREM: true},
}

func main() {
	rout := gin.Default()
	rout.GET("/userData", getUser)
	rout.GET("/userData/:id", findUser)
	rout.POST("/userData", newUsers)
	rout.DELETE("/userData/:id", deleteUser)

	rout.Run("localhost:1337")
}

func getUser(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, DataStructure)
}

func newUsers(c *gin.Context) {
	var newUsers userData

	// newusers.
	if err := c.BindJSON(&newUsers); err != nil {
		return
	}
	DataStructure = append(DataStructure, newUsers)
	c.IndentedJSON(http.StatusCreated, newUsers)
}

func findUser(c *gin.Context) {
	id := c.Param("id")

	for _, a := range DataStructure {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Page Not Found"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range DataStructure {
		if user.ID == id {
			DataStructure = append(DataStructure[:i], DataStructure[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "User Deleted Successfully"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"Message": "User Not Found"})
}

// curl http://localhost:1337/userData \ --include \ --header "c" \ --request "POST" \ --data '{"ID": "7","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
