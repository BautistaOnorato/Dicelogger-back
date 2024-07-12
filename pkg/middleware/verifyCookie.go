package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/pkg/firebaseConnection"
)

func VerifySessionCookie() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Cookie("Session")
		if err != nil {
			ctx.AbortWithStatusJSON(401, err)
			return
		}
		
		authClient := firebaseConnection.CreateFirebaseClient()	
		
		

		_, err = authClient.VerifySessionCookieAndCheckRevoked(ctx, cookie)		
		if err != nil {
			ctx.AbortWithStatusJSON(498, err)
			return
		} else {
			ctx.Next()
		}
		
	}
}
