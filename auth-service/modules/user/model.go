package user

type UserModel struct {
	Id        string `bson:"uid"`
	Email     string `bson:"email"`
	FirstName string `bson:"first_name,omitempty"`
	LastName  string `bson:"last_name,omitempty"`
	Password  string `bson:"-"`
	Salt      string `bson:"salt"`
	IsActive  int64  `bson:"is_active"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}
