package handler

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/user"
)

type UserHandler struct {
	service user.ServiceUsers
}

func NewUserHandler(service *user.ServiceUsers) *UserHandler {
	return &UserHandler{service: *service}
}

// user godoc
// @Summary Register user
// @Tags user
// @Accept json
// @Produce json
// @Param body body domain.User true "User"
// @Success 201 {object} domain.UserResponse
// @Failure 500 {object} error
// @Router /user/register [post]
func (h *UserHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("Register")
		var userTemp domain.User
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			log.Println(1, err)
			ctx.JSON(400, err)
			return
		}

		createdUser, err := h.service.Create(userTemp)
		if err != nil {
			// TEMP ERROR RESPONSE
			log.Println(2, err)
			ctx.JSON(400, err)
			return
		}

		// TEMP SUCCESS RESPONSE
		ctx.JSON(201, createdUser)
	}
}

// user godoc
// @Summary Get all users
// @Tags user
// @Produce json
// @Success 200 {array} domain.UserResponse
// @Failure 500 {object} error
// @Router /user [get]
func (h *UserHandler) HandlerGetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userList, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		//transfer users data to sql in bulk
		// _, err = h.service.TransferDataToSql(userList)
		// if err != nil {
		// 	// TEMP ERROR RESPONSE
		// 	fmt.Println(err)
		// 	ctx.JSON(500, err)
		// 	return
		// }

		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, userList)
	}
}

// user godoc
// @Summary Get user by id
// @Tags user
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} domain.UserResponse
// @Failure 500 {object} error
// @Router /user/{id} [get]
func (h *UserHandler) HandlerGetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		userTemp, err := h.service.GetById(id)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		ctx.JSON(200, userTemp)
	}
}

// user godoc
// @Summary Update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body domain.User true "User"
// @Success 200 {object} domain.User
// @Failure 500 {object} error
// @Router /user/{id} [put]
func (h *UserHandler) HandlerUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")

		var userTemp domain.UserUpdate
		if err := ctx.BindJSON(&userTemp); err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err.Error())
			return
		}
		updatedUser, err := h.service.Update(userTemp, id)
		if err != nil {
			// TEMP ERROR RESPONSE
			ctx.JSON(500, err.Error())
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, updatedUser)
	}
}

// user godoc
// @Summary Partially update user
// @Tags user
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param body body domain.User true "User"
// @Success 200 {object} domain.User
// @Failure 500 {object} error
// @Router /user/{id} [patch]
func (h *UserHandler) HandlerPatch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		jwtClaimsInfo, err := h.service.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id

		var userTemp domain.UserUpdate
		if err := ctx.BindJSON(&userTemp); err != nil {

			ctx.JSON(500, err.Error())
			return
		}
		patchedUser, err := h.service.Patch(userTemp, userId)
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, patchedUser)
	}
}

// user godoc
// @Summary Delete user
// @Tags user
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /user/{id} [delete]
func (h *UserHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := h.service.Delete(id)
		if err != nil {

			ctx.JSON(500, err.Error())
			return
		}

		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, "Deleted user with id: "+id)
	}
}

// user godoc
// @Summary Login user
// @Tags user
// @Accept json
// @Produce json
// @Param body body domain.UserLoginInfo true "UserLoginInfo"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /user/login [post]
func (h *UserHandler) HandlerLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var tempUserInfo domain.UserLoginInfo

		err := ctx.BindJSON(&tempUserInfo)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		cookie, err := h.service.Login(tempUserInfo)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		host := ctx.Request.Host
		domainParts := strings.Split(host, ":")
		domain := domainParts[0]
		ctx.SetCookie("Session", cookie, 36000, "/", domain, false, true)

		ctx.JSON(200, "Setted Cookie")
	}
}

func (h *UserHandler) HandlerGetJwtInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			log.Println(1, err)
			ctx.JSON(400, err.Error())
			return
		}
		// log.Println(cookie.Value)
		jwtClaimsInfo, err := h.service.GetJwtInfo(cookie.Value)
		if err != nil {
			log.Println(2, err)
			ctx.JSON(400, err.Error())
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, jwtClaimsInfo)
	}
}

func (h *UserHandler) HandlerSubPremium() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("SUBSCRIBING TO PREMIUM")

		monthsParam := ctx.Param("months")

		months, err := strconv.Atoi(monthsParam)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		// arreglar para sumar a la fecha guardada en la base checkeando que la fecha de hoy sea posterior al vencimiento de la suscripcion
		log.Print(time.Now().AddDate(0, months, 0).String())
		_, err = h.service.SubscribeToPremium(cookie.Value, time.Now().AddDate(0, months, 0).String())
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		// ctx.SetCookie("Session", newToken, 360000, "/", "localhost", false, true)
		ctx.JSON(200, "Subscribed to Premium")
	}
}

func (h *UserHandler) HandlerCheckSubscriptionExpDate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			log.Println(1, err)
			ctx.JSON(500, err.Error())
			return
		}
		jwtClaimsInfo, err := h.service.GetJwtInfo(cookie.Value)
		if err != nil {
			log.Println(2, err)
			ctx.JSON(400, err.Error())
			return
		}
		
		err = h.service.CheckSubExpiration(jwtClaimsInfo.Id)
		if err != nil {
			ctx.JSON(403, err.Error())
			return
		}
		// TEMP SUCCESS RESPONSE
		ctx.JSON(200, "Still subed to Premium")
	}
}

func (h *UserHandler) HandlerSendEmailVerification() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		jwtClaimsInfo, err := h.service.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		err = h.service.SendVerificationEmail(jwtClaimsInfo.Email)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, "Email sent")
	}
}
