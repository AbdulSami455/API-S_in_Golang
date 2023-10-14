package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Portfolio struct {
	PID           string `json:"PID"`
	PortfolioName string `json:"portfolioname"`
	annualprofit  string `json:"profit"`
}

var portfs = []Portfolio{
	{PID: "1", PortfolioName: "Tech Top 10 Companies", annualprofit: "34"},
	{PID: "2", PortfolioName: "Agriculture", annualprofit: "12"},
}

func getportf(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, portfs)

}

func main() {

	router := gin.Default()
	router.GET("/books", getportf)
	router.Run("localhost:7070")
}
