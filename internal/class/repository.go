package class

import (
	"database/sql"
	"errors"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
)

type repositoryMysqlRepository struct {
	db *sql.DB
}

func NewClassRepository(db *sql.DB) RepositoryCharacterClass {
	return &repositoryMysqlRepository{db: db}
}

func (r *repositoryMysqlRepository) Create(classDto dto.ClassDto) (domain.Class, error) {
	statement, err := r.db.Prepare(QueryInsertClass)
	if err != nil {
		return domain.Class{}, ErrPrepareStatement
	}
	defer statement.Close()
	result, err := statement.Exec(
		classDto.Name,
		classDto.Description,
		classDto.ProficiencyBonus,
		classDto.HitDice,
		classDto.ArmorProficiencies,
		classDto.WeaponProficiencies,
		classDto.ToolProficiencies,
		classDto.SpellcastingAbility,
	)
	if err != nil {
		return domain.Class{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Class{}, ErrGettingLastInsertId
	}

	var createdClass domain.Class
	createdClass.ClassId = int(lastId)
	createdClass.Description = classDto.Description
	createdClass.ArmorProficiencies = classDto.ArmorProficiencies
	createdClass.HitDice = classDto.HitDice
	createdClass.ToolProficiencies = classDto.ToolProficiencies
	createdClass.ProficiencyBonus = classDto.ProficiencyBonus
	createdClass.WeaponProficiencies = classDto.WeaponProficiencies
	createdClass.SpellcastingAbility = classDto.SpellcastingAbility

	return createdClass, nil
}

func (r *repositoryMysqlRepository) GetAll() ([]domain.Class, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Class{}, err
	}
	defer rows.Close()

	var classes []domain.Class

	for rows.Next() {
		var class domain.Class
		if err := rows.Scan(&class.ClassId, &class.Name, &class.Description, &class.ProficiencyBonus, &class.HitDice, &class.ArmorProficiencies, &class.WeaponProficiencies, &class.ToolProficiencies, &class.SpellcastingAbility); err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (r *repositoryMysqlRepository) GetById(id int) (domain.Class, error) {
	var class domain.Class
	if err := r.db.QueryRow(QueryGetById, id).Scan(&class.ClassId, &class.Name, &class.Description, &class.ProficiencyBonus, &class.HitDice, &class.ArmorProficiencies, &class.WeaponProficiencies, &class.ToolProficiencies, &class.SpellcastingAbility); err != nil {
		return domain.Class{}, err
	}
	return class, nil
}

func (r *repositoryMysqlRepository) Update(classDto dto.ClassDto, id int) (domain.Class, error) {
	statement, err := r.db.Prepare(QueryUpdateClass)
	if err != nil {
		return domain.Class{}, err
	}
	defer statement.Close()

	_, err = statement.Exec(classDto.Name, classDto.Description, classDto.ProficiencyBonus, classDto.HitDice, classDto.ArmorProficiencies, classDto.WeaponProficiencies, classDto.ToolProficiencies, classDto.SpellcastingAbility, id)
	if err != nil {
		return domain.Class{}, nil
	}

	var updatedClass domain.Class
	updatedClass.ClassId = id
	updatedClass.Description = classDto.Description
	updatedClass.ArmorProficiencies = classDto.ArmorProficiencies
	updatedClass.HitDice = classDto.HitDice
	updatedClass.ToolProficiencies = classDto.ToolProficiencies
	updatedClass.ProficiencyBonus = classDto.ProficiencyBonus
	updatedClass.WeaponProficiencies = classDto.WeaponProficiencies
	updatedClass.SpellcastingAbility = classDto.SpellcastingAbility

	return updatedClass, nil
}

func (r *repositoryMysqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDeleteClass)
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
