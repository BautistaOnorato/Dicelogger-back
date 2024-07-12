package user

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/pkg/email"
	"google.golang.org/api/iterator"
)

var (
	ctx         = &gin.Context{}
	ErrEmpty    = errors.New("empty list")
	ErrNotFound = errors.New("user not found")
)

type repositoryFirebase struct {
	app        *firebase.App
	authClient *auth.Client
	db         *sql.DB
}

func NewUserFirebaseRepository(app *firebase.App, db *sql.DB) RepositoryUsers {
	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Printf("Error initializing Firebase Auth client: %v", err)
	}
	return &repositoryFirebase{app: app, authClient: authClient, db: db}
}

func (r *repositoryFirebase) Create(user domain.User) (domain.UserResponse, error) {
	log.Println("Register")
	//sql backup
	statement, err := r.db.Prepare(QueryInsertUser)
	if err != nil {
		log.Println(1, err)
		return domain.UserResponse{}, err
	}
	defer statement.Close()
	
	params := (&auth.UserToCreate{}).
	Email(user.Email).
	Password(user.Password).
	DisplayName(user.Username).
	Disabled(false).
	PhotoURL("https://proyecto-dnd.vercel.app/user.png")
	
	//firebase create
	newUser, err := r.authClient.CreateUser(ctx, params)
	if err != nil {
		log.Println(2, err)
		log.Printf("Error creating user: %v", err)
	}
	
	claims := map[string]interface{}{"displayName": user.DisplayName, "subExpiration": time.Unix(0, 0).String()}
	
	err = r.authClient.SetCustomUserClaims(ctx, newUser.UID, claims)
	if err != nil {
		log.Println(3, err)
		fmt.Println("Error setting custom user claims.")
		return domain.UserResponse{}, err
	}
	
	//email verification
	emailVerificationLink, err := r.authClient.EmailVerificationLink(ctx, user.Email)
	if err != nil {
		log.Println(4, err)
		log.Printf("Error creating email verification link: %v", err)
	}
	
	err = email.SendEmailVerificationLink(user.Email, emailVerificationLink)
	if err != nil {
		log.Println(5, err)
		log.Printf("Error sending email verification link: %v", err)
	}
	
	//sql create
	_, err = statement.Exec(newUser.UID, user.Username, user.Email, user.Password, user.DisplayName, "https://proyecto-dnd.vercel.app/user.png")
	if err != nil {
		log.Println(6, err)
		fmt.Println("Error setting custom user claims.")
		return domain.UserResponse{}, err
	}
	
	var userTemp domain.UserResponse
	userTemp.Username = newUser.DisplayName
	userTemp.Email = newUser.Email
	userTemp.Id = newUser.UID
	userTemp.DisplayName = user.DisplayName
	// newUser.PasswordHash = user.Password

	return userTemp, nil
}
func (r *repositoryFirebase) GetAll() ([]domain.UserResponse, error) {

	// var user domain.UserResponse
	var users []domain.UserResponse
	// pager := iterator.NewPager(r.authClient.Users(ctx, ""), 100, "")

	// for {
	// 	var authUsers []*auth.ExportedUserRecord
	// 	nextPageToken, err := pager.NextPage(&authUsers)
	// 	if err != nil {
	// 		log.Printf("paging error %v\n", err)
	// 	}
	// 	for _, u := range authUsers {
	// 		user.Username = u.DisplayName
	// 		user.Email = u.Email
	// 		// user.Password = u.PasswordHash
	// 		user.Id = u.UID
	// 		users = append(users, user)
	// 	}
	// 	if nextPageToken == "" {
	// 		break
	// 	}
	// }

	rows, err := r.db.Query(QueryGetAllUsers)
	if err != nil {
		return []domain.UserResponse{}, err
	}
	defer rows.Close()
	for rows.Next() {
		var user domain.UserResponse
		if err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.DisplayName, &user.Image); err != nil {
			return []domain.UserResponse{}, err
		}

		users = append(users, user)
	}

	return users, nil
}
func (r *repositoryFirebase) GetByName(name string) ([]domain.User, error) {
	var user domain.User
	var users []domain.User
	pager := iterator.NewPager(r.authClient.Users(ctx, ""), 50, "")
	for {
		var authUsers []*auth.ExportedUserRecord
		nextPageToken, err := pager.NextPage(&authUsers)
		if err != nil {
			log.Printf("paging error %v\n", err)
		}
		for _, u := range authUsers {
			if u.DisplayName == name {
				user.Username = u.DisplayName
				user.Id = u.UID
				users = append(users, user)
			}
		}
		if nextPageToken == "" {
			break
		}
	}
	return users, nil
}

