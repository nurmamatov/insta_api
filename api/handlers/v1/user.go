package v1

import (
	"fmt"
	"net/http"
	"tasks/Instagram_clone/insta_api/api/handlers/models"

	"github.com/gin-gonic/gin"
)

// RegisterUserAs godoc
// @Summary Create New User
// @Description This method register new user
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.CreateUserReq true "Register User"
// @Success 201 {object} models.GetUserRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /register [POST]
func (h *handlerV1) RegisterUser(c *gin.Context) {
	body := models.CreateUserReq{}
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

// LogInAs godoc
// @Summary Login User
// @Description This method for Login
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.LoginReq true "Login"
// @Success 201 {object} models.GetUserRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /login [GET]
func (h *handlerV1) Login(c *gin.Context) {
	body := models.LoginReq{}
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

// UpdateAs godoc
// @Summary Update User
// @Description This method for Update
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.UpdateUserReq true "Update"
// @Success 201 {object} models.GetUserRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/update [PUT]
func (h *handlerV1) Update(c *gin.Context) {
	body := models.UpdateUserReq{}
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

// DeleteAs godoc
// @Summary Delete User
// @Description This method for Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.DeleteUserReq true "Delete"
// @Success 201 {object} models.Message
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/delete [PUT]
func (h *handlerV1) Delete(c *gin.Context) {
	body := models.DeleteUserReq{}
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

// SearchAs godoc
// @Summary Search User
// @Description This method search from users
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.SearchUserReq true "Search"
// @Success 201 {object} models.UserList
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /search/user [GET]
func (h *handlerV1) Search(c *gin.Context) {
	body := models.SearchUserReq{}
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

// GetUserPostsAs godoc
// @Summary Get User Posts
// @Description This method for Get User Posts
// @Tags User
// @Accept json
// @Produce json
// @Param todo body models.GetUserReq true "Get User Posts"
// @Success 201 {object} models.GetUserRes
// @Failure 401 {object} models.Err
// @Failure 500 {object} models.Err
// @Router /user/get [GET]
func (h *handlerV1) GetUserPosts(c *gin.Context) {
	body := models.GetUserReq{}
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
