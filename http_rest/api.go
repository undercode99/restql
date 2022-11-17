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
	apiConfig              *ApiConfig
	engine                 *gin.Engine
	listDatabaseConnection *db.ListDatabaseConnect
}

func NewApi(apiConfig *ApiConfig, listCon *db.ListDatabaseConnect) *Api {
	return &Api{
		apiConfig:              apiConfig,
		listDatabaseConnection: listCon,
	}
}

func (a *Api) Run() {

	// if a.apiConfig.Mode != "dev" {
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	a.engine = gin.Default()
	a.routes()

	a.engine.Run(":" + a.apiConfig.Port)
}
