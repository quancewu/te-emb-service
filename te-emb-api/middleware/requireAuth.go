package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"te-emb-api/initalizers"
	"te-emb-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off req
	tokenString, err := c.Cookie("Autherization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Deecode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECERT")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check the exp
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		redisSessionId, err := initalizers.GetSessionIdFromRedis(claims["username"].(string))
		if err != nil {
			// fmt.Println("session timeout")
			// c.JSON(http.StatusUnauthorized, "session timeout")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		if string(redisSessionId) != claims["sessionId"].(string) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		userInfo, err := initalizers.GetSessionIdFromRedis(claims["sessionId"].(string))
		var user models.User
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			json.Unmarshal([]byte(userInfo), &user)
		}
		// Find the user with token sub
		// var user models.User
		// initalizers.DB.First(&user, claims["sub"])

		// if user.ID == 0 {
		// 	c.AbortWithStatus(http.StatusUnauthorized)
		// }
		// Attach to req
		c.Set("user", user)
		// Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
