package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"ginkgo/internal/config"
	"ginkgo/internal/handler/auth"
	"ginkgo/internal/pkg/database"
	"ginkgo/internal/pkg/jwt"
	"ginkgo/internal/pkg/response"
)

func jwtAuth(config *config.Config, db *database.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, response.Error(response.ErrAuth))
			c.Abort()
			return
		}
		j := jwt.NewJWT(config.Jwt.SigningKey)
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == jwt.TokenExpired {
				c.JSON(http.StatusOK, response.Error(response.ErrAuthExpired))
				c.Abort()
				return
			}
			c.JSON(http.StatusOK, response.ErrorUnknown(response.ErrAuthUnknown, err.Error()))
			c.Abort()
			return
		}
		service := auth.UserToken{}
		service.ID = int(claims.ID)
		res, err := service.GetUser(db)
		if err != nil {
			c.JSON(http.StatusOK, res)
			c.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + int64(config.Jwt.Expired)
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			// if global.GVA_CONFIG.System.UseMultipoint {
			// 	err, RedisJwtToken := service.GetRedisJWT(newClaims.Username)
			// 	if err != nil {
			// 		global.GVA_LOG.Error("get redis jwt failed", zap.Any("err", err))
			// 	} else { // 当之前的取成功时才进行拉黑操作
			// 		_ = service.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
			// 	}
			// 	// 无论如何都要记录当前的活跃状态
			// 	_ = service.SetRedisJWT(newToken, newClaims.Username)
			// })
		}
		c.Set("claims", claims)
		c.Next()
	}
}
