package mongodb

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/model"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/repository"
	"github.com/huytran2000-hcmus/bTaskee/booking/internal/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	taskCollection = "tasks"
)

var _ repository.Repository = (*Repository)(nil)

type Repository struct {
	db *mongo.Database
}

func NewRepository(databaseName string, uri string) (*Repository, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opt := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opt)
	if err != nil {
		return nil, fmt.Errorf("connect to mongodb: %w", err)
	}

	var result bson.M
	db := client.Database(databaseName)
	err = db.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("ping mongodb: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}

func (r *Repository) InsertTask(ctx context.Context, taskModel *model.Task) error {
	coll := r.db.Collection(taskCollection)

	task := toMongoDBTask(taskModel)

	_, err := coll.InsertOne(ctx, task)
	if err != nil {
		return err
	}

	return err
}

func (r *Repository) GetTaskByID(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	coll := r.db.Collection(taskCollection)

	task := &Task{}
	filter := bson.D{{"_id", id}}
	err := coll.FindOne(ctx, filter).Decode(task)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, service.ErrNotFound
		}

		return nil, err
	}

	return toModelTask(task), nil
}
