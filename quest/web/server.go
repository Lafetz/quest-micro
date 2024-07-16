package web

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	quest "github.com/lafetz/quest-demo/quest/core"
)

type App struct {
	questService quest.QuestServiceApi
	gin          *gin.Engine
	port         int
	logger       *slog.Logger
}

func NewApp(questService quest.QuestServiceApi, port int, logger *slog.Logger) *App {
	a := &App{
		gin:          gin.Default(),
		questService: questService,
		port:         port,
		logger:       logger,
	}
	a.gin.Use(corsMiddleware())
	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return a.gin.Run()
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
