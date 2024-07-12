package backgroundXproficiency

import "github.com/proyecto-dnd/backend/internal/domain"

type BackgroundXProficiencyRepository interface {
	Create(backgroundXProficiency domain.BackgroundXProficiency) (domain.BackgroundXProficiency, error)
	Delete(backgroundXProficiency domain.BackgroundXProficiency) error
}

type BackgroundXProficiencyService interface {
	Create(backgroundXProficiency domain.BackgroundXProficiency) (domain.BackgroundXProficiency, error)
	Delete(backgroundXProficiency domain.BackgroundXProficiency) error
}
