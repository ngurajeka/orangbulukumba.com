package middlewares

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

// RegisterDB Middleware for mongo Session
func RegisterDB(session *mgo.Session) gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Set("db", session)
		c.Next()
	}
}

// EnableCORS Middleware for Enable CORS
func EnableCORS() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Range, Content-Disposition, Content-Type, Authorization")
		c.Header("Access-Control-Allow-Methods", "GET,PATCH,PUT,POST,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Origin", "http://192.168.1.101:8100")
		c.Next()
	}
}
