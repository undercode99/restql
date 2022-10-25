package http_rest

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/undercode99/sqlt/db"
// 	"github.com/undercode99/sqlt/sqltgo"
// )

// func (a *Api) queryTemplateEngine(c *gin.Context) {

// 	urlQuery := c.Request.URL.Query()

// 	dataParam := map[string]interface{}{}
// 	for key, value := range urlQuery {
// 		if len(value) == 1 {
// 			dataParam[key] = value[0]
// 		}
// 	}

// 	sqltgoTemplate := `
// 		SELECT id, uuid::text, fullname FROM customer
// 		{{ if .offset }}OFFSET {{ .offset }} {{ end }}
// 		{{ if .limit }}LIMIT {{ .limit }} {{end}};
// 	`

// 	t, err := sqltgo.NewParse("getUser", sqltgoTemplate, dataParam)
// 	if err != nil {
// 		// response error to client
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 			"sql":   t.SqlRaw,
// 		})
// 		return
// 	}

// 	doQuery := db.NewQuery(a.db.GetDB(), t.SqlRaw)
// 	err = doQuery.Do()
// 	if err != nil {
// 		// response error to client
// 		c.JSON(500, gin.H{
// 			"error": err.Error(),
// 			"sql":   t.SqlRaw,
// 		})
// 		return
// 	}

// 	// response to client
// 	c.JSON(200, gin.H{
// 		"data":    doQuery.GetResultMap(),
// 		"sql":     t.SqlRaw,
// 		"columns": doQuery.GetMapColumnNameAndType(),
// 	})

// }
