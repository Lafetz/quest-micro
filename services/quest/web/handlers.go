package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	quest "github.com/lafetz/quest-demo/services/quest/core"
)

type addQuestReq struct {
	KnightID    uuid.UUID `json:"knightId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func (app *App) addQuest(c *gin.Context) {
	//user, exists := c.Get("user")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"Error": "Unauthorized",
	// 	})
	// 	return
	// }
	var questReq addQuestReq
	if err := c.ShouldBind(&questReq); err != nil {
		_, ok := err.(validator.ValidationErrors)

		if ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"Errors": ValidateModel(err),
			})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error processing request body",
		})
		return
	}
	name := "sir. ironborn"
	questD := quest.NewQuest(name,
		questReq.KnightID, questReq.Name, questReq.Description)
	qst, err := app.questService.AddQuest(c, *questD)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "internal server Error",
		})
		return
	}
	res := toJsonQuest(qst)
	c.JSON(http.StatusCreated, gin.H{
		"msg":   "Quest added",
		"quest": res,
	})

}
func (app *App) getAssignedQuests(c *gin.Context) {

	id, err := uuid.Parse("temo user id")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "error parsing id uuid",
		})
	}

	quests, err := app.questService.GetAssignedQuests(c, id)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "internal server Error",
		})
		return
	}
	res := toJsonQuests(quests)
	c.JSON(http.StatusOK, gin.H{
		"quests": res,
	})
}
func (app *App) completeQuest(c *gin.Context) {
	questId := c.Param("questId")
	id, err := uuid.Parse(questId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "quest not found",
		})
		return
	}
	err = app.questService.CompleteQuest(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "internal server Error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": "quest completed",
	})
}
func (app *App) getQuest(c *gin.Context) {
	questId := c.Param("questId")
	id, err := uuid.Parse(questId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "quest not found",
		})
		return
	}
	qst, err := app.questService.GetQuest(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "internal server Error",
		})
		return
	}
	res := toJsonQuest(qst)
	c.JSON(http.StatusOK, gin.H{
		"quest": res,
	})
}
