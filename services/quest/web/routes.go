package web

func (a *App) initAppRoutes() {

	a.gin.GET("/quest/:questId", a.getQuest)
	a.gin.GET("/quest", a.getAssignedQuests)
	a.gin.PUT("/quest", a.completeQuest)
	a.gin.POST("/quest", a.addQuest)
}