func (r *repositoryFirebase) GetById(id string) (domain.UserResponse, error) {

	_, err := r.authClient.GetUser(ctx, id)
	if err != nil {
		//TODO RETURN ERROR
		log.Printf("error getting user %s: %v\n", id, err)
	}

	row, err := r.db.Query(QueryGetUserById, id)
	if err != nil {
		return domain.UserResponse{}, err
	}
	defer row.Close()

	var user domain.UserResponse
	for row.Next() {
		if err := row.Scan(&user.Id, &user.Username, &user.Email, &user.DisplayName, &user.Image); err != nil {
			return domain.UserResponse{}, err
		}
	}
	// fmt.Println(user)
	// var user domain.User
	// user.Username = u.DisplayName
	// user.Email = u.Email
	// user.Id = u.UID

	return user, nil
}
func (r *repositoryFirebase) Update(user domain.UserUpdate, id string) (domain.UserUpdate, error) {
	params := (&auth.UserToUpdate{}).
		Email(user.Email).
		Password(user.Password).
		DisplayName(user.Username)
	_, err := r.authClient.UpdateUser(ctx, id, params)
	if err != nil {
		//TODO RETURN ERROR
		log.Printf("error updating user: %v\n", err)
	}
	// log.Printf("Successfully updated user: %v\n", u)
	result, err := r.db.Exec(QueryUpdateUser,
		user.Username,
		user.Email,
		user.Password,
		user.Image,
		user.DisplayName,
		id,
	)
	if err != nil {
		return domain.UserUpdate{}, err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return domain.UserUpdate{}, err
	}
	user.Id = id

	return user, nil
}
func (r *repositoryFirebase) Delete(id string) error {
	// userList, err := r.GetAll()
	// if err != nil {
	// 	return err
	// }
	// //extract uid from userList

	// var idList []string
	// for _, u := range userList {
	// 	id = u.Id
	// 	idList = append(idList, id)
	// }
	// r.authClient.DeleteUsers(ctx, idList)

	// fmt.Println("https://media1.tenor.com/m/LBWyQg647MoAAAAC/execute-order66-order66.gif")

	err := r.authClient.DeleteUser(ctx, id)
	if err != nil {
		log.Printf("error deleting user: %v\n", err)
	}

	result, err := r.db.Exec(QueryDeleteUser, id)
	if err != nil {
		return err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return err
	}

	log.Printf("Successfully deleted user: %s\n", id)

	return nil
}

