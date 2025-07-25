package repositories

import (
	"context"
	"errors"
	"time"

	infrastructure "restfulapi/Infrastructure"
	"restfulapi/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type MongoUserRepository struct {
	Collection *mongo.Collection
}

func NewMongoUserRepository(col *mongo.Collection) domain.UserRepository {
	return &MongoUserRepository{Collection: col}
}

func (r *MongoUserRepository) Create(user *domain.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check duplicate username/email
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}
	var existing domain.User
	err := r.Collection.FindOne(ctx, filter).Decode(&existing)
	if err == nil {
		return primitive.NilObjectID, errors.New("username or email already in use")
	}

	// Hash password
	hashedPassword, err := infrastructure.HashPassword(user.Password)
	if err != nil {
		return primitive.NilObjectID, err
	}
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	// Set default role: first user admin else user
	count, err := r.Collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return primitive.NilObjectID, err
	}
	if count == 0 {
		user.Role = "admin"
	} else if user.Role == "" {
		user.Role = "user"
	}

	_, err = r.Collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return user.ID, nil
}

func (r *MongoUserRepository) Authenticate(usernameOrEmail, password string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

func (r *MongoUserRepository) GetByID(id primitive.ObjectID) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MongoUserRepository) GetAll() ([]*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*domain.User
	for cursor.Next(ctx) {
		var user domain.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *MongoUserRepository) PromoteUser(userID primitive.ObjectID, newRole string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"role": newRole}}
	result, err := r.Collection.UpdateByID(ctx, userID, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *MongoUserRepository) DeleteByID(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
