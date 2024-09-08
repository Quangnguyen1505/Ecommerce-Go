package po

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID       int64  `gorm:"column:id;type:int;primaryKey;autoIncrement;comment:'Primary key'"`
	RoleName string `gorm:"column:role_name;type:varchar(50)"`
	RoleNote string `gorm:"column:role_note;type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
