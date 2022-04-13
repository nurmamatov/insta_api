package v1

import (
	"fmt"
	"net/http"
	"tasks/Instagram_clone/insta_api/api/handlers/models"

	"github.com/gin-gonic/gin"
)

// CreateCommentAs godoc
// @Summary Create New Comment
// @Description This method create new comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body models.CreateCommentReq true "Create Comment"
// @Success 201 {object} models.GetCommentRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/create [POST]
func (h *handlerV1) CreateComment(c *gin.Context) {
	body := models.CreateCommentReq{}
	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "salome emas",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "salom",
	})
}

// GetCommentAs godoc
// @Summary Get Comment
// @Description This method Get comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body models.GetCommentReq true "Get Comment"
// @Success 201 {object} models.GetCommentRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/get [GET]
func (h *handlerV1) GetComment(c *gin.Context) {
	body := models.GetCommentReq{}
	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "salome emas",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "salom",
	})
}

// UpdateCommentAs godoc
// @Summary Update Comment
// @Description This method Update comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body models.UpdateCommentReq true "Update Comment"
// @Success 201 {object} models.GetCommentRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/update [PUT]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	body := models.UpdateCommentReq{}
	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "salome emas",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "salom",
	})
}

// DeleteCommentAs godoc
// @Summary Delete Comment
// @Description This method Delete comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body models.DeleteCommentReq true "Delete Comment"
// @Success 201 {object} models.Empty
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/delete [PUT]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	body := models.DeleteCommentReq{}
	err := c.ShouldBindJSON(&body)
	fmt.Println(body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"message": "salome emas",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "salom",
	})
}
