package backgroundxSkill

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryBackgroundXSkills interface {
	CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error)
	GetByIdBackgroundXSkills(id int) (domain.BackgroundXSkills, error)
	GetByBackgroundId(id int) ([]domain.BackgroundXSkills, error)
	UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	DeleteBackgroundXSkills(id int) error
	DeleteByBackgroundId(id int) error
}

type ServiceBackgroundXSkills interface {
	CreateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	GetAllBackgroundXSkills() ([]domain.BackgroundXSkills, error)
	GetByIdBackgroundXSkills(id int) (domain.BackgroundXSkills, error)
	GetByBackgroundId(id int) ([]domain.BackgroundXSkills, error)
	UpdateBackgroundXSkills(data domain.BackgroundXSkills) (domain.BackgroundXSkills, error)
	DeleteBackgroundXSkills(id int) error
	DeleteByBackgroundId(id int) error
}
