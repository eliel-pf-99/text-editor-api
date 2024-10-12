package users

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertUser(ctx context.Context, user User) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	DeleteUser(ctx context.Context, id string) error
	FindUserById(ctx context.Context, id string) (User, error)
	FindUserByEmail(ctx context.Context, email string) (User, error)
}

type repository struct {
	db              *mongo.Client
	database_name   string
	collection_name string
}

func NewRepository(db *mongo.Client, db_name, collection_name string) Repository {
	return &repository{db: db, database_name: db_name, collection_name: collection_name}
}

// DeleteUser implements Repository.
func (r *repository) DeleteUser(ctx context.Context, id string) error {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: id}}
	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

// FindUserById implements Repository.
func (r *repository) FindUserById(ctx context.Context, id string) (User, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: id}}

	var user User
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// InsertUser implements Repository.
func (r *repository) InsertUser(ctx context.Context, user User) (User, error) {
	_, err := r.db.Database(r.database_name).Collection(r.collection_name).InsertOne(ctx, user)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

// UpdateUser implements Repository.
func (r *repository) UpdateUser(ctx context.Context, user User) (User, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: user.ID}}

	_, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: user}})
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// FindUserByEmail implements Repository.
func (r *repository) FindUserByEmail(ctx context.Context, email string) (User, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "email", Value: email}}

	var user User
	err := coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return User{}, err
	}

	return user, nil
}
