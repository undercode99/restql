package http_rest

import (
	"github.com/gin-gonic/gin"
	"github.com/undercode99/restql/db"
)

type ApiConfig struct {
	Port string
	Mode string
}

type Api struct {
	apiConfig *ApiConfig
	engine    *gin.Engine
	db        *db.Database
}

func NewApi(apiConfig *ApiConfig, dbConfig *db.Database) *Api {
	return &Api{
		apiConfig: apiConfig,
		db:        dbConfig,
	}
}

func (a *Api) Run() {

	if a.apiConfig.Mode != "dev" {
		gin.SetMode(gin.ReleaseMode)
	}

	a.engine = gin.Default()
	a.routes()

	a.engine.Run(":" + a.apiConfig.Port)
}
