package user

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func InsertUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserModel
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := NewSQLStore(db)
		if err := store.Insert(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}

func UpdateUser(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserModel
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := NewSQLStore(db)
		if err := store.Update(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}

func GetAll(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		store := NewSQLStore(db)
		data, err := store.GetAll(c.Request.Context())
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)
	}
}
func GetById(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		store := NewSQLStore(db)
		data, err := store.GetOne(c, id)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, data)

	}
}
