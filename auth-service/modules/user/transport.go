package user

import (
	"auth/modules/helper"
	"auth/modules/token"
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
		data.Salt = helper.GenerateRandomSalt()
		data.Password = helper.HashPassword(data.Password, data.Salt)
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
func Login(db *mongo.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		secret := "JWTSECRET"
		var data *UserLogin
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		store := NewSQLStore(db)
		res, err := store.FindByEmail(c, data.Email)
		if err != nil {
			panic(err)
		}
		//Check password
		checkPassword := helper.ConstainsPassword(res.Password, data.Password, res.Salt)
		if checkPassword == true {
			store := token.NewTokenJWTProvider(secret)
			var token token.TokenPayload
			token.Role = "Admin"
			token.UserId = res.Id
			rs, err := store.Generate(token, 200000)
			if err != nil {
				c.JSON(http.StatusBadRequest, err)
			} else {
				c.JSON(http.StatusOK, rs)
			}
		}

		c.JSON(http.StatusOK, checkPassword)

	}
}
