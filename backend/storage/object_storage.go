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

//go:generate mockgen -source=./object_storage.go -destination=./mocks/object_storage.go -package=mocks

type MongoClient interface {
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

type MongoCollection interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

type MongoDatabaseHandler struct {
	courseCollection MongoCollection
	userCollection   MongoCollection
}

func NewMustMongoDatabaseHandler() MongoDatabaseHandler {
	uri := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return MongoDatabaseHandler{
		courseCollection: client.Database(dbName).Collection("courses"),
		userCollection:   client.Database(dbName).Collection("users"),
	}
}

func (mongo MongoDatabaseHandler) SaveCourse(course models.Course) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := mongo.courseCollection.InsertOne(ctx, course)
	if err != nil {
		return err
	}
	return nil
}

func (mongo MongoDatabaseHandler) UserKnowledge(course string) (string, error) {
	return "no knowledge about anything", nil
}

func (mongo MongoDatabaseHandler) EnrollUserInCourse(userID, courseID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the user's enrolled courses array
	filter := bson.M{"user_id": userID}
	update := bson.M{
		"$addToSet": bson.M{"enrolled_courses": courseID}, // add courseID to enrolled_courses array
	}

	_, err := mongo.userCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (mongo MongoDatabaseHandler) GetCourse(courseID string) (*models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var course models.Course
	err := mongo.courseCollection.FindOne(ctx, bson.M{"course_id": courseID}).Decode(&course)
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (mongo MongoDatabaseHandler) GetCourses(userID string) ([]models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Step 1: Fetch the enrolled course IDs for the user
	filter := bson.M{"user_id": userID}
	projection := bson.M{"_id": 0, "enrolled_courses": 1}

	var userEnrollment struct {
		EnrolledCourses []string `bson:"enrolled_courses"`
	}

	err := mongo.userCollection.FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&userEnrollment)
	if err != nil {
		return nil, err // Return error if user is not found or query fails
	}

	// Step 2: If no courses are enrolled, return an empty list
	if len(userEnrollment.EnrolledCourses) == 0 {
		return []models.Course{}, nil
	}

	// Step 3: Fetch course details for the enrolled course IDs
	courseFilter := bson.M{"course_id": bson.M{"$in": userEnrollment.EnrolledCourses}} // Match courses in the array
	cursor, err := mongo.courseCollection.Find(ctx, courseFilter)
	if err != nil {
		return nil, err // Return error if query fails
	}
	defer cursor.Close(ctx)

	// Step 4: Decode the course details into a slice
	var result []models.Course
	for cursor.Next(ctx) {
		var course models.Course
		if err := cursor.Decode(&course); err != nil {
			return nil, err // Return error if decoding fails
		}
		result = append(result, course)
	}

	// Step 5: Return the list of enrolled courses
	return result, nil
}
