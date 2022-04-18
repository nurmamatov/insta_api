package v1

import (
	"context"
	"net/http"

	// "tasks/Instagram_clone/insta_api/api/handlers/user_proto"
	pu "tasks/Instagram_clone/insta_api/genproto/user_proto"
	l "tasks/Instagram_clone/insta_api/pkg/logger"
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
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Error parse json"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().CreateUser(ctx, &body)

	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, err)
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
// @Param todo body user_proto.LoginReq true "Login User"
// @Success 201 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /login [PUT]
func (h *handlerV1) Login(c *gin.Context) {
	body := pu.LoginReq{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusBadRequest, pu.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().Login(ctx, &body)

	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// ChangePassAs godoc
// @Summary Update Password
// @Description This method for Update password
// @Tags User
// @Accept json
// @Produce json
// @Param todo body user_proto.UpdatePass true "Password update"
// @Success 202 {object} user_proto.Message
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/password [PUT]
func (h *handlerV1) ChangePassword(c *gin.Context) {
	body := pu.UpdatePass{}
	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusBadRequest, pu.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().UpdatePassword(ctx, &body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// UpdateAs godoc
// @Summary Update User
// @Description This method for Update
// @Tags User
// @Accept json
// @Produce json
// @Param todo body user_proto.UpdateUserReq true "Update"
// @Success 200 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/update [PUT]
func (h *handlerV1) Update(c *gin.Context) {
	body := pu.UpdateUserReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusBadRequest, pu.Err{Error: err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().UpdateUser(ctx, &body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// DeleteAs godoc
// @Summary Delete User
// @Description This method for Delete user
// @Tags User
// @Accept json
// @Produce json
// @Param todo body user_proto.DeleteUserReq true "Delete"
// @Success 200 {object} user_proto.Message
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/delete [PUT]
func (h *handlerV1) Delete(c *gin.Context) {
	body := pu.DeleteUserReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusBadRequest, pu.Err{Error: err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().DeleteUser(ctx, &body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchAs godoc
// @Summary Search User
// @Description This method search from users
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} user_proto.UserList
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /search/{username} [GET]
func (h *handlerV1) Search(c *gin.Context) {
	username := c.Param("username")
	body := pu.SearchUserReq{Username: username}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().SearchUser(ctx, &body)
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

// SearchAs godoc
// @Summary Get Posts
// @Description This method Get User posts
// @Tags User
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} user_proto.GetUserRes
// @Failure 401 {object} user_proto.Err
// @Failure 500 {object} user_proto.Err
// @Router /user/{username} [GET]
func (h *handlerV1) GetUserPosts(c *gin.Context) {
	username := c.Param("username")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(5))
	defer cancel()

	res, err := h.serviceManager.UserService().GetUser(ctx, &pu.GetUserReq{Username: username})
	if err != nil {
		h.log.Error("Error: ", l.Error(err))
		c.JSON(http.StatusInternalServerError, pu.Err{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
