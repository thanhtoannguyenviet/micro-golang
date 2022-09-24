package token

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Generate(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}
}
func Validate(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Payload struct {
			Payload string `json:"payload"`
		}
		var data Payload
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := NewTokenJWTProvider(secret)
		rs, err := store.Validate(data.Payload)
		if err != nil {
			c.JSON(http.StatusNonAuthoritativeInfo, nil)
		}
		c.JSON(http.StatusOK, rs)

	}
}
