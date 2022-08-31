package user

type UserModel struct {
	Id        string `json:"uid"`
	Email     string `json:"email"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"`
	Salt      string `json:"salt"`
	IsActive  int64  `json:"is_active"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}
