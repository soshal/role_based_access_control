package middleware

import (
    "daily-api/models"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "strings"
	"time"
	"fmt"
)

var jwtSecret = []byte("your_secret_key")

// GenerateToken generates a JWT token
func GenerateToken(user models.User) (string, error) {
    
	//jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{...}): Creates a new JWT token with the specified claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": user.Username,
        "role":     user.Role,
        "exp":      jwt.TimeFunc().Add(time.Hour * 72).Unix(),
    })
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    return tokenString, nil
}

// AuthMiddleware authenticates the JWT token
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        c.Set("username", claims["username"])
        c.Set("role", claims["role"])
        c.Next()
    }
}

// RoleMiddleware authorizes based on user role
func RoleMiddleware(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        role, exists := c.Get("role")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
            c.Abort()
            return
        }

        userRole := role.(string)
        for _, r := range roles {
            if userRole == r {
                c.Next()
                return
            }
        }

        c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
        c.Abort()
    }
}
