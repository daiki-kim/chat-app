package auth

import (
	"context"
	"time"

	"github.com/daiki-kim/chat-app/configs"
	"github.com/daiki-kim/chat-app/pkg/logger"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
)

const (
	Subject                = "AccessToken"
	Issuer                 = "github.com/daiki-kim/chat-app"
	Audience               = "github.com/daiki-kim/chat-app"
	TokenExpiration        = time.Minute * time.Duration(10)
	RefreshTokenExpiration = time.Hour * time.Duration(24)
)

var JwtSecret = []byte(configs.Config.JwtSecret)

type CustomClaim struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// Deifne unique type to avoid collosions in context
type UserIDKeyForContext struct{}

func NewClaim(userID int) *CustomClaim {
	return &CustomClaim{
		UserID: userID,
	}
}

func (c *CustomClaim) GenerateToken() (string, error) {
	claims := jwt.RegisteredClaims{
		ID:        uuid.New().String(),
		Issuer:    Issuer,
		Subject:   Subject,
		Audience:  []string{Audience},
		IssuedAt:  jwt.NewNumericDate(time.Now().Local()),
		ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(TokenExpiration)),
	}
	if err := copier.CopyWithOption(c, claims, copier.Option{IgnoreEmpty: true, DeepCopy: true}); err != nil {
		logger.Error("failed to copy claim", zap.Error(err))
		return "", err
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	token, err := generatedToken.SignedString(JwtSecret)
	if err != nil {
		logger.Error("failed to generate token", zap.Error(err))
		return "", err
	}
	return token, nil
}

func ParseToken(tokenStr string) (*CustomClaim, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&CustomClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(JwtSecret), nil
		})
	if err != nil {
		logger.Error("failed to parse token", zap.Error(err))
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaim)
	if !ok || !token.Valid {
		logger.Error("invalid token")
		return nil, err
	}

	return claims, nil
}

func GetUserIDFromContext(ctx context.Context) int {
	return ctx.Value(UserIDKeyForContext{}).(int)
}

func SetUserIDToContext(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, UserIDKeyForContext{}, userID)
}
