package main

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func ResponseFormatMiddleware(c *gin.Context) {
	if strings.Contains(c.Request.URL.Path, "/api") {
		c.Set("rsp_fmt", "json")
	} else {
		c.Set("rsp_fmt", "html")
	}
	c.Next()
}

var DatabaseMiddleware = func(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
