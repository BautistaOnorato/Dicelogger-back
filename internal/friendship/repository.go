package friendship

import (
	"database/sql"
	"errors"
	"log"
	"strings"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/user"
)

type repositoryFriendship struct {
	db             *sql.DB
	userRepository user.RepositoryUsers
	firebaseApp    *firebase.App
	authClient     *auth.Client
}

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrDuplicateName    = errors.New("name already exists")
	ctx                 = &gin.Context{}
)

func NewFriendshipRepository(db *sql.DB, userRepository user.RepositoryUsers, app *firebase.App) FriendshipRepository {
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFriendship{db: db, userRepository: userRepository, authClient: authClient}
}

func (r *repositoryFriendship) SearchFollowers(mutuals domain.Mutuals) ([]domain.UserResponse, error) {

	usersListChan := make(chan []domain.UserResponse)
	user1FriendsChan := make(chan []domain.FriendUserData)
	tempFriendListChan := make(chan []domain.UserResponse)

	go func() {
		user1Friends, err := r.GetAllFriends(mutuals.User1Id)
		if err != nil {
			user1FriendsChan <- []domain.FriendUserData{}
			return
		}
		user1FriendsChan <- user1Friends
		close(user1FriendsChan)
	}()

	go func() {
		usersList, err := r.userRepository.GetAll()
		if err != nil {
			usersListChan <- []domain.UserResponse{}
			return
		}
		usersListChan <- usersList
		close(usersListChan)
	}()
	go func() {
		var usersListByName []domain.UserResponse
		for _, user := range <-usersListChan {
			if strings.HasPrefix(strings.ToLower(user.Username), strings.ToLower(mutuals.User2Name)) {
				usersListByName = append(usersListByName, user)
			}
		}
		filteredList := []domain.UserResponse{}
		for _, user := range usersListByName {
			isFriend := false
			for _, friend := range <-user1FriendsChan {
				if user.Id == friend.UserId {
					isFriend = true
					break
				}
			}
			if !isFriend && user.Id != mutuals.User1Id {
				filteredList = append(filteredList, user)
			}
		}

		tempFriendListChan <- filteredList
		close(tempFriendListChan)

	}()

	tempFriendList := <-tempFriendListChan

	return tempFriendList, nil
}

func (r *repositoryFriendship) Create(friendship domain.Friendship) (domain.Friendship, error) {

	statement, err := r.db.Prepare(QueryCreate)
	if err != nil {
		return domain.Friendship{}, ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		friendship.User1Id,
		friendship.User2Id,
	)
	if err != nil {
		return domain.Friendship{}, err
	}

	return friendship, nil
}

func (r *repositoryFriendship) Delete(friendship domain.Friendship) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(
		friendship.User1Id,
		friendship.User2Id,
	)
	if err != nil {
		return err
	}

	return nil
}

// func (r *repositoryFriendship) IsFriends(userId1 string, userId2 string) (bool, error) {
// 	statement, err := r.db.Prepare(QueryCheckFriendship)
// 	if err != nil {
// 		return false, ErrPrepareStatement
// 	}
// 	defer statement.Close()

// 	var count int
// 	err = statement.QueryRow(userId1, userId2, userId2, userId1).Scan(&count)
// 	if err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }

