package tradeevent

import "github.com/proyecto-dnd/backend/internal/domain"

type RepositoryTradeEvent interface {
	Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error)
	GetBySessionId(sessionId int) ([]domain.TradeEvent, error)
	GetBySender(sender int) ([]domain.TradeEvent, error)
	GetByReceiver(receiver int) ([]domain.TradeEvent, error)
	GetBySenderOrReciever(id int) ([]domain.TradeEvent, error)
	Delete(id int) (error)
}

type ServiceTradeEvent interface {
	Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error)
	GetBySessionId(sessionId int) ([]domain.TradeEvent, error)
	GetBySender(sender int) ([]domain.TradeEvent, error)
	GetByReceiver(receiver int) ([]domain.TradeEvent, error)
	GetBySenderOrReciever(id int) ([]domain.TradeEvent, error)
	Delete(id int) (error)
	DeleteBySenderOrReciever(id int) (error)
}