package store

import (
	"context"
	"fmt"
	"log"

	"github.com/m-yosefpor/httpmon/internal/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const Collection = "users"

type MongoDB struct {
	db *mongo.Database
}

func NewMongoDBStore(db *mongo.Database) MongoDB {
	return MongoDB{
		db: db,
	}
}

func (m MongoDB) CreateUser(ctx context.Context, u model.User) error {
	log.Println(u)

	record := m.db.Collection(Collection).FindOne(ctx, bson.M{"id": u.ID})
	if record.Err() == nil {
		return ErrUserDuplicate
	}

	res, err := m.db.Collection(Collection).InsertOne(ctx, u)
	if err != nil {
		return fmt.Errorf("mongodb insert failed %w", err)
	}
	log.Println(res)
	return nil
}

func (m MongoDB) CreateEndpoint(ctx context.Context, id string, ep model.Endpoint) error {
	update := bson.M{
		"$addToSet": bson.M{
			"endpoints": ep,
		},
	}
	res, err := m.db.Collection(Collection).UpdateOne(ctx, bson.M{"id": id}, update)
	if err != nil {
		return fmt.Errorf("mongodb insert failed %w", err)
	}
	log.Println(res)
	return nil
}

func (m MongoDB) ListEndpoints(ctx context.Context, id string) ([]model.Endpoint, error) {
	var user model.User
	log.Println(id)
	record := m.db.Collection(Collection).FindOne(ctx, bson.M{"id": id})

	if err := record.Decode(&user); err != nil {
		return []model.Endpoint{}, fmt.Errorf("reading from mongodb failed %w", err)
	}

	return user.Endpoints, nil
}

func (m MongoDB) ListAlerts(ctx context.Context, id string) ([]model.Endpoint, error) {
	var user model.User

	record := m.db.Collection(Collection).FindOne(ctx, bson.M{"id": id})

	if err := record.Decode(&user); err != nil {
		return []model.Endpoint{}, fmt.Errorf("reading from mongodb failed %w", err)
	}

	return user.Endpoints, nil
}

func (m MongoDB) StatEndpoint(ctx context.Context, id, url string) (model.Endpoint, error) {
	var user model.User

	record := m.db.Collection(Collection).FindOne(ctx, bson.M{"id": id})

	if err := record.Decode(&user); err != nil {
		return model.Endpoint{}, fmt.Errorf("reading from mongodb failed %w", err)
	}

	for _, ep := range user.Endpoints {
		if ep.URL == url {
			return ep, nil
		}
	}

	return model.Endpoint{}, fmt.Errorf("endpoint not found")
}

// func (m MongoDB) Load(ctx context.Context) ([]model.User, error) {
// 	var users []model.User

// 	records, err := m.db.Collection(Collection).Find(ctx, bson.M{})
// 	if err != nil {
// 		return nil, fmt.Errorf("mongo find failed %w", err)
// 	}

// 	for records.Next(ctx) {
// 		var user model.User

// 		if err := records.Decode(&user); err != nil {
// 			return users, fmt.Errorf("mongo record decoding failed %w", err)
// 		}

// 		users = append(users, user)
// 	}

// 	return users, nil
// }
