package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var timeoutDuration time.Duration = 20 * time.Second

func GetDB(host, user, pass, dbName string) (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.MergeClientOptions(
		&options.ClientOptions{
			Hosts:          []string{host},
			ConnectTimeout: &timeoutDuration,
			Auth: &options.Credential{
				Username: user,
				Password: pass,
			},
		}))
	if err != nil {
		return nil, err
	}
	return client.Database(dbName), nil
}
