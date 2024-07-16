package web

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
func (a *App) initAppRoutes() {

	a.gin.GET("/quest/:questId", a.getQuest)
	a.gin.GET("/quest", a.getAssignedQuests)
	a.gin.PUT("/quest", a.completeQuest)
	a.gin.POST("/quest", a.addQuest)
	a.gin.GET("/metrics", prometheusHandler())
}
