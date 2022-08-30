package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	"ginkgo/internal/config"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/hash"
	jwtPkg "ginkgo/internal/pkg/jwt"
	"ginkgo/internal/pkg/response"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Registration struct {
	Username string `json:"userName"`
	Password string `json:"passWord"`
	Nickname string `json:"nickName"`
}

type UserToken struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

func (u *UserLogin) Login(config *config.Config,
	db *database.Client) response.Response {
	u.Password = hash.MD5V([]byte(u.Password))
	var modelU User
	modelU.Username = u.Username
	modelU.Password = u.Password
	err := db.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&modelU).Error
	if err != nil {
		return response.Error(response.ErrUsernameOrPassword)
	}
	j := &jwtPkg.JWT{SigningKey: []byte(config.Jwt.SigningKey)} // 唯一签名
	claims := jwtPkg.CustomClaims{
		ID:       modelU.ID,
		NickName: modelU.Nickname.String,
		Username: modelU.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                      // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(config.Jwt.Expired), // 过期时间 7天  配置文件
			Issuer:    "qmPlus",                                      // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		return response.Error(response.ErrJwtToken)
	}
	modelU.Token = token
	db.DB.Model(&User{}).Where("id = ?", modelU.ID).Update("token", token)
	return response.Success(modelU)
}

func (r *Registration) Create(config *config.Config,
	db *database.Client) response.Response {
	var u User
	u.Username = r.Username
	u.Password = r.Password
	u.Nickname = sql.NullString{Valid: true, String: r.Nickname}

	if !errors.Is(db.DB.Where("username = ?", u.Username).First(&u).Error, gorm.ErrRecordNotFound) {
		return response.Error(response.UserExist)
	}

	u.Password = hash.MD5V([]byte(r.Password))

	if err := db.DB.Create(&u).Error; err != nil {
		return response.Error(response.ErrSQL)
	}
	return response.Success("成功创建用户")
}

func (t *UserToken) GetUser() (err error, res response.Response) {
	// var u User
	// err = global.GDB.Where("id = ?", t.ID).First(&u).Error
	// if err != nil {
	// 	return err, response.Success("认证错误")
	// }
	return nil, response.Success("认证成功")
}
