package v1

import (
	"context"
	"fmt"
	"net/http"
	"tasks/Instagram_clone/insta_api/api/handlers/models"
	pp "tasks/Instagram_clone/insta_api/genproto/post_proto"
	"time"

	"github.com/gin-gonic/gin"
)

// CreatePostAs godoc
// @Summary Create New Post
// @Description This method create new post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.CreatePostReq true "Create Post"
// @Success 201 {object} post_proto.GetPostRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/create [POST]
func (h *handlerV1) CreatePost(c *gin.Context) {
	body := pp.CreatePostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().CreatePost(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// GetPostAs godoc
// @Summary Get Post
// @Description This method for Get Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.GetPostReq true "Get Post"
// @Success 200 {object} post_proto.GetPostRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/get [PUT]
func (h *handlerV1) GetPost(c *gin.Context) {
	body := pp.GetPostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().GetPost(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// UpdatePostAs godoc
// @Summary Update Post
// @Description This method for Update Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.UpdatePostReq true "Update Post"
// @Success 200 {object} post_proto.GetPostRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/update [PUT]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	body := pp.UpdatePostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().UpdatePost(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeletePostAs godoc
// @Summary Delete Post
// @Description This method for Delete Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.DeletePostReq true "Delete Post"
// @Success 200 {object} post_proto.Message
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/delete [PUT]
func (h *handlerV1) DeletePost(c *gin.Context) {
	body := pp.DeletePostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().DeletePost(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ListUserPostsAs godoc
// @Summary List user posts
// @Description This method for list user posts
// @Tags Post
// @Accept json
// @Produce json
// @Param id path string true "user_id"
// @Success 200 {object} post_proto.ListPostsRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /list/{id}/posts [GET]
func (h *handlerV1) ListUserPosts(c *gin.Context) {
	user_id := c.Param("id")
	body := pp.ListPostsReq{UserId: user_id}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().ListUserPosts(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// LikePostAs godoc
// @Summary Like Post
// @Description This method for Like Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.LikePostReq true "Like Post"
// @Success 200 {object} post_proto.Bool
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/like [POST]
func (h *handlerV1) Like(c *gin.Context) {
	body := pp.LikePostReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().Like(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, res)
}

// DeleteLikeAs godoc
// @Summary Delete Like
// @Description This method for Delete Like
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body post_proto.LikeDeleteReq true "Delete Like"
// @Success 200 {object} post_proto.Bool
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/like [DELETE]
func (h *handlerV1) DeleteLike(c *gin.Context) {
	body := pp.LikeDeleteReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, models.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.PostService().DeleteLike(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, models.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
