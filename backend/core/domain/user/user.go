package user

type User struct {
	Id       string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`

	// TODO
}

type BasicUserInfo struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CreateUserRequest struct {
	Id       string `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
}
