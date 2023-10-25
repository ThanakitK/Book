package repositories

import (
	"context"
	"golang-db/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type bookDBRepo struct {
	db         *mongo.Database
	collection string
}

func NewBookDBRepository(db *mongo.Database, collection string) BookRepository {
	return bookDBRepo{db, collection}
}

func (r bookDBRepo) GetAll() (result []models.RepoBookModel, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.db.Collection(r.collection).Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r bookDBRepo) Create(payload models.RepoBookModel) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = r.db.Collection(r.collection).InsertOne(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (r bookDBRepo) Update(book_id string, payload models.UpdateRepoBookModel) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": book_id}
	update := bson.M{"$set": payload}
	_, err = r.db.Collection(r.collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (r bookDBRepo) Delete(book_id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.M{"id": book_id}
	_, err = r.db.Collection(r.collection).DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
