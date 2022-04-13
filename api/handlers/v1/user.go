package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterUserAs godoc
// @Summary Creates new task
// @Description This method create new task
// @Security BearerAuth
// @Tags Task
// @Accept json
// @Produce json
// @Param todo body models.CreateTaskReq true "New Task"
// @Success 201 {object} models.TaskRes
// @Router /register [POST]
func (h *handlerV1) RegisterUser(c *gin.Context) {
	c.JSON(int(http.StateActive), gin.H{
		"message": "salom",
	})
}
