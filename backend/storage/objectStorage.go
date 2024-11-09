package storage

import (
	"backend/models"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var courseCollection *mongo.Collection
var userCollection *mongo.Collection

func InitDatabase() error {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	// Select collections
	courseCollection = client.Database(dbName).Collection("courses")
	userCollection = client.Database(dbName).Collection("users")
	return nil
}

func SaveCourse(course models.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := courseCollection.InsertOne(ctx, course)
	if err != nil {
		return err
	}
	return nil
}

func EnrollUserInCourse(userID, courseID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the user's enrolled courses array
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$addToSet": bson.M{"enrolled_courses": courseID}, // add courseID to enrolled_courses array
	}

	_, err := userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func GetCourse(courseID string) (*models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err := courseCollection.FindOne(ctx, bson.M{"course_id": courseID}).Decode(&course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}
