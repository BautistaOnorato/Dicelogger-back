package middleware

// func VerifySessionCookie() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		cookie, err := ctx.Cookie("session")
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(401, err)
// 			return
// 		}

// 		authClient := firebaseConnection.CreateFirebaseClient()
// 		_, err = authClient.VerifySessionCookieAndCheckRevoked(ctx, cookie)
// 		if err != nil {
// 			ctx.AbortWithStatusJSON(498, err)
// 			return
// 		} else {
// 			ctx.Next()
// 		}
// 	}
// }
