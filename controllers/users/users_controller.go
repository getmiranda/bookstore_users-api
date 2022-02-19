package users

import (
	"net/http"
	"strconv"

	"github.com/getmiranda/bookstore_users-api/domain/users"
	"github.com/getmiranda/bookstore_users-api/services/users_service"
	"github.com/getmiranda/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (uint64, errors.APIError) {
	userId, err := strconv.ParseUint(userIdParam, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestError("user id should be a valid number")
	}
	return userId, nil
}

func Get(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	user, getErr := users_service.Users.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
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

	result, err := users_service.Users.CreateUser(&user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Update(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
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

	result, err := users_service.Users.UpdateUser(isPartial, &user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result.Marshal(c.GetHeader("X-Public") == "true"))
}

func Delete(c *gin.Context) {
	userId, err := getUserId(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	if err := users_service.Users.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := users_service.Users.Search(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, users.Marshal(c.GetHeader("X-Public") == "true"))
}
