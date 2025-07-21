package data

import (
	"context"
	"errors"
	"restfulapi/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// CreateUser inserts a new user into the database after hashing the password
func CreateUser(user *models.User) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Check if username or email already exists
	filter := bson.M{
		"$or": []bson.M{
			{"username": user.Username},
			{"email": user.Email},
		},
	}
	var existing models.User
	err := UserCollection.FindOne(ctx, filter).Decode(&existing)
	if err == nil {
		return primitive.NilObjectID, errors.New("username or email already in use")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}

	// Set default values
	user.ID = primitive.NewObjectID()
	user.Password = string(hashedPassword)

	count, err := UserCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return primitive.NilObjectID, err
	}
	if count == 0 {
		user.Role = "admin"
	} else if user.Role == "" {
		user.Role = "user"
	}

	_, err = UserCollection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return user.ID, nil
}

// AuthenticateUser validates user credentials
func AuthenticateUser(usernameOrEmail, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	filter := bson.M{
		"$or": []bson.M{
			{"username": usernameOrEmail},
			{"email": usernameOrEmail},
		},
	}
	err := UserCollection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// GetUserByID fetches a user by ID
func GetUserByID(id primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetAllUsers fetches all users from the database
func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

// promote a user
func PromoteUser(userID primitive.ObjectID, newRole string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"role": newRole}}
	result, err := UserCollection.UpdateByID(ctx, userID, update)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

// DeleteUser deletes a user by their ID
func DeleteUser(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, err := UserCollection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
