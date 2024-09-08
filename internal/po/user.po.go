package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID      uuid.UUID `gorm:"column:uuid;type:uuid;index:idx_uuid;primaryKey;unique;not null"`
	UserName  string    `gorm:"column:user_name;type:varchar(50);not null"`
	IsActived bool      `gorm:"column:is_actived;type:boolean"`
	Roles     []Role    `gorm:"many2many:go_user_roles"`
}

func (r *User) TableName() string {
	return "go_db_users"
}
