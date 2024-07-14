package web

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	quest "github.com/lafetz/quest-demo/services/quest/core"
)

type addQuestReq struct {
	KntUsername string `json:"KntUsername" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	Description string `json:"description" binding:"required"`
}
type getQuestsReq struct {
	KntUsername string `json:"KntUsername" binding:"required"`
}

func (app *App) addQuest(c *gin.Context) {
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

	questD := quest.NewQuest(questReq.Owner,
		questReq.KntUsername, questReq.Name, questReq.Description)
	qst, err := app.questService.AddQuest(c, *questD)
	if err != nil {
		if errors.Is(err, quest.ErrKntUnavailable) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		app.logger.Error("err", err, "stack")
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
	var questReq getQuestsReq
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
	quests, err := app.questService.GetAssignedQuests(c, questReq.KntUsername)
	if err != nil {
		app.logger.Error("err", err, "stack")
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
		if errors.Is(err, quest.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		app.logger.Error("err", err, "stack")
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
		if errors.Is(err, quest.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": err.Error(),
			})
			return
		}
		app.logger.Error("err", err, "stack")
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
