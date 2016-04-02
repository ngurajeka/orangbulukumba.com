package views

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ngurajeka/orangbulukumba.com/forms"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Users in orangbulukumba
func Users(c *gin.Context) {

	var users []forms.User
	session := c.MustGet("db").(*mgo.Session)
	col := session.DB("orangbulukumba").C("users")

	err := col.Find(bson.M{"deleted": false}).All(&users)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "Data Fetched",
	})

}

// User in orangbulukumba
func User(c *gin.Context) {

	username := c.Param("username")
	if username == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Username",
		})
		return
	}

	var user forms.User
	session := c.MustGet("db").(*mgo.Session)
	col := session.DB("orangbulukumba").C("users")

	err := col.Find(bson.M{"username": username, "deleted": false}).One(&user)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user: " + username + " " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    user,
		"message": "Data Fetched",
	})

}

// NewUser in orangbulukumba
func NewUser(c *gin.Context) {

	var user forms.User

	session := c.MustGet("db").(*mgo.Session)

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// TODO What if user already exist?
	// find existing user first

	// default user deleted -> false
	user.Deleted = false
	// default created time
	user.Created = time.Now()

	col := session.DB("orangbulukumba").C("users")

	err = col.Insert(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}

// UpdateUser in orangbulukumba
func UpdateUser(c *gin.Context) {

	var userDb forms.User
	var userUpdate forms.User

	session := c.MustGet("db").(*mgo.Session)

	// param username
	username := c.Param("username")

	if username == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Id",
		})
		return
	}

	// finding user
	col := session.DB("orangbulukumba").C("users")
	err := col.Find(bson.M{"username": username, "deleted": false}).One(&userDb)
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user: " + username + ", " + err.Error(),
		})
		return
	}

	// binding form-data / json-data
	err = c.Bind(&userUpdate)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = col.Update(&userDb, userUpdate)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    userDb,
		"message": "Data Updated",
	})
}

// DeleteUser in orangbulukumba
func DeleteUser(c *gin.Context) {

	session := c.MustGet("db").(*mgo.Session)

	var user forms.User

	// param username
	username := c.Param("username")

	if username == "" {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty Username",
		})
		return
	}

	// finding user
	col := session.DB("orangbulukumba").C("users")
	err := col.Find(bson.M{"username": username, "deleted": false}).One(&user)

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"error": "user: " + username + ", " + err.Error(),
		})
		return
	}

	// flag as deleted
	userUpdate := forms.User{
		Deleted: true,
	}

	err = col.Update(&user, userUpdate)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Data Deleted",
	})
}