func (r *repositoryFriendship) GetAllFriends(userId string) ([]domain.FriendUserData, error) {
	statement, err := r.db.Prepare(QueryGetFriends)
	if err != nil {
		return nil, ErrPrepareStatement
	}
	defer statement.Close()

	rows, err := statement.Query(userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friendships []domain.Friendship
	for rows.Next() {
		var friend domain.Friendship
		err = rows.Scan(&friend.User1Id, &friend.User2Id)
		if err != nil {
			return nil, err
		}
		if friend.User1Id == userId {
			friendships = append(friendships, friend)
		} else {
			friend.User1Id, friend.User2Id = friend.User2Id, friend.User1Id
			friendships = append(friendships, friend)
		}
	}
	var friendshipList []domain.UserResponse
	fullUserList, err := r.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, friendship := range friendships {
		for _, user := range fullUserList {
			if friendship.User2Id == user.Id {
				friendshipList = append(friendshipList, user)
			}
		}
	}

	friendshipsListrows, err := r.db.Query(QueryGetAllFriendships)
	if err != nil {
		return []domain.FriendUserData{}, err
	}
	defer friendshipsListrows.Close()

	var friendshipsList []domain.Friendship
	for friendshipsListrows.Next() {
		var friendship domain.Friendship
		err = friendshipsListrows.Scan(&friendship.User1Id, &friendship.User2Id)
		if err != nil {
			return []domain.FriendUserData{}, err
		}
		friendshipsList = append(friendshipsList, friendship)
	}
	var friendStatus domain.FriendUserData
	var friends []domain.FriendUserData

	for _, friend := range friendshipList {
		if r.isReciprocated(friendshipsList, userId, friend.Id) == "reciprocated" {
			friendStatus.UserId = friend.Id
			friendStatus.Username = friend.Username
			friendStatus.DisplayName = friend.DisplayName
			friendStatus.Email = friend.Email
			friendStatus.Image = friend.Image
			friendStatus.Following = true
			friendStatus.FollowsYou = true
			friends = append(friends, friendStatus)
			continue
		} else if r.isReciprocated(friendshipsList, userId, friend.Id) == "follows you" {
			friendStatus.UserId = friend.Id
			friendStatus.Username = friend.Username
			friendStatus.DisplayName = friend.DisplayName
			friendStatus.Email = friend.Email
			friendStatus.Image = friend.Image
			friendStatus.Following = true
			friendStatus.FollowsYou = false
			friends = append(friends, friendStatus)
			continue
		} else if r.isReciprocated(friendshipsList, userId, friend.Id) == "you follow" {
			friendStatus.UserId = friend.Id
			friendStatus.Username = friend.Username
			friendStatus.DisplayName = friend.DisplayName
			friendStatus.Email = friend.Email
			friendStatus.Image = friend.Image
			friendStatus.Following = false
			friendStatus.FollowsYou = true
			friends = append(friends, friendStatus)
			continue
		} else if r.isReciprocated(friendshipsList, userId, friend.Id) == "unknowns" {
			friendStatus.UserId = friend.Id
			friendStatus.Username = friend.Username
			friendStatus.DisplayName = friend.DisplayName
			friendStatus.Email = friend.Email
			friendStatus.Image = friend.Image
			friendStatus.Following = false
			friendStatus.FollowsYou = false
			friends = append(friends, friendStatus)
			continue
		}
	}

	return friends, nil
}

func (r *repositoryFriendship) GetBySimilarName(input string, userId string) ([]domain.FriendUserData, error) {
	usersList, err := r.userRepository.GetAll()
	if err != nil {
		return []domain.FriendUserData{}, err
	}

	// friendList, err := r.GetAllFriends(userId)
	// if err != nil {
	// 	return []domain.FriendUserData{}, err
	// }

	friendshipList, err := r.GetListFriendships()
	if err != nil {
		return []domain.FriendUserData{}, err
	}

	//filter the list of all users by people that is not in the friendlist
	// listByName := []domain.UserResponse{}
	foundUsersList := []domain.FriendUserData{}
	var foundUser domain.FriendUserData

	for _, user := range usersList {
		if strings.HasPrefix(strings.ToLower(user.Username), strings.ToLower(input)) && user.Id != userId {
			if r.isReciprocated(friendshipList, userId, user.Id) == "reciprocated" {
				foundUser.UserId = user.Id
				foundUser.Username = user.Username
				foundUser.DisplayName = user.DisplayName
				foundUser.Email = user.Email
				foundUser.Image = user.Image
				foundUser.Following = true
				foundUser.FollowsYou = true
				foundUsersList = append(foundUsersList, foundUser)
				continue
			} else if r.isReciprocated(friendshipList, user.Id, userId) == "follows you" {
				foundUser.UserId = user.Id
				foundUser.Username = user.Username
				foundUser.DisplayName = user.DisplayName
				foundUser.Email = user.Email
				foundUser.Image = user.Image
				foundUser.Following = false
				foundUser.FollowsYou = true
				foundUsersList = append(foundUsersList, foundUser)
				continue
			} else if r.isReciprocated(friendshipList, user.Id, userId) == "you follow" {
				foundUser.UserId = user.Id
				foundUser.Username = user.Username
				foundUser.DisplayName = user.DisplayName
				foundUser.Email = user.Email
				foundUser.Image = user.Image
				foundUser.Following = true
				foundUser.FollowsYou = false
				foundUsersList = append(foundUsersList, foundUser)
				continue
			} else if r.isReciprocated(friendshipList, userId, user.Id) == "unknowns" {
				foundUser.UserId = user.Id
				foundUser.Username = user.Username
				foundUser.DisplayName = user.DisplayName
				foundUser.Email = user.Email
				foundUser.Image = user.Image
				foundUser.Following = false
				foundUser.FollowsYou = false
				foundUsersList = append(foundUsersList, foundUser)
				continue
			}
			break
		}

	}

	//  START

	return foundUsersList, nil
}
func (r *repositoryFriendship) isReciprocated(friendshipList []domain.Friendship, user1Id string, user2Id string) string {
	var youBefriended string
	var youAreBefriended string
	for _, friendship := range friendshipList {
		if friendship.User1Id == user1Id && friendship.User2Id == user2Id {
			youBefriended = "yes"
		}
		if friendship.User1Id == user2Id && friendship.User2Id == user1Id {
			youAreBefriended = "yes"
		}
	}
	if youBefriended == "yes" && youAreBefriended == "yes" {
		return "reciprocated"
	} else if youAreBefriended == "yes" && youBefriended != "yes" {
		return "you follow"
	} else if youAreBefriended != "yes" && youBefriended == "yes" {
		return "follows you"
	} else if youAreBefriended != "yes" && youBefriended != "yes" {
		return "unknowns"
	}
	return "unknowns"
}

func (r *repositoryFriendship) GetListFriendships() ([]domain.Friendship, error) {
	friendshipsListrows, err := r.db.Query(QueryGetAllFriendships)
	if err != nil {
		return []domain.Friendship{}, err
	}
	defer friendshipsListrows.Close()

	var friendshipsList []domain.Friendship
	for friendshipsListrows.Next() {
		var friendship domain.Friendship
		err = friendshipsListrows.Scan(&friendship.User1Id, &friendship.User2Id)
		if err != nil {
			return []domain.Friendship{}, err
		}
		friendshipsList = append(friendshipsList, friendship)
	}
	return friendshipsList, nil
}

/*



 */
