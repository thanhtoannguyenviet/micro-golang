package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type sqlStore struct {
	db *mongo.Client
}

func NewSQLStore(db *mongo.Client) *sqlStore {
	return &sqlStore{db: db}
}

// Insert inserts a new user into the database, and returns the ID of the newly inserted row
func (s *sqlStore) Insert(ctx context.Context, data *UserModel) error {
	//db := s.db.Connect(ctx)
	collection := s.db.Database("micro-go").Collection("auth-user")
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
