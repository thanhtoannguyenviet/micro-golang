package user

type UserModel struct {
	Id        string `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string `bson:"email" json:"email"`
	FirstName string `bson:"first_name,omitempty" json:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty" json:"last_name,omitempty"`
	Password  string `bson:"password" json:"password"`
	Salt      string `bson:"salt" json:"salt"`
	IsActive  int64  `bson:"is_active" json:"is_active"`
	CreatedAt int64  `bson:"created_at" json:"created_at"`
	UpdatedAt int64  `bson:"updated_at" json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" form:"email" bson:"email"`
	Password string `json:"password" form:"password" bson:"password"`
}
