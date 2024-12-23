package serializer

import "github.com/mhrlife/tonference/internal/ent"

type User struct {
	ID          string `json:"id"`
	DisplayName string `json:"display_name"`
}

func NewUser(user *ent.User) User {
	return User{
		ID:          user.GameID,
		DisplayName: user.DisplayName,
	}
}

type UserWithToken struct {
	User
	Token string `json:"token"`
}

func NewUserWithJwt(user *ent.User, token string) UserWithToken {
	return UserWithToken{
		User:  NewUser(user),
		Token: token,
	}
}
