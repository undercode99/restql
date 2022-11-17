package http_rest

import (
	"github.com/gin-gonic/gin"
	"github.com/undercode99/restql/db"
)

// bind the query to the route
type BindJson struct {
	NameConnection string `json:"name_connection" binding:"required"`
	Query          string `json:"query" binding:"required"`
}

func (a *Api) query(c *gin.Context) {

	// bind the query to the route
	var bindJson BindJson
	if err := c.ShouldBindJSON(&bindJson); err != nil {
		c.JSON(400, gin.H{"error": "Query json is required"})
		return
	}

	// check if the connection already exists
	if !a.listDatabaseConnection.CheckExist(bindJson.NameConnection) {
		c.JSON(400, gin.H{"error": "Connection not found"})
		return
	}

	// get the connection
	con, err := a.listDatabaseConnection.GetConnection(bindJson.NameConnection)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// do query
	doQuery := db.NewQuery(con.GetDB(), bindJson.Query)
	err = doQuery.Do()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// response to client
	c.JSON(200, gin.H{
		"data":    doQuery.GetResultMapRows(),
		"sql":     bindJson.Query,
		"columns": doQuery.GetMapColumnNameAndType(),
	})
}

func (a *Api) addDatabaseConnection(c *gin.Context) {

	var bindJson db.Database

	if err := c.ShouldBindJSON(&bindJson); err != nil {
		c.JSON(400, gin.H{"error": "Database connection json is required"})
		return
	}

	// check if the connection already exists
	if a.listDatabaseConnection.CheckExist(bindJson.Name) {
		c.JSON(400, gin.H{"error": "Database connection already exists"})
		return
	}

	con, err := db.NewDatabaseConnect(&bindJson)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	a.listDatabaseConnection.AddConnection(bindJson.Name, con)

	c.JSON(200, gin.H{"message": "Database connection added"})
}

func (a *Api) listDatabaseConnections(c *gin.Context) {
	c.JSON(200, gin.H{"data": a.listDatabaseConnection.GetList()})
}
