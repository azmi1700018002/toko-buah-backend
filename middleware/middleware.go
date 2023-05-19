package middleware

import (
	"time"
	"toko-buah/model/m_user"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var identityKey = "id"

func AuthMiddleware() (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*m_user.User); ok {
				return jwt.MapClaims{
					identityKey:      v.Username,
					"UserID":         v.UserID.String(), // tambahkan IDUser ke payload
					"Username":       v.Username,
					"ProfilePicture": v.ProfilePicture,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &m_user.User{
				Email:  claims[identityKey].(string),
				UserID: uuid.MustParse(claims["UserID"].(string)), // konversi string ke uuid.UUID
			}
		},
	})

	if err != nil {
		return nil, err
	}

	if errInit := authMiddleware.MiddlewareInit(); errInit != nil {
		return nil, errInit
	}

	return authMiddleware, nil
}
