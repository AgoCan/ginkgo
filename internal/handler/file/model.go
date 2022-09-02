package file

import "github.com/jinzhu/gorm"

type File struct {
	gorm.Model
	Name    string `gorm:"type:varchar(128);not null"`
	Path    string `gorm:"type:varchar(512)"`
	Size    int64  `gorm:"comment:'文件大小'"`
	Md5     string
	Type    string `gorm:"comment:'文件类型'"`
	BaseUrl string `gorm:"comment:'文件url，请求使用'"`
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&File{})
}
