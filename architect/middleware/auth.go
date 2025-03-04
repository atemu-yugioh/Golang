package middleware

import (
	"architect/common"
	"context"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (common.Requester, error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))

		if err != nil {
			common.WriteErrorResponse(ctx, err)
			ctx.Abort()
			return
		}

		requester, err := ac.IntrospectToken(ctx, token)

		if err != nil {
			common.WriteErrorResponse(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(common.KeyRequester, requester)

		ctx.Next()
	}
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("missing access token")
	}

	return parts[1], nil
}
