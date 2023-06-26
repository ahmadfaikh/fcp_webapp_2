package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		// Request Cookie
		cookie, err := ctx.Request.Cookie("session_token")

		// Error Handling
		if err != nil {
			if ctx.GetHeader("Content-Type") != "application/json" {
				ctx.JSON(http.StatusSeeOther, model.ErrorResponse{Error: "Content-type undefined"})
				return
			}
			if err == http.ErrNoCookie {
				ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
				return
			}
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}

		// Kane new variable Model Claims
		claim := &model.Claims{}

		// JWT
		token, err := jwt.ParseWithClaims(cookie.Value, claim, func(tkn *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
			return
		}
		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: err.Error()})
			return
		}

		ctx.Writer.WriteHeader(http.StatusOK)
		ctx.Set("email", claim.Email)
		ctx.Next()
	})
}
