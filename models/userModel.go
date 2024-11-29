package models

import(
  "time"
  "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	 ID primitive.ObjectID `json:"id"`
     FirstName string `json:"firstName"`
	 LastName string `json:"lastName"`
	 Email string `json:"email"`
	 Password string `json:"password"`
	 Role string `json:"role"`
	 Phone string `json:"phone"`
	 Token string
	 RefreshToken string
	 CreatedAt time.Time
	 UpdatedAt time.Time
	 UserID string `json:"userID"`
}
