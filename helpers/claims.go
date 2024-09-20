package helpers

import (
	"context"
	"defterdar-go/models"
	"errors"
)

func GetClaims(ctx context.Context) (*models.Claims, error) {
	claims, ok := ctx.Value("user").(*models.Claims)
	if !ok {
		return nil, errors.New("claims not found")
	}
	return claims, nil
}
