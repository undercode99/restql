package http_rest

// ApiConfig is a struct that contains the configuration for the API

func (a *Api) routes() {
	// a.engine.GET("/api/v1/template/:name", a.queryTemplateEngine)
	a.engine.GET("/api/v1/query", a.query)
}
