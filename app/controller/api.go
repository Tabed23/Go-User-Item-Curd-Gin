package controller

import (
	"gin-crud/app/service"
	"gin-crud/app/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatUser(c *gin.Context) {

	var usr types.User
	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "invalid json",
		})
		return
	}
	resp, err := service.CreatUser(&usr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Cannot insert data into databsae",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": resp,
	})

}
func GetUser(c *gin.Context) {

}

func GetUsers(c *gin.Context) {
	usr, err := service.GetUsers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "there are no users",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   &usr,
	})
}

func UpdateUser(c *gin.Context) {

}

func DeleteUser(c *gin.Context) {

}
