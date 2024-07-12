package spell

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

var (
	ErrPrepareStatement    = errors.New("error preparing statement")
	ErrGettingLastInsertId = errors.New("error getting last insert id")
	ErrNotFound            = errors.New("spell not found")
)

type spellMySqlRepository struct {
	db *sql.DB
}

func NewSpellRepository(db *sql.DB) RepositorySpell {
	return &spellMySqlRepository{db: db}
}

func (r *spellMySqlRepository) Create(spell dto.SpellDto) (domain.Spell, error) {
	statement, err := r.db.Prepare(QueryCreateSpell)
	if err != nil {
		return domain.Spell{}, ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(
		spell.Name,
		spell.Description,
		spell.Range,
		spell.Ritual,
		spell.Duration,
		spell.Concentration,
		spell.CastingTime,
		spell.Level,
		spell.DamageType,
		spell.DifficultyClass,
		spell.Aoe,
		spell.School,
	)
	if err != nil {
		return domain.Spell{}, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Spell{}, ErrGettingLastInsertId
	}

	var createdSpell domain.Spell
	createdSpell.SpellId = int(lastId)
	createdSpell.Name = spell.Name
	createdSpell.Description = spell.Description
	createdSpell.Range = spell.Range
	createdSpell.Ritual = spell.Ritual
	createdSpell.Duration = spell.Duration
	createdSpell.Concentration = spell.Concentration
	createdSpell.CastingTime = spell.CastingTime
	createdSpell.Level = spell.Level
	createdSpell.DamageType = spell.DamageType
	createdSpell.DifficultyClass = spell.DifficultyClass
	createdSpell.Aoe = spell.Aoe
	createdSpell.School = spell.School

	return createdSpell, nil
}

func (r *spellMySqlRepository) GetAll() ([]domain.Spell, error) {
	rows, err := r.db.Query(QueryGetAll)
	if err != nil {
		return []domain.Spell{}, err
	}
	spells := []domain.Spell{}
	for rows.Next() {
		var spell domain.Spell
		if err := rows.Scan(
			&spell.SpellId,
			&spell.Name,
			&spell.Description,
			&spell.Range,
			&spell.Ritual,
			&spell.Duration,
			&spell.Concentration,
			&spell.CastingTime,
			&spell.Level,
			&spell.DamageType,
			&spell.DifficultyClass,
			&spell.Aoe,
			&spell.School); err != nil {
			return nil, err
		}
		spells = append(spells, spell)
	}

	return spells, nil
}

func (r *spellMySqlRepository) GetById(id int) (domain.Spell, error) {
	var spell domain.Spell
	err := r.db.QueryRow(QueryGetById, id).Scan(
		&spell.SpellId,
		&spell.Name,
		&spell.Description,
		&spell.Range,
		&spell.Ritual,
		&spell.Duration,
		&spell.Concentration,
		&spell.CastingTime,
		&spell.Level,
		&spell.DamageType,
		&spell.DifficultyClass,
		&spell.Aoe,
		&spell.School)
	if err != nil {
		fmt.Println(err)
		return domain.Spell{}, err
	}
	return spell, nil
}

func (r *spellMySqlRepository) Update(spell dto.SpellDto, id int) (domain.Spell, error) {
	statement, err := r.db.Prepare(QueryUpdate)
	if err != nil {

		return domain.Spell{}, ErrPrepareStatement
	}
	defer statement.Close()
	_, err = statement.Exec(&spell.Name, &spell.Description, &spell.Range, &spell.Ritual, &spell.Duration, &spell.Concentration, &spell.CastingTime, &spell.Level, &spell.DamageType, &spell.DifficultyClass, &spell.Aoe, &spell.School, id)
	if err != nil {
		return domain.Spell{}, err
	}
	var updatedSpell domain.Spell
	updatedSpell.SpellId = id
	updatedSpell.Name = spell.Name
	updatedSpell.Description = spell.Description
	updatedSpell.Range = spell.Range
	updatedSpell.Ritual = spell.Ritual
	updatedSpell.Duration = spell.Duration
	updatedSpell.Concentration = spell.Concentration
	updatedSpell.CastingTime = spell.CastingTime
	updatedSpell.Level = spell.Level
	updatedSpell.DamageType = spell.DamageType
	updatedSpell.DifficultyClass = spell.DifficultyClass
	updatedSpell.Aoe = spell.Aoe
	updatedSpell.School = spell.School

	return updatedSpell, nil
}

func (r *spellMySqlRepository) Delete(id int) error {
	statement, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return ErrPrepareStatement
	}
	defer statement.Close()

	result, err := statement.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNotFound
	}

	return nil
}

func (r *spellMySqlRepository) GetByCharacterDataId(characterId int) ([]domain.Spell, error) {
	rows, err := r.db.Query(QueryGetByCharacterDataId, characterId)

	if err != nil {
		return []domain.Spell{}, err
	}
	spells := []domain.Spell{}
	for rows.Next() {
		var spell domain.Spell
		if err := rows.Scan(
			&spell.SpellId,
			&spell.Name,
			&spell.Description,
			&spell.Range,
			&spell.Ritual,
			&spell.Duration,
			&spell.Concentration,
			&spell.CastingTime,
			&spell.Level,
			&spell.DamageType,
			&spell.DifficultyClass,
			&spell.Aoe,
			&spell.School); err != nil {
			return nil, err
		}
		spells = append(spells, spell)
	}

	return spells, nil
}

func (r *spellMySqlRepository) GetByClassId(classId int) ([]domain.Spell, error) {
	rows, err := r.db.Query(QueryGetByClassId, classId)

	if err != nil {
		return []domain.Spell{}, err
	}
	spells := []domain.Spell{}
	for rows.Next() {
		var spell domain.Spell
		err := rows.Scan(
			&spell.SpellId,
			&spell.Name,
			&spell.Description,
			&spell.Range,
			&spell.Ritual,
			&spell.Duration,
			&spell.Concentration,
			&spell.CastingTime,
			&spell.Level,
			&spell.DamageType,
			&spell.DifficultyClass,
			&spell.Aoe,
			&spell.School)

		if err != nil {
			return []domain.Spell{}, err
		}
		spells = append(spells, spell)
	}

	return spells, nil
}
