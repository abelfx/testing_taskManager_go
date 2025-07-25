package infrastructure

import (
	"net/http"
	"restfulapi/usecases"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var jwtSecret = []byte("my_secret_key_but_will_be_replaced_by_env_later")

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			ctx.Abort()
			return
		}

		// Header format: "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			ctx.Abort()
			return
		}

		tokenStr := parts[1]

		// ✅ Parse token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			ctx.Abort()
			return
		}

		// ✅ Extract user ID from claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if userID, ok := claims["user_id"].(string); ok {
				ctx.Set("userID", userID)
			}
		}

		ctx.Next()
	}
}

func AdminOnly(userUsecase *usecases.UserUsecase) gin.HandlerFunc {
    return func(ctx *gin.Context) {
        userIDstr, exists := ctx.Get("userID")
        if !exists {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
            ctx.Abort()
            return
        }
        objectID, err := primitive.ObjectIDFromHex(userIDstr.(string))
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
            ctx.Abort()
            return
        }

        user, err := userUsecase.GetUserByID(objectID)
        if err != nil {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            ctx.Abort()
            return
        }
        if user.Role != "admin" {
            ctx.JSON(http.StatusForbidden, gin.H{"error": "Admin access required"})
            ctx.Abort()
            return
        }
        ctx.Next()
    }
}
