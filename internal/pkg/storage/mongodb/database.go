package mongodb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)




type DatabaseConfig struct {
	Auth string
	Host string
	Port string
	User string
	Pass string
	Name string
}



func NewDatabase(ctx context.Context, cnf DatabaseConfig) (*mongo.Database, error) {
	md, err := mongo.Connect(ctx, options.Client().SetAuth(
		options.Credential{
			AuthMechanism:           cnf.Auth,
			AuthMechanismProperties: nil,
			AuthSource:              cnf.Name,
			Username:                cnf.User,
			Password:                cnf.Pass,
		},
	).ApplyURI(fmt.Sprintf("mongo://%s:%s", cnf.Host, cnf.Port)),)
	if err != nil {
		return nil, err
	}
	return md.Database(cnf.Name), nil
}
