package controllers

import (
	m "employee_subscription/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_"strconv"
)

func UserGet(c *gin.Context){

	authorization := c.Request.Header.Get("Authorization")
	if  authorization != "12345" {
		response := &m.Response{
			Message: []string{"Unauthorized access"},
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	users, err := m.GetUsers()
	if err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: []string{"Get users"},
		Data: users,
	}

	c.JSON(http.StatusOK, response)
}

func UserDetail(c *gin.Context){

	authorization := c.Request.Header.Get("Authorization")
	if  authorization != "12345" {
		response := &m.Response{
			Message: []string{"Unauthorized access"},
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	idStr := c.Param("id")
	/*id, err := strconv.Atoi(idStr)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}*/

	user, err := m.GetUser(idStr)
	if err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: []string{"Get user"},
		Data: user,
	}

	c.JSON(http.StatusOK, response)
}

func UserCreate(c *gin.Context){

	authorization := c.Request.Header.Get("Authorization")
	if  authorization != "12345" {
		response := &m.Response{
			Message: []string{"Unauthorized access"},
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	req := &m.User{}
	err := c.BindJSON(&req)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	user, err := m.CreateUser(req)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: []string{"User has been created"},
		Data: user,
	}

	c.JSON(http.StatusCreated, response)
}

func UserUpdate(c *gin.Context){

	authorization := c.Request.Header.Get("Authorization")
	if  authorization != "12345" {
		response := &m.Response{
			Message: []string{"Unauthorized access"},
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	idStr := c.Param("id")
	/*id, err := strconv.Atoi(idStr)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}*/

	req := &m.User{}
	err := c.BindJSON(&req)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	req.Phone = idStr
	user, err := m.UpdateUser(req)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: []string{"User has been updated"},
		Data: user,
	}

	c.JSON(http.StatusOK, response)
}

func UserDelete(c *gin.Context){

	authorization := c.Request.Header.Get("Authorization")
	if  authorization != "12345" {
		response := &m.Response{
			Message: []string{"Unauthorized access"},
		}
		c.JSON(http.StatusUnauthorized, response)
		c.Abort()
		return
	}

	idStr := c.Param("id")
	/*id, err := strconv.Atoi(idStr)
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}*/

	err := m.DeleteUser(idStr)
	if err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusServiceUnavailable, response)
		c.Abort()
		return
	}

	response := &m.Response{
		Message: []string{"User has been deleted"},
	}

	c.JSON(http.StatusOK, response)
}
