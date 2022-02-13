package users

import (
	"net/http"
	"strconv"

	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/services"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := errors.NewBadRequestError("user id should be a valid number")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		apiErr := errors.NewBadRequestError("invalid json body")
		c.JSON(apiErr.Status(), apiErr)
		return
	}

	result, err := services.CreateUser(&user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func UpdateUser(c *gin.Context) {
	userId, userErr := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if userErr != nil {
		apiErr := errors.NewBadRequestError("user id should be a valid number")
		c.JSON(apiErr.Status(), apiErr)
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

	result, err := services.UpdateUser(isPartial, &user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result)
}
