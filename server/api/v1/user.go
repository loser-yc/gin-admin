package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	c.JSON(http.StatusOK, gin.H{
		"msg" : "success",
	})
}
