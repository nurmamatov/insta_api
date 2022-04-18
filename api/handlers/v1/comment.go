package v1

import (
	"context"
	"fmt"
	"net/http"
	"tasks/Instagram_clone/insta_api/api/handlers/models"
	pc "tasks/Instagram_clone/insta_api/genproto/comment_proto"
	"time"

	"github.com/gin-gonic/gin"
)

// CreateCommentAs godoc
// @Summary Create New Comment
// @Description This method create new comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body comment_proto.CreateCommentReq true "Create Comment"
// @Success 201 {object} comment_proto.Res
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/create [POST]
func (h *handlerV1) CreateComment(c *gin.Context) {
	body := pc.CreateCommentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.CommentService().CreateComment(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateCommentAs godoc
// @Summary Update Comment
// @Description This method Update comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body comment_proto.UpdateCommentReq true "Update Comment"
// @Success 201 {object} comment_proto.Res
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/update [PUT]
func (h *handlerV1) UpdateComment(c *gin.Context) {
	body := pc.UpdateCommentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.CommentService().UpdateComment(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteCommentAs godoc
// @Summary Delete Comment
// @Description This method Delete comment
// @Tags Comment
// @Accept json
// @Produce json
// @Param todo body comment_proto.DeleteCommentReq true "Delete Comment"
// @Success 201 {object} comment_proto.Empty
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /comment/delete [PUT]
func (h *handlerV1) DeleteComment(c *gin.Context) {
	body := pc.DeleteCommentReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.CommentService().DeleteComment(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
