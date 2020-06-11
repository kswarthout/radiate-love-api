package userapi

import "go.mongodb.org/mongo-driver/bson/primitive"

// Credentials is a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// Address represents a street address
type Address struct {
	Street  string `json:"street,omitempty" bson:"street,omitempty"`
	City    string `json:"city,omitempty" bson:"city,omitempty"`
	State   string `json:"state,omitempty" bson:"state,omitempty"`
	Country string `json:"country,omitempty" bson:"country,omitempty"`
	Zip     string `json:"zip,omitempty" bson:"zip,omitempty"`
}

// User represents an app user
type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Mobile    string             `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Age       int                `json:"age,omitempty" bson:"age,omitempty"`
	WhatsApp  string             `json:"whatsapp,omitempty" bson:"whatsapp,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	Address   Address            `json:"address,omitempty" bson:"address,omitempty"`
	Role      string             `json:"role,omitempty" bson:"role,omitempty"`
	StartDate uint32             `json:"start_date,omitempty" bson:"start_date,omitempty"`
}
