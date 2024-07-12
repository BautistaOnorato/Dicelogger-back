package domain

type User struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"displayName"`
}

type UserResponse struct {
	Id          string  `json:"id"`
	Username    string  `json:"name"`
	Email       string  `json:"email"`
	DisplayName string  `json:"displayName"`
	Image       *string `json:"image"`
}
type UserUpdate struct {
	Id          string  `json:"id"`
	Username    string  `json:"name"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	DisplayName string  `json:"displayName"`
	Image       *string `json:"image"`
}
type UserLoginInfo struct {
	IdToken string `json:"idToken"`
}

type UserTokenClaims struct {
	Id                string `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	DisplayName       string `json:"displayName"`
	SubExpirationDate string `json:"subExpiration"`
	Image             string `json:"image"`
	EmailVerified     bool   `json:"emailVerified"`
}

type UserResponseFull struct {
	Id                string  `json:"id"`
	Username          string  `json:"name"`
	Email             string  `json:"email"`
	DisplayName       string  `json:"displayName"`
	Image             *string `json:"image"`
	SubExpirationDate string  `json:"subExpiration"`
}
