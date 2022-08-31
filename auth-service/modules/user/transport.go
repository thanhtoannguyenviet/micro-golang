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
