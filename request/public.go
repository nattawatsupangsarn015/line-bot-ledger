package request

type Register struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Name     string `json:"name" bson:"name"`
	Phone    string `json:"phone" bson:"phone"`
}

type Login struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
