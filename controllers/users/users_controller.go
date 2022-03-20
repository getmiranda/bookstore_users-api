package users

import (
	"net/http"
	"strconv"

	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/services"

	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (uint64, error) {
	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a valid number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		err, _ := getErr.(errors.APIError)
		c.JSON(err.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}

func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.UsersService.CreateUser(&user)
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	user.ID = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UsersService.UpdateUser(isPartial, &user)
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}

	if err := services.UsersService.DeleteUser(userId); err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UsersService.Search(status)
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, err := services.UsersService.LoginUser(&request)
	if err != nil {
		err, _ := err.(errors.APIError)
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshal(c.GetHeader("X-Public") == "true"))
}
