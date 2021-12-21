package forwarding

import (
	"code.meikeland.com/me/forwarding/internal/conf"
	"github.com/gin-gonic/gin"
	"github.com/meikeland/logger"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initRouter(router *gin.Engine) {
	// debug模式下打印请求日志
	logConf := conf.GetLogConfig()
	if logConf.Level == "debug" {
		router.Use(logger.LogRequest())
	}

	// api文档，地址http://localhost:8080/swagger/
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 其它路由
	router.GET("/api/index", index)
	router.POST("/api/getmsg", getMsg)

}

// @Summary Some DESC
// @Tags sample
// @Accept json
// @Produce json
// @Param name query string true "name"
// @Success 200 {string} string "result"
// @Router /index [get]
func index(c *gin.Context) {

}
