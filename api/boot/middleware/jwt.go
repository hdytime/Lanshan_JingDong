package middleware

import (
	"Lanshan_JingDong/api/boot/model"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

func NewToken(c *gin.Context, username string) string {
	claims := model.MyClaims{
		Username:       username,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString(model.MySecret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to sign token"})
		return ""
	}
	return result
}

func ParseToken(tokenstring string) (*model.MyClaims, error) {
	var claims model.MyClaims
	token, err := jwt.ParseWithClaims(tokenstring, &claims, func(token *jwt.Token) (interface{}, error) {
		return model.MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"error": "请求头中auth为空",
			})
			c.Abort()
			return
		}
		//按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"error": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		//parts[1]是获取到的tokenstring,使用解析jwt函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": "无效的token",
			})
			c.Abort()
			return
		}
		//将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)
		c.Next() //后续的处理函数可以用c.Get("username")来获取当前请求的用户信息
	}
}
