package gintest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RunServer() error{
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	err := r.Run()
	if err != nil {
		return err
	} //

	return nil
}


