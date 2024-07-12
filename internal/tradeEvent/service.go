package tradeevent

import (
	"errors"
	"fmt"

	"github.com/proyecto-dnd/backend/internal/armorXCharacterData"
	charactertrade "github.com/proyecto-dnd/backend/internal/characterTrade"
	"github.com/proyecto-dnd/backend/internal/domain"
	itemxcharacterdata "github.com/proyecto-dnd/backend/internal/itemXCharacterData"
	weaponxcharacterdata "github.com/proyecto-dnd/backend/internal/weaponXCharacterData"
)

var (
	ErrCannotBeNegative = errors.New("cannot be negative")
)

type serviceTradeEvent struct {
	tradeEventRepo          RepositoryTradeEvent
	characterTradeService   charactertrade.ServiceCharacterTrade
	weaponXCharacterService weaponxcharacterdata.ServiceWeaponXCharacterData
	armorXCharacterService  armorXCharacterData.ServiceArmorXCharacterData
	itemXCharacterService   itemxcharacterdata.ServiceItemXCharacterData
}

func NewTradeEventService(tradeEventRepo RepositoryTradeEvent, characterTradeService charactertrade.ServiceCharacterTrade, weaponService weaponxcharacterdata.ServiceWeaponXCharacterData, armorService armorXCharacterData.ServiceArmorXCharacterData, itemService itemxcharacterdata.ServiceItemXCharacterData) ServiceTradeEvent {
	return &serviceTradeEvent{tradeEventRepo, characterTradeService, weaponService, armorService, itemService}
}

// DeleteBySenderOrReciever implements ServiceTradeEvent.
func (s *serviceTradeEvent) DeleteBySenderOrReciever(id int) error {
	events, err := s.GetBySenderOrReciever(id)
	if err != nil {
		return err
	}
	for i := range events {
		err = s.Delete(events[i].TradeEvent_Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// Create implements ServiceTradeEvent.
// TO DO: Implement method in itemXCharacterData to search by characterData_Id and Item_Id, must also be implemented in itemXcharacterData's create to prevent duping of entries
func (s *serviceTradeEvent) Create(tradeEvent domain.TradeEvent) (domain.TradeEvent, error) {
	newTradeEvent, err := s.tradeEventRepo.Create(tradeEvent)
	if err != nil {
		return domain.TradeEvent{}, err
	}

	for i, tradingItems := range tradeEvent.TradingItems {
		tradeEvent.TradingItems[i].TradeEvent_Id = newTradeEvent.TradeEvent_Id
		if tradingItems.WeaponXCharacter != nil {
			err = s.weaponXCharacterService.UpdateOwnership(domain.WeaponXCharacterData{Character_Weapon_Id: *tradingItems.WeaponXCharacter, CharacterData_Id: tradingItems.ItemReciever, Weapon: domain.Weapon{}, Equipped: false})
			fmt.Println("updating weapon")
			if err != nil {
				return domain.TradeEvent{}, err
			}
		}
		if tradingItems.ArmorXCharacter != nil {
			fmt.Println("updating armor")
			err = s.armorXCharacterService.UpdateOwnership(domain.ArmorXCharacterData{ArmorXCharacterData_Id: *tradingItems.ArmorXCharacter, Armor: domain.Armor{}, CharacterData_Id: tradingItems.ItemReciever, Equipped: false})
			if err != nil {
				return domain.TradeEvent{}, err
			}
		}

		if tradingItems.ItemXCharacter != nil {
			itemXCharacterToUpdate, err := s.itemXCharacterService.GetById(*tradingItems.ItemXCharacter)
			if err != nil {
				return domain.TradeEvent{}, err
			}
			if itemXCharacterToUpdate.Quantity < *tradingItems.Quantity {
				return domain.TradeEvent{}, ErrCannotBeNegative
			}
			if itemXCharacterToUpdate.Quantity == *tradingItems.Quantity {
				err = s.itemXCharacterService.UpdateOwnership(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.ItemXCharacter, CharacterData_Id: tradeEvent.Receiver, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: itemXCharacterToUpdate.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
			} else {
				_, err = s.itemXCharacterService.Update(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.ItemXCharacter, CharacterData_Id: tradeEvent.Sender, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: itemXCharacterToUpdate.Quantity - *tradingItems.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
				_, err = s.itemXCharacterService.Create(domain.ItemXCharacterData{Character_Item_Id: *tradingItems.ItemXCharacter, CharacterData_Id: tradeEvent.Receiver, Item: domain.Item{Item_Id: itemXCharacterToUpdate.Item.Item_Id}, Quantity: *tradingItems.Quantity})
				if err != nil {
					return domain.TradeEvent{}, err
				}
			}

		}
	}
	err = s.characterTradeService.BulkCreateCharacterTrade(tradeEvent.TradingItems)
	if err != nil {
		return domain.TradeEvent{}, err
	}
	newTradeEvent.TradingItems, err = s.characterTradeService.GetByTradeEventId(newTradeEvent.TradeEvent_Id)
	if err != nil {
		return domain.TradeEvent{}, err
	}
	return newTradeEvent, nil
}

// Delete implements ServiceTradeEvent.
func (s *serviceTradeEvent) Delete(id int) error {
	s.characterTradeService.DeleteByTradeEventId(id)
	return s.tradeEventRepo.Delete(id)
}

// GetByReceiver implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetByReceiver(receiver int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetByReceiver(receiver)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySender implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySender(sender int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySender(sender)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySenderOrReciever implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySenderOrReciever(id int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySenderOrReciever(id)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

// GetBySessionId implements ServiceTradeEvent.
func (s *serviceTradeEvent) GetBySessionId(sessionId int) ([]domain.TradeEvent, error) {
	tradeEvents, err := s.tradeEventRepo.GetBySessionId(sessionId)
	if err != nil {
		return []domain.TradeEvent{}, err
	}
	for i := range tradeEvents {
		tradeEvents[i].TradingItems, err = s.characterTradeService.GetByTradeEventId(tradeEvents[i].TradeEvent_Id)
		if err != nil {
			return []domain.TradeEvent{}, err
		}
	}
	return tradeEvents, nil
}

