package v1

import (
	"fmt"
	"net/http"
	"tasks/Instagram_clone/insta_api/api/handlers/models"

	"github.com/gin-gonic/gin"
)

// CreatePostAs godoc
// @Summary Create New Post
// @Description This method create new post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.CreatePostReq true "Create Post"
// @Success 201 {object} models.GetPostRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/create [POST]
func (h *handlerV1) CreatePost(c *gin.Context) {
	body := models.CreatePostReq{}
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

// GetPostAs godoc
// @Summary Get Post
// @Description This method for Get Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.GetPostReq true "Get Post"
// @Success 201 {object} models.GetUserRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/get [GET]
func (h *handlerV1) GetPost(c *gin.Context) {
	body := models.GetPostReq{}
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

// UpdatePostAs godoc
// @Summary Update Post
// @Description This method for Update Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.UpdatePostReq true "Update Post"
// @Success 201 {object} models.GetPostRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/update [PUT]
func (h *handlerV1) UpdatePost(c *gin.Context) {
	body := models.UpdatePostReq{}
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

// DeletePostAs godoc
// @Summary Delete Post
// @Description This method for Delete Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.DeletePostReq true "Delete Post"
// @Success 201 {object} models.Message
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/delete [PUT]
func (h *handlerV1) DeletePost(c *gin.Context) {
	body := models.DeletePostReq{}
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

// ListUserPostsAs godoc
// @Summary List user posts
// @Description This method for list user posts
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.ListUserPosts true "List posts"
// @Success 201 {object} models.ListPostsRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /list/user/posts [GET]
func (h *handlerV1) ListUserPosts(c *gin.Context) {
	body := models.ListPostsReq{}
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

// LikePostAs godoc
// @Summary Like Post
// @Description This method for Like Post
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.LikePostReq true "Like Post"
// @Success 201 {object} models.Bool
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/like [POST]
func (h *handlerV1) Like(c *gin.Context) {
	body := models.LikePostReq{}
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

// DeleteLikeAs godoc
// @Summary Delete Like
// @Description This method for Delete Like
// @Tags Post
// @Accept json
// @Produce json
// @Param todo body models.LikeDeleteReq true "Delete Like"
// @Success 201 {object} models.Bool
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /post/like [DELETE]
func (h *handlerV1) DeleteLike(c *gin.Context) {
	body := models.LikeDeleteReq{}
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
