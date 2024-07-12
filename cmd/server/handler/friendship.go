package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/friendship"
	"github.com/proyecto-dnd/backend/internal/user"
)

type FriendshipHandler struct {
	service     friendship.FriendshipService
	userService user.ServiceUsers
}

func NewFriendshipHandler(service *friendship.FriendshipService, userService *user.ServiceUsers) *FriendshipHandler {
	return &FriendshipHandler{service: *service, userService: *userService}
}

// friendship godoc
// @Summary Create friendship
// @Tags friendship
// @Accept json
// @Produce json
// @Param body body domain.Friendship true "Friendship"
// @Success 201 {object} domain.Friendship
// @Failure 500 {object} error
// @Router /friendship [post]
func (h *FriendshipHandler) HandlerCreate() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		friend := ctx.Param("friend")
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id
		var tempFriendship domain.Friendship
		tempFriendship.User1Id = userId
		tempFriendship.User2Id = friend

		createdFriendship, err := h.service.Create(tempFriendship)
		if err != nil {
			ctx.JSON(400, "Failed to create friendship ;( ")
			return
		}

		ctx.JSON(201, "Created succesfully friendship with user1_id: "+createdFriendship.User1Id+" and user2_id: "+createdFriendship.User2Id)
	}
}

// friendship godoc
// @Summary Delete friendship
// @Tags friendship
// @Accept json
// @Produce json
// @Param body body domain.Friendship true "Friendship"
// @Success 200 {object} string
// @Failure 500 {object} error
// @Router /friendship [delete]
func (h *FriendshipHandler) HandlerDelete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		friend := ctx.Param("friend")
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id
		var tempFriendship domain.Friendship
		tempFriendship.User1Id = userId
		tempFriendship.User2Id = friend

		err = h.service.Delete(tempFriendship)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}
		ctx.JSON(200, "Deleted succesfully friendship with user1_id: "+tempFriendship.User1Id+" and user2_id: "+tempFriendship.User2Id)
	}
}

func (h *FriendshipHandler) HandlerSearchFollowers() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		name := ctx.Param("name")
		var tempFriendship domain.Mutuals
		tempFriendship.User2Name = name
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		tempFriendship.User1Id = jwtClaimsInfo.Id

		followers, err := h.service.SearchFollowers(tempFriendship)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, followers)
	}
}

func (h *FriendshipHandler) HandlerGetAllFriends() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id

		friends, err := h.service.GetAllFriends(userId)
		if err != nil {
			ctx.JSON(500, err.Error())
			return
		}

		ctx.JSON(200, friends)
	}
}

func (h *FriendshipHandler) HandlerGetBySimilarName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cookie, err := ctx.Request.Cookie("Session")
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		jwtClaimsInfo, err := h.userService.GetJwtInfo(cookie.Value)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}
		userId := jwtClaimsInfo.Id

		name := ctx.Param("name")
		users, err := h.service.GetBySimilarName(name, userId)
		if err != nil {
			ctx.JSON(400, err.Error())
			return
		}

		ctx.JSON(200, users)
	}
}
