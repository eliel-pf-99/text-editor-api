package users

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type Handler interface {
	Login(c *gin.Context)
	Signup(c *gin.Context)
	Auth(c *gin.Context)
}

type handler struct {
	service Service
}

// Auth implements Handler.
func (h *handler) Auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpect signing method")
		}
		return []byte(os.Getenv("secret")), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		user, err := h.service.FindUserById(c, claims["sub"].(string))
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user.ID)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

// Signup implements Handler.
func (h *handler) Signup(c *gin.Context) {
	var body UserSignUp

	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	_, err := h.service.FindUserByEmail(c, body.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User already exists",
		})
		return
	}

	_, err = h.service.InsertUser(c, body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User create with success!",
	})
}

// Login implements Handler.
func (h *handler) Login(c *gin.Context) {
	var body UserLogin
	if c.BindJSON(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read request",
		})
		return
	}

	user, err := h.service.FindUserByEmail(c, body.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User don't find ",
		})
		return
	}

	if !CheckPassword(body.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password invalid",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failded to create token",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name":  user.Name,
		"token": tokenString,
	})

}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}
