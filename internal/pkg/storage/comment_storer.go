package storage

import (
	"context"
	"time"
)

type CommentStorer interface {
	Insert (ctx context.Context, com Comment) error
	Find(ctx context.Context, uuid string) (Comment, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, com Comment) error
	List(ctx context.Context, page, limit int) ([]*Comment, error)
}


type Comment struct {
	UUID      string    `bson:"uuid"`
	Text      string    `bson:"text"`
	CreatedAt time.Time `bson:"created_at"`
}