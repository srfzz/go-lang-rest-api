package auth

import (
	"net/http"

	"go-lang-restapi/models"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not Parse data 	" + err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "An Error Occurred 	" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "New User Created Successfully!"})
}
