package middleware

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	auth "github.com/hafizdkn/toko-belanja/apps/domain/auth/jwt"
	"github.com/hafizdkn/toko-belanja/apps/domain/handler/views"
	"github.com/hafizdkn/toko-belanja/apps/domain/user"
)

func AuthenticationMiddleware(authService auth.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		unauthorizedResponse := "Unauthorized"
		cantParseTokenReponse := "Can't parse token"
		tokenIvalidResponse := "Invalid Token"

		headerToken := c.Request.Header.Get("Authorization")
		bearer := strings.HasPrefix(headerToken, "Bearer")
		if !bearer {
			err := errors.New(unauthorizedResponse)
			resp := views.UnauthorizedResponse(err, unauthorizedResponse)
			views.AbortJsonRespnse(c, resp)
			return
		}

		stringToken := strings.Split(headerToken, " ")
		if len(stringToken) != 2 {
			err := errors.New(cantParseTokenReponse)
			resp := views.UnauthorizedResponse(err, "")
			views.AbortJsonRespnse(c, resp)
			return
		}

		token, err := authService.ValidateToken(stringToken[1])
		if err != nil {
			resp := views.UnauthorizedResponse(err, "")
			views.AbortJsonRespnse(c, resp)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok && !token.Valid {
			err := errors.New(tokenIvalidResponse)
			resp := views.UnauthorizedResponse(err, "")
			views.AbortJsonRespnse(c, resp)
			return
		}

		if err != nil {
			resp := views.UnauthorizedResponse(err, "")
			views.AbortJsonRespnse(c, resp)
			return
		}

		c.Set("userData", claim)
		c.Next()
	}
}

func AuthorizationMiddleware(userService user.IService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userData := c.MustGet("userData").(jwt.MapClaims)
		userEmail := userData["email"].(string)

		user, err := userService.GetUserByEmail(userEmail)
		if err != nil {
			resp := views.InternalServerError(err)
			views.AbortJsonRespnse(c, resp)
			return
		}

		if user.Role != "admin" {
			err := errors.New("Unauthorized")
			msg := "You are not authorized to access this resource"
			resp := views.UnauthorizedResponse(err, msg)
			views.AbortJsonRespnse(c, resp)
			return
		}

		c.Next()
	}
}
