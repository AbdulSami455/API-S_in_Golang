package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Portfolio struct {
	PID           string `json:"PID"`
	PortfolioName string `json:"portfolioname"`
	AnnualProfit  string `json:"profit"`
}

var portfs = []Portfolio{
	{PID: "1", PortfolioName: "Tech Top 10 Companies", AnnualProfit: "34"},
	{PID: "2", PortfolioName: "Agriculture", AnnualProfit: "12"},
	{PID: "3", PortfolioName: "Energy", AnnualProfit: "13"},
}

// Get all portfolios
func getPortfolios(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, portfs)
}

// Get a portfolio by ID
func getPortfolioByID(c *gin.Context) {
	id := c.Param("id")

	for _, p := range portfs {
		if p.PID == id {
			c.IndentedJSON(http.StatusOK, p)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
}

// Create a new portfolio
func createPortfolio(c *gin.Context) {
	var newPortfolio Portfolio

	if err := c.ShouldBindJSON(&newPortfolio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new PID (You may want to use a more robust method for generating IDs)
	newPortfolio.PID = strconv.Itoa(len(portfs) + 1)

	portfs = append(portfs, newPortfolio)

	c.JSON(http.StatusCreated, newPortfolio)
}

func main() {
	router := gin.Default()

	// Route to get all portfolios
	router.GET("/portfolios", getPortfolios)

	// Route to get a portfolio by ID
	router.GET("/portfolios/:id", getPortfolioByID)

	// Route to create a new portfolio
	router.POST("/portfolios", createPortfolio)

	router.Run("localhost:7080")
}
