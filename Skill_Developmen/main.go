package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Flight struct {
	Id          string
	Number      string
	ArilineName string
	Source      string
	Destination string
	Capacity    int
	Price       float32
}

func readAllFlight(c *gin.Context) {
	flights := []Flight{
		{Id: "1001", Number: "AI 786", ArilineName: "Air India", Source: "Mumbai",
			Destination: "Abu Dhabi", Capacity: 100,
			Price: 15000.0},
		{Id: "1001", Number: "AI 706", ArilineName: "Air India", Source: "Abu Dhabi",
			Destination: "Mumbai", Capacity: 100,
			Price: 15000.0},
	}
	c.JSON(http.StatusOK, flights)
}

func readAllFlightById(c *gin.Context) {
	id := c.Param("id")
	flight := Flight{Id: id, Number: "AI 786", ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 100,
		Price: 15000.0}

	c.JSON(http.StatusOK, flight)
}
func createFlight(c *gin.Context) {
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	createdFlight := Flight{Id: "1006", Number: "AI 790", ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 100,
		Price: 15000.0}
	c.JSON(http.StatusCreated, gin.H{"message": "Flight Created Sucessfully.", "flight": createdFlight})
}

func updateFlight(c *gin.Context) {
	id := c.Param("id")
	var jbodyFlight Flight
	err := c.BindJSON(&jbodyFlight)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Server Error." + err.Error()})
		return
	}
	updatedFlight := Flight{Id: id, Number: "AI 790", ArilineName: "Air India", Source: "Mumbai",
		Destination: "Abu Dhabi", Capacity: 100,
		Price: 15000.0}
	c.JSON(http.StatusOK, gin.H{"message": "Flight Updated Sucessfully.", "flight": updatedFlight})
}

func deleteFlight(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Flight Deleted Sucessfully."})
}
func main() {
	/*flight := Flight{Id:"1001",Number: "AI 786",ArilineName: "Air India",Source: "Mumbai"}*/
	r := gin.Default()
	r.GET("/flights", readAllFlight)
	r.GET("/flights/:id", readAllFlightById)
	r.POST("/flights", createFlight)
	r.PUT("/flights/:id", updateFlight)
	r.DELETE("/flights/:id", deleteFlight)
	r.Run(":8080")
}

/*fmt.Println(cars)
var car1 Car = Car{Id: 101, Number: "TN48 z7878", Model: "Ambasiter", Type: "Seden"}
fmt.Println(car1)
fmt.Println(car1.Type)

var cars []Car = []Car{
	{Id: 101, Number: "TN48 z7878", Model: "Ambasiter", Type: "Seden"},
	{Id: 107, Number: "TN48 z1234", Model: "Toyoto", Type: "SUV"},*/
/*fmt.Println("Hello world")
var first int = 10
var second int = 20
var sum int = first + second
fmt.Println(sum)*/
