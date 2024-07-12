package charactertrade

import "github.com/proyecto-dnd/backend/internal/domain"

type characterTradeService struct {
	characterTradeRepository RepositoryCharacterTrade
}

func NewCharacterTradeService(characterTradeRepository RepositoryCharacterTrade) ServiceCharacterTrade {
	return &characterTradeService{characterTradeRepository: characterTradeRepository}
}

// BulkCreateCharacterTrade implements ServiceCharacterTrade.
func (c *characterTradeService) BulkCreateCharacterTrade(characterTradeList []domain.CharacterTrade) error {
	return c.characterTradeRepository.BulkCreateCharacterTrade(characterTradeList)
}

// DeleteByTradeEventId implements ServiceCharacterTrade.
func (c *characterTradeService) DeleteByTradeEventId(tradeEventId int) error {
	return c.characterTradeRepository.DeleteByTradeEventId(tradeEventId)
}

// GetByTradeEventId implements ServiceCharacterTrade.
func (c *characterTradeService) GetByTradeEventId(tradeEventId int) ([]domain.CharacterTrade, error) {
	return c.characterTradeRepository.GetByTradeEventId(tradeEventId)
}

