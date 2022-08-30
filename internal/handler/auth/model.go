package auth

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string         `gorm:"type:varchar(128);not null"`
	Password  string         `gorm:"type:varchar(128);not null"`
	Nickname  sql.NullString `gorm:"type:varchar(128)"`
	Role      Role
	RoleID    sql.NullInt64
	IsSupper  uint8
	IsActive  uint8
	LastLogin sql.NullTime
	LastIP    string `gorm:"type:varchar(64)"`
	Token     string `gorm:"type:varchar(512)"`
}

type Role struct {
	gorm.Model
	Name        string         `gorm:"type:varchar(128);not null"`
	Desc        sql.NullString `gorm:"type:varchar(128)"`
	PagePerms   string         `gorm:"type:text"`
	DeployPerms string         `gorm:"type:text"`
	HostPerms   string         `gorm:"type:text"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
}
