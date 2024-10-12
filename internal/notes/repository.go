package notes

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	InsertNote(ctx context.Context, note Note) (Note, error)
	UpdateNote(ctx context.Context, note Note) (Note, error)
	DeleteNote(ctx context.Context, id string) error
	FindNoteById(ctx context.Context, id string) (Note, error)
	GetNotes(ctx context.Context, user_id string) ([]Note, error)
	DeleteNotes(ctx context.Context, user_id string) error
}

type repository struct {
	db              *mongo.Client
	database_name   string
	collection_name string
}

func NewRepository(db *mongo.Client, database_name, collection_name string) (Repository, error) {
	return &repository{db: db, database_name: database_name, collection_name: collection_name}, nil
}

func (r *repository) InsertNote(ctx context.Context, note Note) (Note, error) {
	_, err := r.db.Database(r.database_name).Collection(r.collection_name).InsertOne(ctx, note)
	if err != nil {
		return Note{}, err
	}
	return note, nil
}

func (r *repository) UpdateNote(ctx context.Context, note Note) (Note, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: note.ID}}

	_, err := coll.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: note}})
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (r *repository) DeleteNote(ctx context.Context, id string) error {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: id}}

	_, err := coll.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) FindNoteById(ctx context.Context, id string) (Note, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "id", Value: id}}

	var note Note
	err := coll.FindOne(ctx, filter).Decode(&note)
	if err != nil {
		return Note{}, err
	}

	return note, nil
}

func (r *repository) GetNotes(ctx context.Context, user_id string) ([]Note, error) {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "user_id", Value: user_id}}

	var notes []Note
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return []Note{}, err
	}

	if err = cursor.All(ctx, &notes); err != nil {
		return []Note{}, err
	}

	return notes, nil
}

func (r *repository) DeleteNotes(ctx context.Context, user_id string) error {
	coll := r.db.Database(r.database_name).Collection(r.collection_name)
	filter := bson.D{{Key: "user_id", Value: user_id}}
	_, err := coll.DeleteMany(ctx, filter)
	return err
}
