package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"github.com/roobiuli/MongoCRUD/internal/pkg/storage"
)

var _ storage.CommentStorer = CommentStorage{}

type CommentStorage struct {
	Database *mongo.Database
	Timeout time.Duration
}


func (c CommentStorage) Insert(ctx context.Context, com storage.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	if _, err := c.Database.Collection("Comments").InsertOne(ctx, com); err != nil {
		log.Println(err)
		
		return err
	}
	return nil
}

func (c CommentStorage) Find(ctx context.Context, uuid string) (storage.Comment, error) {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	var com storage.Comment

	query := bson.M{"uuid": uuid}

	if  err := c.Database.Collection("Comments").FindOne(ctx, query).Decode(&com); err != nil {
		log.Println(err)
		return storage.Comment{}, err
	}
	return com, nil
}

func (c CommentStorage) Delete(ctx context.Context, uuid string)  error {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()
	query :=  bson.M{"uuid": uuid}
	 del, err := c.Database.Collection("Comments").DeleteOne(ctx, query)

	 if err != nil {
		 log.Println(err)
		 return err
	 }
	 if del.DeletedCount == 0 {
		 return errors.New("No Match Found")
	 }
	return nil
}

func (c CommentStorage) Update(ctx context.Context, com storage.Comment) error {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	qry := bson.M{"uuid": com.UUID}
	upd := bson.M{"$set": bson.M{"text": com.Text}}

	res, err := c.Database.Collection("Comments").UpdateOne(ctx, qry, upd)

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return errors.New("Comment not found, nothing was updated")
	}
	return nil

}

func (c CommentStorage) List(ctx context.Context, page, limit int) ([]*storage.Comment, error)  {
	ctx, cancel := context.WithTimeout(ctx, c.Timeout)
	defer cancel()

	var comments []*storage.Comment

	var skip int
	if page > 1 {
		skip = (page - 1) * limit
	}

	qry := bson.M{}
	opt := options.FindOptions{}
	opt.SetSkip(int64(skip))
	opt.SetLimit(int64(limit))
	opt.SetSort(bson.M{"created_at": -1})

	cur, err := c.Database.Collection("Comments").Find(ctx, qry, &opt)

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {

		comment := &storage.Comment{}

		err := cur.Decode(comment)

		if err != nil {
			log.Println(err)

			return nil, err
		}

		comments = append(comments, comment)
	}

	defer cur.Close(context.Background())

	if err := cur.Err(); err != nil {
		log.Println(err)

		return nil, err
	}

	return comments, nil

}