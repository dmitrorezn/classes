package classes

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var mySigningKey = []byte("2secret3")

type obj bson.M

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	Email          string             `json:"email" bson:"email"`
	HashedPassword string             `json:"password" bson:"password"`
	Order          Order              `json:"order" bson:"order"`
	Status         string             `json:"status" bson:"status"`
	//Status 	for              Admin 	 -> admin
	//Status 	for      		 User	 -> user
	//Status    for Announcement Author  -> author
}

type MyClaim struct {
	jwt.StandardClaims
	Username   string
	Status     string
	Expiration int64
}

type Service struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

type Announcement struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	IDString      string             `json:"idstr" bson:"_idstr"`
	AuthorLogin   string             `json:"author_login" bson:"auth_login"`
	Activity      Activity           `json:"activity" bson:"activity"`
	UserRate      int                `json:"user_rate" bson:"user_rate"`
	Email         string             `json:"email" bson:"email"`
	PhoneNumber   string             `json:"phone_number" bson:"phone_number"`
	StartWeekDays []string           `json:"start_dates" bson:"start_dates"`
	Comments      []Comment          `json:"comments" bson:"comments"`
}

type Activity struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	AnnounsIDStr string             `json:"announsidstr" bson:"_announsidstr"`
	Name         string             `json:"name" bson:"name"`
	Type         string             `json:"type" bson:"type"`
	Price        float64            `json:"price" bson:"price"`
	StartDate    primitive.DateTime `json:"start_date" bson:"start_date"`
	Description  string             `json:"description" bson:"description"`
}

type Order struct {
	OrderList  []Activity `json:"order_list" bson:"order_list"`
	TotalPrice float64    `json:"total_price" bson:"total_price"`
}

type Comment struct {
	CommentText string             `json:"comment_text" bson:"comment_text"`
	UserLogin   string             `json:"user_login" bson:"user_login"`
	Date        primitive.DateTime `json:"date" bson:"date"`
}

func (u *User) GetToken() string {
	// Создаем новый токен
	token := jwt.New(jwt.SigningMethodHS256)
	// Устанавливаем набор параметров для токена
	token.Claims = MyClaim{
		Username:   u.Username,
		Status:     u.Status,
		Expiration: time.Now().Add(time.Minute * 5).Unix(),
	}
	// Подписываем токен нашим секретным ключем
	tokenString, _ := token.SignedString(mySigningKey)
	// Отдаем токен клиенту
	return tokenString
}


