package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func ResponseFormatMiddleware(c *gin.Context) {
	if strings.Contains(c.Request.URL.Path, "/api") {
		c.Set("rsp_fmt", "json")
	} else {
		c.Set("rsp_fmt", "html")
	}
	c.Next()
}
