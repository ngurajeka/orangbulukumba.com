package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ngurajeka/orangbulukumba.com/middlewares"
	"github.com/ngurajeka/orangbulukumba.com/views"
	"gopkg.in/mgo.v2"
)

func main() {

	h := gin.Default()

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	h.Use(middlewares.RegisterDB(session))
	h.Use(middlewares.EnableCORS())

	h.GET("/user", views.Users)
	h.POST("/user", views.NewUser)

	h.GET("/user/:username", views.User)
	h.PUT("/user/:username", views.UpdateUser)
	h.DELETE("/user/:username", views.DeleteUser)

	h.Run(":3030")
}
