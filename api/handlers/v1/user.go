package v1

import (
	"context"
	"fmt"
	"net/http"

	// "tasks/Instagram_clone/insta_api/api/handlers/user_proto"
	pu "tasks/Instagram_clone/insta_api/genproto/user_proto"
	"time"

	"github.com/gin-gonic/gin"
)

// RegisterUserAs godoc
// @Summary Create New User
// @Description This method register new user
// @Tags User
// @Accept json
// @Produce json
// @Param todo body user_proto.CreateUserReq true "Register User"
// @Success 201 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /register [POST]
func (h *handlerV1) RegisterUser(c *gin.Context) {
	body := pu.CreateUserReq{}
	
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error parse json"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()
	fmt.Println(body)
	res, err := h.serviceManager.UserService().CreateUser(ctx, &body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

// LogInAs godoc
// @Summary Login User
// @Description This method for Login
// @Tags User
// @Accept json
// @Produce json
// @Param todo body user_proto.LoginReq true "Login"
// @Success 201 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /login [GET]
func (h *handlerV1) Login(c *gin.Context) {
	body := pu.LoginReq{}
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
// @Param todo body user_proto.UpdateUserReq true "Update"
// @Success 201 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/update [PUT]
func (h *handlerV1) Update(c *gin.Context) {
	body := pu.UpdateUserReq{}
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
// @Param todo body user_proto.DeleteUserReq true "Delete"
// @Success 201 {object} user_proto.Message
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/delete [PUT]
func (h *handlerV1) Delete(c *gin.Context) {
	body := pu.DeleteUserReq{}
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
// @Param todo body user_proto.SearchUserReq true "Search"
// @Success 201 {object} user_proto.UserList
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /search/user [GET]
func (h *handlerV1) Search(c *gin.Context) {
	body := pu.SearchUserReq{}
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
// @Param todo body user_proto.GetUserReq true "Get User Posts"
// @Success 201 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/get [GET]
func (h *handlerV1) GetUserPosts(c *gin.Context) {
	body := pu.GetUserReq{}
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
