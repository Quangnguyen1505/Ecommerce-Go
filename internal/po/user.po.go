package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gore:"column:id; type:int; primaryKey; autoIncrement; comment:'Primary key'"`
	RoleName string `gore:"column:role_name"`
	RoleNote string `gore:"column:user_note; type:text;"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
