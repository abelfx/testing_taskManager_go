package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username string             `bson:"username" json:"username"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
	Role     string             `bson:"role" json:"role"` // e.g., "admin", "user"
}

type UserRepository interface {
	Create(user *User) (primitive.ObjectID, error)
	Authenticate(usernameOrEmail, password string) (*User, error)
	GetByID(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	PromoteUser(userID primitive.ObjectID, newRole string) error
	DeleteByID(id primitive.ObjectID) error
}
