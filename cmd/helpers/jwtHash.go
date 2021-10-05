package helpers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
	"time"
)
func init() {
	os.Setenv("SECRET_KEY", "TheJWTSecretToken")
}
func GenerateToken(id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expiredAt" : time.Now().Add(time.Hour * 24).Unix(),
		"userId" : id,
		"Authorized": true,
		"issuedAt": time.Now().Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET_KEY")))
}

func VerifyTokenMiddleware() gin.HandlerFunc  {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No Token Given"})
			return
		}
		token, err := jwt.Parse(header, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Invalid Token"})
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		expiryTime := claims["expiredAt"].(float64)
		if int64(expiryTime) < time.Now().Unix() {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message" : "Invalid Token"})
			return
		}
		userId := claims["userId"].(string)
		c.Set("userId", userId)
		c.Next()
	}

}