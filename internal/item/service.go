package item

import "github.com/proyecto-dnd/backend/internal/domain"

type service struct {
	repo RepositoryItem
}

func NewItemService(repo RepositoryItem) ServiceItem {
    return &service{repo: repo}
}

// GetAllGeneric implements ServiceItem.
func (s *service) GetAllGeneric() ([]domain.Item, error) {
	items, err := s.repo.GetAllGeneric()
	if err != nil {
		return []domain.Item{}, err
	}
	return items, nil
}


// Create implements ServiceItem.
func (s *service) Create(item domain.Item) (domain.Item, error) {
	newItem, err := s.repo.Create(item)
	if err != nil {
		return domain.Item{}, err
	}
	return newItem, nil
}

// Delete implements ServiceItem.
func (s *service) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

// GetAll implements ServiceItem.
func (s *service) GetAll() ([]domain.Item, error) {
	items, err := s.repo.GetAll()
	if err != nil {
		return []domain.Item{}, err
	}
	return items, nil
}

// GetByCampaignId implements ServiceItem.
func (s *service) GetByCampaignId(campaignId int) ([]domain.Item, error) {
	items, err := s.repo.GetByCampaignId(campaignId)
	if err != nil {
		return []domain.Item{}, err
	}
	return items, nil
}

// GetById implements ServiceItem.
func (s *service) GetById(id int) (domain.Item, error) {
	item, err := s.repo.GetById(id)
	if err != nil {
		return domain.Item{}, err
	}
	return item, nil
}

// Update implements ServiceItem.
func (s *service) Update(item domain.Item) (domain.Item, error) {
	item, err := s.repo.Update(item)
	if err != nil {
		return domain.Item{}, err
	}
	return item, nil
}
