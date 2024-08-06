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

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Make sure that the token method conforms to "SigningMethodHMAC"
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return []byte("secret"), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Optionally, you can set user information in the context
        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("userID", claims["id"])
            c.Set("userRole", claims["role"])
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        c.Next()
    }
}



func RoleMiddleware(roles ...string) gin.HandlerFunc {
    return func(c *gin.Context) {
        userRole, exists := c.Get("userRole")
        if !exists {
            c.JSON(http.StatusForbidden, gin.H{"error": "No role found in token"})
            c.Abort()
            return
        }

        roleValid := false
        for _, role := range roles {
            if userRole == role {
                roleValid = true
                break
            }
        }

        if !roleValid {
            c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
            c.Abort()
            return
        }

        c.Next()
    }
}