func (r *repositoryFirebase) Patch(user domain.UserUpdate, id string) (domain.UserResponse, error) {
	var fieldsToUpdate []string
	var args []interface{}

	if user.Username != "" {
		fieldsToUpdate = append(fieldsToUpdate, "name = ?")
		args = append(args, user.Username)
		_, err := r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).DisplayName(user.Username))
		if err != nil {
			return domain.UserResponse{}, err
		}
	}
	if user.Email != "" {
		fieldsToUpdate = append(fieldsToUpdate, "email = ?")
		args = append(args, user.Email)
		_, err := r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).Email(user.Email))
		if err != nil {
			return domain.UserResponse{}, err
		}
	}
	if user.Password != "" {
		fieldsToUpdate = append(fieldsToUpdate, "password = ?")
		args = append(args, user.Password)
		_, err := r.authClient.UpdateUser(ctx, id, (&auth.UserToUpdate{}).Password(user.Password))
		if err != nil {
			return domain.UserResponse{}, err
		}
	}
	if user.Image != nil && *user.Image != "" {
		fieldsToUpdate = append(fieldsToUpdate, "image = ?")
		args = append(args, user.Image)
	}
	if user.DisplayName != "" {
		fieldsToUpdate = append(fieldsToUpdate, "display_name = ?")
		args = append(args, user.DisplayName)
	}

	if len(fieldsToUpdate) == 0 {
		return domain.UserResponse{}, ErrEmpty
	}

	queryString := "UPDATE user SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE uid = ?"
	args = append(args, id)

	fmt.Println(queryString)

	patchStatement, err := r.db.Exec(queryString, args...)
	if err != nil {
		return domain.UserResponse{}, err
	}

	_, err = patchStatement.RowsAffected()
	if err != nil {
		return domain.UserResponse{}, err
	}

	patchedUser, err := r.GetById(id)
	if err != nil {
		return domain.UserResponse{}, err
	}

	return patchedUser, nil
}

func (r *repositoryFirebase) Login(userInfo domain.UserLoginInfo) (string, error) {

	expiresIn := time.Hour * 24 * 2

	cookie, err := r.authClient.SessionCookie(ctx, userInfo.IdToken, expiresIn)
	if err != nil {
		fmt.Printf("error creating session cookie: %v\n", err)
		return "error creating session cookie", err
	}

	return cookie, nil
}

func (r *repositoryFirebase) GetJwtInfo(cookieToken string) (domain.UserTokenClaims, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(cookieToken, jwt.MapClaims{})
	if err != nil {
		return domain.UserTokenClaims{}, err
	}

	var tokenClaims domain.UserTokenClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["user_id"] != nil {
			tokenClaims.Id = claims["user_id"].(string)
		} else {
			tokenClaims.Id = claims["claims"].(map[string]interface{})["user_id"].(string)
		}
		// if claims["name"] != nil {
		// 	tokenClaims.Username = claims["name"].(string)
		// } else {
		// 	tokenClaims.Username = claims["claims"].(map[string]interface{})["name"].(string)
		// }
		// if claims["email"] != nil {
		// 	tokenClaims.Email = claims["email"].(string)
		// } else {
		// 	tokenClaims.Email = claims["claims"].(map[string]interface{})["email"].(string)
		// }
		// if claims["displayName"] != nil {
		// 	tokenClaims.DisplayName = claims["displayName"].(string)
		// } else {
		// 	tokenClaims.DisplayName = claims["claims"].(map[string]interface{})["displayName"].(string)
		// }
		// if claims["subExpiration"] != nil {
		// 	tokenClaims.SubExpirationDate = claims["subExpiration"].(string)
		// } else {
		// 	tokenClaims.SubExpirationDate = claims["claims"].(map[string]interface{})["subExpiration"].(string)
		// }
	}
	userData, err := r.getFullDataById(tokenClaims.Id)
	if err != nil {
		return domain.UserTokenClaims{}, err
	}
	userAuthData, err := r.authClient.GetUser(ctx, tokenClaims.Id)
	if err != nil {
		return domain.UserTokenClaims{}, err
	}
	var tokenInfo domain.UserTokenClaims
	tokenInfo.Id = userData.Id
	tokenInfo.Username = userData.Username
	tokenInfo.Email = userData.Email
	tokenInfo.DisplayName = userData.DisplayName
	tokenInfo.SubExpirationDate = userData.SubExpirationDate
	tokenInfo.Image = *userData.Image
	tokenInfo.EmailVerified = userAuthData.EmailVerified
	return tokenInfo, nil
}

