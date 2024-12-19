package controllers

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"os"
	"te-emb-api/initalizers"
	"te-emb-api/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Singup(c *gin.Context) {
	// Get the email/pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	// Create the user

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initalizers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	// Response

	c.JSON(http.StatusOK, gin.H{
		"message": "success add user",
	})
}

func Login(c *gin.Context) {
	// Get the email and pass off req body
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	// Look up requested user
	var user models.User
	initalizers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	// Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	// Generate a jwt token
	sessionId := GenerateSessionId()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":       user.ID,
		"exp":       time.Now().Add(time.Hour * 24 * 30).Unix(),
		"username":  user.Email,
		"sessionId": sessionId,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERT")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fail to locate token",
		})
		return
	}
	// store to Redis
	initalizers.SetSessionIdToRedis(user.Email, []byte(sessionId))
	datastr, _ := json.Marshal(user)
	initalizers.SetSessionIdToRedis(sessionId, datastr)

	// send it back
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Autherization", tokenString, 3600*24*1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"messsage": "login",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	username := user.(models.User).Email

	// user.(models.User).
	c.JSON(http.StatusOK, gin.H{
		"username": username,
	})
}

func GenerateSessionId() string {
	u4 := uuid.New()
	// fmt.Println(u4.String())
	// fmt.Println(base64.URLEncoding.EncodeToString([]byte(u4.String())))
	return base64.URLEncoding.EncodeToString([]byte(u4.String()))
}
