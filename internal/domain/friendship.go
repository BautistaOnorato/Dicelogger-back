package domain

type Friendship struct {
	User1Id string `json:"user1_id"`
	User2Id string `json:"user2_id"`
}

type Mutuals struct {
	User1Id   string `json:"user1_id"`
	User2Name string `json:"user2_name"`
}

type FriendUserData struct {
	UserId      string  `json:"user_id"`
	Username    string  `json:"username"`
	DisplayName string  `json:"display_name"`
	Image       *string `json:"image"`
	Email       string  `json:"email"`
	Following   bool    `json:"following"`
	FollowsYou  bool    `json:"follows_you"`
}
