package http_rest

func (a *Api) routes() {
	a.engine.GET("/api/v1/query", a.query)
	a.engine.POST("/api/v1/db/connect", a.addDatabaseConnection)
	a.engine.GET("/api/v1/db/connect", a.listDatabaseConnections)
}
