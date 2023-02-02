package user

import "sixels.io/manekani/ent/schema"

type User struct {
	Id       string `json:"-"`
	Username string `json:"username"`
	Email    string `json:"email"`

	Level int32 `json:"level"`

	PendingActions []schema.PendingAction `json:"-"`
}

type UserBasic struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Level    int32  `json:"level"`
}

type CreateUserRequest struct {
	Id       string `json:"id" form:"id"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
}
