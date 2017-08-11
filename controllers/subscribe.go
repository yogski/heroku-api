package controllers

import (
	m "employee_subscription/models"
	"github.com/gin-gonic/gin"
	"net/http"
	_"strconv"
)

func SubsGet(c *gin.Context){

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

func SubsDetail(c *gin.Context){

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
	//id, err := strconv.Atoi(idStr)
	c.String(http.StatusOK, "Hello, your ID is %s", idStr)

	/*
	if  err != nil {
		response := &m.Response{
			Message: []string{err.Error()},
		}
		c.JSON(http.StatusBadRequest, response)
		c.Abort()
		return
	}

	user, err := m.GetUser(id)
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
	*/
}

func SubsCreate(c *gin.Context){

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
		Message: []string{"New subscription update"},
		Data: user,
	}

	c.JSON(http.StatusCreated, response)
}