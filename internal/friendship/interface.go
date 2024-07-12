package friendship

import "github.com/proyecto-dnd/backend/internal/domain"

type FriendshipRepository interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
	SearchFollowers(friendship domain.Mutuals) ([]domain.UserResponse, error)
	GetAllFriends(userId string) ([]domain.FriendUserData, error)
	GetBySimilarName(input string, userid string) ([]domain.FriendUserData, error)
}

type FriendshipService interface {
	Create(friendship domain.Friendship) (domain.Friendship, error)
	Delete(friendship domain.Friendship) error
	SearchFollowers(friendship domain.Mutuals) ([]domain.UserResponse, error)
	GetAllFriends(userId string) ([]domain.FriendUserData, error)
	GetBySimilarName(input string, userid string) ([]domain.FriendUserData, error)
}