func (r *repositoryFirebase) TransferDataToSql(users []domain.User) (string, error) {

	insertString, err := r.BulkInsertString(users)

	if err != nil {
		return "", err
	}

	result, err := r.db.Exec(insertString)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	if rowsAffected < 1 {
		return "", errors.New("no rows affected")
	}

	return insertString, nil
}

func (r *repositoryFirebase) BulkInsertString(users []domain.User) (string, error) {
	var values strings.Builder

	// values.WriteString("(")

	for i, user := range users {

		values.WriteString(fmt.Sprintf("('%s', '%s', '%s', '%s', '%s')", user.Id, user.Username, user.Email, user.Password, user.DisplayName))

		if i < len(users)-1 {
			values.WriteString(", ")
		}
	}

	// values.WriteString(")")

	insertSQL := fmt.Sprintf("INSERT INTO user (uid, name, email, password, display_name) VALUES %s;", values.String())
	return insertSQL, nil
}

func (r *repositoryFirebase) SubscribeToPremium(id string, date string) (string, error) {

	token, _, err := new(jwt.Parser).ParseUnverified(id, jwt.MapClaims{})
	if err != nil {
		return "", err
	}

	var tokenClaims domain.UserTokenClaims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if claims["user_id"] != nil {
			tokenClaims.Id = claims["user_id"].(string)
		} else {
			tokenClaims.Id = claims["claims"].(map[string]interface{})["user_id"].(string)

		}
	}

	statement, err := r.db.Prepare(QueryUpdateSubExpirationDate)
	if err != nil {
		return "couldn't sub", errors.New("Error preparing statement: " + err.Error())
	}
	defer statement.Close()

	_, err = statement.Exec(date, tokenClaims.Id)
	if err != nil {
		return "couldn't sub", errors.New("Error executing statement: " + err.Error())
	}

	return "Subbed succesfully", nil
}

func (r *repositoryFirebase) getFullDataById(id string) (domain.UserResponseFull, error) {
	statement, err := r.db.Prepare(QueryGetFullData)
	if err != nil {
		return domain.UserResponseFull{}, err
	}
	defer statement.Close()

	row := statement.QueryRow(id)
	var userResponseFull domain.UserResponseFull
	err = row.Scan(&userResponseFull.Id, &userResponseFull.Username, &userResponseFull.Email, &userResponseFull.DisplayName, &userResponseFull.Image, &userResponseFull.SubExpirationDate)
	if err != nil {
		return domain.UserResponseFull{}, err
	}
	return userResponseFull, nil
}
func (r *repositoryFirebase) CheckSubExpiration(userId string) error {
	statement, err := r.db.Prepare(QueryGetSubExpirationDate)
	if err != nil {
		log.Println("Error preparing statement:", err)
		return err
	}
	defer statement.Close()
	
	row := statement.QueryRow(userId)
	var subExpirationDate string
	err = row.Scan(&subExpirationDate)
	if err != nil {
		log.Println("Error scanning row:", err)
		return err
	}

	expirationDateParsed, err := time.Parse(time.RFC3339, subExpirationDate)
	if err != nil {
		log.Println("Error parsing expiration date:", err)
		ctx.JSON(500, "Error parsing expiration date from token claims:"+err.Error())
	}

	if !time.Now().Before(expirationDateParsed) {
		return errors.New("sub expired")
	}

	return nil
}

func (r *repositoryFirebase) SendVerificationEmail(emailAddress string) error {
	emailVerificationLink, err := r.authClient.EmailVerificationLink(ctx, emailAddress)
	if err != nil {
		log.Printf("Error creating email verification link: %v", err)
		return err
	}

	err = email.SendEmailVerificationLink(emailAddress, emailVerificationLink)
	if err != nil {
		log.Printf("Error sending email verification link: %v", err)
		return err
	}
	return nil
}
