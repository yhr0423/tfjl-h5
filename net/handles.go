package net

import (
	"github.com/gin-gonic/gin"
)

func WsHandler(c *gin.Context) {
	GWServer.Serve(c)
}
