package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `gore:"column:uuid; type:vachar(36); index:idx_uuid; primaryKey not_null; unique"`
	UserName  string    `gore:"column:user_name; type:varchar(50); not_null"`
	IsActived bool      `gore:"column:is_actived; type:boolean"`
	Roles     []Role    `gore:"many2many: go_user_roles"`
}

func (r *User) TableName() string {
	return "go_db_users"
}
