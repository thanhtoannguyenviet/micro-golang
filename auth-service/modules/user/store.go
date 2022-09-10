package user

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
func (s *sqlStore) GetAll(ctx context.Context) ([]*UserModel, error) {
	var rs []*UserModel
	collection := s.db.Database("micro-go").Collection("auth-user")
	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	res, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	if err = res.All(ctx, &rs); err != nil {
		panic(err)
	}
	return rs, nil
}
func (s *sqlStore) GetOne(ctx context.Context, id string) (*UserModel, error) {
	var rs *UserModel
	collection := s.db.Database("micro-go").Collection("auth-user")
	docId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err = collection.FindOne(ctx, bson.M{"_id": docId}).Decode(&rs); err != nil {
		return nil, err
	}
	return rs, err
}
func (s *sqlStore) Update(ctx context.Context, model *UserModel) error {
	collection := s.db.Database("micro-go").Collection("auth-user")
	docId, err := primitive.ObjectIDFromHex(model.Id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(ctx, bson.M{"_id": docId}, model)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlStore) FindByEmail(ctx context.Context, email string) (*UserModel, error) {
	var rs *UserModel
	collection := s.db.Database("micro-go").Collection("auth-user")
	if err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&rs); err != nil {
		return nil, err
	}
	return rs, nil
}
