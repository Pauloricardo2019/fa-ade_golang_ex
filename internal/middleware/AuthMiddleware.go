package middleware

import (
	"facade/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SecurityFacade interface {
	ValidateToken(*dto.ValidateTokenRequest) (*dto.ValidateTokenResponse, error)
}

type AuthMiddleware struct {
	securityFacade SecurityFacade
}

func NewAuthMiddleware(securityFacade SecurityFacade) *AuthMiddleware {
	return &AuthMiddleware{
		securityFacade: securityFacade,
	}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const bearerSchema = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := header[len(bearerSchema):]

		if header != (bearerSchema + token) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenDTO := &dto.ValidateTokenRequest{Value: token}

		response, err := a.securityFacade.ValidateToken(tokenDTO)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError,
				dto.Error{
					Message: err.Error(),
				},
			)
			return
		}

		if response == nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("user_id", response.UserID)
	}
}
