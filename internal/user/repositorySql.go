package user

// import (
// 	"database/sql"
// 	"errors"
// 	"strconv"
// 	"strings"

// 	"github.com/proyecto-dnd/backend/internal/domain"
// )

// var (
// 	ErrPrepareStatement = errors.New("error prepare statement")
// 	ErrExecStatement    = errors.New("error exec statement")
// 	ErrLastInsertedId   = errors.New("error last inserted id")
// 	ErrEmpty            = errors.New("empty list")
// 	ErrNotFound         = errors.New("dentist not found")
// )

// type repositoryUserSql struct {
// 	db *sql.DB
// }

// func NewUserSqlRepository(db *sql.DB) RepositoryUsers {
// 	return &repositoryUserSql{db: db}
// }

// func (r *repositoryUserSql) Create(user domain.User) (domain.User, error) {
// 	statement, err := r.db.Prepare(QueryInsertUser)
// 	if err != nil {
// 		return domain.User{}, ErrPrepareStatement
// 	}

// 	defer statement.Close()

// 	result, err := statement.Exec(
// 		user.Username,
// 		user.Email,
// 		user.Password,
// 	)
// 	if err != nil {
// 		return domain.User{}, ErrExecStatement
// 	}

// 	lastId, err := result.LastInsertId()
// 	if err != nil {
// 		return domain.User{}, ErrLastInsertedId
// 	}

// 	user.Id = strconv.FormatInt(lastId, 10)

// 	return user, nil
// }
// func (r *repositoryUserSql) GetAll() ([]domain.User, error) {
// 	rows, err := r.db.Query(QueryGetAllUsers)
// 	if err != nil {
// 		return []domain.User{}, err
// 	}
// 	defer rows.Close()

// 	var users []domain.User

// 	for rows.Next() {
// 		var user domain.User
// 		err := rows.Scan(
// 			&user.Id,
// 			&user.Username,
// 			&user.Email,
// 			&user.Password,
// 		)
// 		if err != nil {
// 			return []domain.User{}, err
// 		}
// 		users = append(users, user)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return []domain.User{}, err
// 	}

// 	return users, nil
// }

// func (r *repositoryUserSql) GetById(id string) (domain.User, error) {
// 	row := r.db.QueryRow(QueryGetUserById, id)

// 	var user domain.User
// 	err := row.Scan(
// 		&user.Id,
// 		&user.Username,
// 		&user.Email,
// 		&user.Password,
// 	)
// 	if err != nil {
// 		return domain.User{}, err
// 	}
// 	return user, nil
// }

// func (r *repositoryUserSql) Update(user domain.User, id string) (domain.User, error) {
// 	result, err := r.db.Exec(QueryUpdateUser,
// 		id,
// 		user.Username,
// 		user.Email,
// 		user.Password,
// 	)
// 	if err != nil {
// 		return domain.User{}, err
// 	}
// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return domain.User{}, ErrEmpty
// 	}

// 	user.Id = id

// 	return user, nil
// }

// func (r *repositoryUserSql) Delete(id string) error {
// 	result, err := r.db.Exec(QueryDeleteUser, id)
// 	if err != nil {
// 		return err
// 	}

// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return err
// 	}

// 	if rowsAffected < 1 {
// 		return ErrNotFound
// 	}
// 	return nil
// }

// func (r *repositoryUserSql) Patch(user domain.User, id string) (domain.User, error) {
// 	var fieldsToUpdate []string
// 	var args []interface{}

// 	if user.Username != "" {
// 		fieldsToUpdate = append(fieldsToUpdate, "username = ?")
// 		args = append(args, user.Username)
// 	}

// 	if user.Email != "" {
// 		fieldsToUpdate = append(fieldsToUpdate, "email = ?")
// 		args = append(args, user.Email)
// 	}

// 	if user.Password != "" {
// 		fieldsToUpdate = append(fieldsToUpdate, "password = ?")
// 		args = append(args, user.Password)
// 	}

// 	if len(fieldsToUpdate) == 0 {
// 		return domain.User{}, ErrEmpty
// 	}

// 	queryString := "UPDATE user SET " + strings.Join(fieldsToUpdate, ", ") + " WHERE user_id = ?"
// 	args = append(args, id)

// 	result, err := r.db.Exec(queryString, args...)
// 	if err != nil {
// 		return domain.User{}, err
// 	}

// 	_, err = result.RowsAffected()
// 	if err != nil {
// 		return domain.User{}, ErrEmpty
// 	}

// 	return r.GetById(id)
// }
