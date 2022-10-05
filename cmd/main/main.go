package main

import (
	"net/http"

	util "github.com/Korppi/love/pkg/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/lovecheck", getLoveScore)
	router.Run("localhost:8080")
}

// Returns lovescore
func getLoveScore(c *gin.Context) {
	fname, _ := c.GetQuery("fname")
	sname, _ := c.GetQuery("sname")
	score, err := util.CalculateLove(fname, sname)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
	} else {
		c.IndentedJSON(http.StatusOK, score)
	}

}
