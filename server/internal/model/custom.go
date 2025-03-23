package model

import (
	"github.com/gogf/gf/v2/util/gmeta"
	"server/internal/model/entity"
)

type WithUser struct {
	gmeta.Meta `orm:"table:users"`
	*entity.Users
	//Authenticate *entity.Authenticates `orm:"with:id=authenticate_id" json:"authenticate"`
}
