package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/model/entity"
)

type UserSignInInput struct {
	LoginName string
	Password  string
}

type UserInfo struct {
	gmeta.Meta `orm:"table:users"`
	*entity.Users
	AccountValidTime []int64 `json:"accountValidTime,omitempty"`
}
