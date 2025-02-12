package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Game struct {
	Id              string
	Name            string
	NumberOfPlayers string
	Description     string
	Amount          float32
}

func readAllGame(c *gin.Context) {
	games := []Game{
		{Id: "1001", Name: "Kabbadi", NumberOfPlayers: "7 + 5", Amount: 50,
			Description: "Rule 1: 30 sec for raid,Rule2 :Weight under 65 kg"},
		{Id: "1002", Name: "ThrowBall", NumberOfPlayers: "9 + 3", Amount: 30,
			Description: "Rule 1:  25 points Win,"},
	}
	c.JSON(http.StatusOK, games)
}

func readAllGameById(c *gin.Context) {
	id := c.Param("id")
	game := Game{Id: id, Name: "Kabbadi", NumberOfPlayers: "7 + 5", Amount: 50,
		Description: "Rule 1: 30 sec for raid,Rule2 :Weight under 65 kg"}

	c.JSON(http.StatusOK, game)
}
func createGame(c *gin.Context) {
	var jbodyGame Game
	err := c.BindJSON(&jbodyGame)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdGame := Game{Id: "1001", Name: "Kho-kho", NumberOfPlayers: "9 + 3", Amount: 50,
		Description: "Rule 1: 5 minutes aset"}
	c.JSON(http.StatusCreated, gin.H{"message": "Game Created Sucessfully.", "flight": createdGame})
}

func deleteGame(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Game Deleted Sucessfully."})
}
func main() {
	/*flight := Flight{Id:"1001",Number: "AI 786",ArilineName: "Air India",Source: "Mumbai"}*/
	r := gin.Default()
	r.GET("/games", readAllGame)
	r.GET("/games/:id", readAllGameById)
	r.POST("/games", createGame)
	r.DELETE("/gamess/:id", deleteGame)
	r.Run()
}
