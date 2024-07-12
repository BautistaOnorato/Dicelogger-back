package user_campaign

import (
	"database/sql"
	"errors"
	"log"

	"github.com/proyecto-dnd/backend/internal/domain"
)

var (
	ErrPrepareStatement = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type userCampaignMySqlRepository struct {
	db *sql.DB
}

func NewUserCampaignRepository(db *sql.DB) UserCampaignRepository {
	return &userCampaignMySqlRepository{db: db}
}

func (r *userCampaignMySqlRepository) Create(userCampaign domain.UserCampaign) (domain.UserCampaign, error) {
	statement, err := r.db.Prepare(QueryCreateUserCampaign)
	if err != nil {
		log.Println(1, err)
		return domain.UserCampaign{}, ErrPrepareStatement
	}
	
	defer statement.Close()
	result, err := statement.Exec(
		userCampaign.CampaignId,
		userCampaign.UserId,
		userCampaign.CharacterId,
		userCampaign.IsOwner,
	)
	if err != nil {
		log.Println(2, err)
		return domain.UserCampaign{}, err
	}
	
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Println(3, err)
		return domain.UserCampaign{}, ErrGettingLastInsertId
	}
	userCampaign.UserCampaignId = int(lastId)

	return userCampaign, nil
}

func (r *userCampaignMySqlRepository) GetAll() ([]domain.UserCampaign, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userCampaigns []domain.UserCampaign
	for rows.Next() {
		var userCampaign domain.UserCampaign
		err := rows.Scan(
			&userCampaign.UserCampaignId,
			&userCampaign.CampaignId,
			&userCampaign.UserId,
			&userCampaign.CharacterId,
			&userCampaign.IsOwner,
		)
		if err != nil {
			return nil, err
		}
		userCampaigns = append(userCampaigns, userCampaign)
	}

	return userCampaigns, nil
}

func (r *userCampaignMySqlRepository) GetById(id int) (domain.UserCampaign, error) {
	var userCampaign domain.UserCampaign
	err := r.db.QueryRow(QueryGetById, id).Scan(
		&userCampaign.UserCampaignId,
		&userCampaign.CampaignId,
		&userCampaign.UserId,
		&userCampaign.CharacterId,
		&userCampaign.IsOwner,
	)
	if err != nil {
		return domain.UserCampaign{}, err
	}

	return userCampaign, nil
}

func (r *userCampaignMySqlRepository) GetByCampaignId(id int) ([]domain.UserCampaign, error) {
	rows, err := r.db.Query(QueryGetByCampaignId, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var userCampaigns []domain.UserCampaign
	for rows.Next() {
		var userCampaign domain.UserCampaign
		err := rows.Scan(
			&userCampaign.UserCampaignId,
			&userCampaign.CampaignId,
			&userCampaign.UserId,
			&userCampaign.CharacterId,
			&userCampaign.IsOwner,
		)
		if err != nil {
			return nil, err
		}

		userCampaigns = append(userCampaigns, userCampaign)
	}

	return userCampaigns, nil
}


func (r *userCampaignMySqlRepository) GetByUserId(id string) ([]domain.UserCampaign, error) {
	rows, err := r.db.Query(QueryGetByUserId, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var userCampaigns []domain.UserCampaign
	for rows.Next() {
		var userCampaign domain.UserCampaign
		err := rows.Scan(
			&userCampaign.UserCampaignId,
			&userCampaign.CampaignId,
			&userCampaign.UserId,
			&userCampaign.CharacterId,
			&userCampaign.IsOwner,
		)
		if err != nil {
			return nil, err
		}
		userCampaigns = append(userCampaigns, userCampaign)
	}

	return userCampaigns, nil
}

func (r *userCampaignMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userCampaignMySqlRepository) DeleteUserCampaignByCampaignId(id int) error {
	statement, err := r.db.Prepare(QueryDeleteByCampaignID)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userCampaignMySqlRepository) AddFriendsToUserCampaign(userIds []string, campaignId int) (error) {
	sqlQuery := QueryCreateMultipleUserCampaign
	vals := []interface{}{}

	for _, userId  := range userIds {
		sqlQuery += "(?, ?, ?, ?),"
		vals = append(vals, campaignId, userId, nil, 0)
	}
	
	sqlQuery = sqlQuery[0:len(sqlQuery) - 1]

	statement, err := r.db.Prepare(sqlQuery)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(vals...)
	if err != nil {
		return err
	}

	return nil
}

func (r *userCampaignMySqlRepository) AddCharacterToCampaign(characterId int, campaignId int, userId string) (error) {
	statement, err := r.db.Prepare(QueryUpdateCharacterCampaign)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(campaignId, characterId)
	if err != nil {
		return  err
	}

	statement, err = r.db.Prepare(QueryUpdateUserCharacter)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(characterId, userId, campaignId)
	if err != nil {
		return  err
	}

	return nil
}