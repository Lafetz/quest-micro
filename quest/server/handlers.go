package questserver

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	commonerrors "github.com/lafetz/quest-micro/common/errors"
	quest "github.com/lafetz/quest-micro/quest/core"
)

type addQuestReq struct {
	Email       string `json:"email" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Owner       string `json:"owner" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (app *App) addQuest(c *gin.Context) {
	var questReq addQuestReq
	if err := c.ShouldBind(&questReq); err != nil {

		_, ok := err.(validator.ValidationErrors)

		if ok {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   "There was a problem with the provided data",
				"details": ValidateModel(err),
			})
			return

		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error processing request body",
		})
		return
	}

	questD := quest.NewQuest(questReq.Owner,
		questReq.Email, questReq.Name, questReq.Description)
	qst, err := app.questService.AddQuest(c, *questD)
	if err != nil {
		if errors.Is(err, quest.ErrKntUnavailable) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		if errors.Is(err, commonerrors.ErrKnightNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		app.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server Error",
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
	kntName := c.Query("kntName")

	if kntName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "kntName query parameter is required",
		})
		return
	}
	quests, err := app.questService.GetAssignedQuests(c, kntName)
	if err != nil {
		app.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server Error",
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
			"error": "quest not found",
		})
		return
	}
	err = app.questService.CompleteQuest(c, id)
	if err != nil {
		if errors.Is(err, quest.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		app.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server Error",
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
			"error": "quest not found",
		})
		return
	}
	qst, err := app.questService.GetQuest(c, id)
	if err != nil {
		if errors.Is(err, quest.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		app.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server Error",
		})
		return
	}
	res := toJsonQuest(qst)
	c.JSON(http.StatusOK, gin.H{
		"quest": res,
	})
}
