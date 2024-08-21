package middleware

import (
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/yusufguntav/hospital-management/pkg/entities"
	"github.com/yusufguntav/hospital-management/pkg/state"
)

func CheckAuth(authRole ...entities.AuthRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(400, gin.H{"error": "Token is required"})
			c.AbortWithStatus(400)
			return
		}

		authToken := strings.Split(authHeader, " ")
		if len(authToken) != 2 || authToken[0] != "Bearer" {
			c.JSON(400, gin.H{"error": "Invalid/Malformed auth token"})
			c.AbortWithStatus(400)
			return
		}

		myJwt := authToken[1]

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(myJwt, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			c.AbortWithStatus(400)
			return
		}

		if !token.Valid {
			c.JSON(400, gin.H{"error": "Token is not valid"})
			c.AbortWithStatus(400)
			return
		}

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.JSON(400, gin.H{"error": "token expired"})
			c.AbortWithStatus(400)
			return
		}

		role := entities.AuthRole(claims["role"].(float64))
		permission := false
		if len(authRole) == 0 {
			permission = true
		} else {
			for _, r := range authRole {
				if r == role {
					permission = true
					break
				}
			}
		}

		if !permission {
			c.JSON(400, gin.H{"error": "Permission denied"})
			c.AbortWithStatus(400)
			return
		}

		c.Set(state.CurrentUserId, claims["id"])
		c.Set(state.CurrentUserROLE, role)
		c.Next()
	}
}
