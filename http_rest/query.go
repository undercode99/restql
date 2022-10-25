package http_rest

import (
	"github.com/gin-gonic/gin"
	"github.com/undercode99/restql/db"
)

// bind the query to the route
type BindJson struct {
	Query string `json:"query" binding:"required"`
}

func (a *Api) query(c *gin.Context) {

	// bind the query to the route
	var bindJson BindJson
	if err := c.ShouldBindJSON(&bindJson); err != nil {
		c.JSON(400, gin.H{"error": "Query json is required"})
		return
	}

	// do query
	doQuery := db.NewQuery(a.db.GetDB(), bindJson.Query)
	err := doQuery.Do()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// response to client
	c.JSON(200, gin.H{
		"data":    doQuery.GetResultMap(),
		"sql":     bindJson.Query,
		"columns": doQuery.GetMapColumnNameAndType(),
	})
}
