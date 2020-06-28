package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irvanherz/aldebaran/internal/data"
	userM "github.com/irvanherz/aldebaran/internal/models/user"
)

func ReadOne(c *gin.Context) {
	_userId := c.Param("userId")
	userId, err0 := strconv.ParseInt(_userId, 10, 64)
	if err0 != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, data.ResponseData{
			Code:    "INVALID_INPUT",
			Message: "Invlid user id",
		})
	}
	res, err := userM.ReadById(userId)
	if err == "" && res != nil {
		res.Password = ""
		c.JSON(http.StatusOK, data.ResponseData{
			Code: "SUCCESS",
			Data: res,
		})
	} else if err == "" && res == nil {
		c.JSON(http.StatusOK, data.ResponseData{
			Code:    "SUCCESS",
			Message: "User doesn't exist",
		})
	} else {
		c.JSON(http.StatusInternalServerError, data.ResponseData{
			Code:    err,
			Message: "Failed reading database",
		})
	}
}

func ReadMany(c *gin.Context) {
	res, err := userM.ReadMany(1, 10, "")
	if err == "" && res != nil {
		for i := 0; i < len(*res); i++ {
			(*res)[i].Password = ""
		}
		c.JSON(http.StatusOK, data.ResponseData{
			Code: "SUCCESS",
			Data: res,
		})
	} else if err == "" && res == nil {
		c.JSON(http.StatusOK, data.ResponseData{
			Code:    "SUCCESS",
			Message: "User doesn't exist",
		})
	} else {
		c.JSON(http.StatusInternalServerError, data.ResponseData{
			Code:    err,
			Message: "Failed reading database",
		})
	}
}

func Update(c *gin.Context) {
	return
}

func Delete(c *gin.Context) {
	return
}
