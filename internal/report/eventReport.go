package report

import (
	"bytes"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/proyecto-dnd/backend/internal/attackEvent"
	characterdata "github.com/proyecto-dnd/backend/internal/characterData"
	"github.com/proyecto-dnd/backend/internal/dice_event"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
	tradeevent "github.com/proyecto-dnd/backend/internal/tradeEvent"
)

type ReportGenerator struct {
	tradeEventService  tradeevent.ServiceTradeEvent
	attackEventService attackEvent.AttackEventService
	diceEventService   dice_event.DiceEventService
	characterDataService characterdata.ServiceCharacterData
}

func NewReportGenerator(tradeEventService tradeevent.ServiceTradeEvent, attackEventService attackEvent.AttackEventService, diceEventService dice_event.DiceEventService, characterDataService characterdata.ServiceCharacterData) *ReportGenerator {
	return &ReportGenerator{
		tradeEventService:  tradeEventService,
		attackEventService: attackEventService,
		diceEventService:   diceEventService,
		characterDataService: characterDataService,
	}
}

func (r *ReportGenerator) GenerateEventSessionReport(id int) (*bytes.Buffer, error) {
	excelFile := excelize.NewFile()
	tradeSheetIndex := excelFile.NewSheet("Sheet1")
	excelFile.SetSheetName("Sheet1", "Trade Events")
	characterTradeSheetIndex := excelFile.NewSheet("Character X Trades")
	attackEventSheetIndex := excelFile.NewSheet("Attack Event")
	affectedByAttackEventSheetIndex := excelFile.NewSheet("Affected By Attack Event")
	diceEventSheetIndex := excelFile.NewSheet("Dice Event")

	tradeEvents, err := r.tradeEventService.GetBySessionId(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	attackEvents, err := r.attackEventService.GetEventsBySessionId(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	diceEvents, err := r.diceEventService.GetBySessionId(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	generateCharacterTradeHeaders(excelFile)
	excelFile.SetActiveSheet(characterTradeSheetIndex)
	generateTradeEventHeaders(excelFile)
	excelFile.SetActiveSheet(tradeSheetIndex)
	characterTradeIndex := 0
	for index1, tradeEvent := range tradeEvents {
		excelFile.SetActiveSheet(characterTradeSheetIndex)
		for _, characterTrade := range tradeEvent.TradingItems {
			insertCharacterTradeRow(excelFile, characterTrade, characterTradeIndex)
			characterTradeIndex++
		}
		excelFile.SetActiveSheet(tradeSheetIndex)
		insertTradeEventRow(excelFile, tradeEvent, index1)
	}

	generateAttackEventHeaders(excelFile)
	generateAffectedByAttackEventHeaders(excelFile)
	affectedCharactersIndex := 0
	for index1, attackEvent := range attackEvents {
		excelFile.SetActiveSheet(affectedByAttackEventSheetIndex)
		for _, affectedByAttackEvent := range attackEvent.Affected {
			insertAffectedByAttackEventRow(excelFile, affectedByAttackEvent, affectedCharactersIndex)
			affectedCharactersIndex++
		}
		excelFile.SetActiveSheet(attackEventSheetIndex)
		insertAttackEventRow(excelFile, attackEvent, index1)
	}

	generateDiceEventHeaders(excelFile)
	excelFile.SetActiveSheet(diceEventSheetIndex)
	for i, diceEvent := range diceEvents{
		insertDiceEventRow(excelFile, diceEvent, i)
	}

	excelBytes, err := excelFile.WriteToBuffer()
	if err != nil { // Should change to return buffer
		return &bytes.Buffer{}, err
	}
	return excelBytes, nil
}

func generateCharacterTradeHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Trades", "A1", "character_trade_id")
	excelFile.SetCellValue("Character X Trades", "B1", "trade_event_id")
	excelFile.SetCellValue("Character X Trades", "C1", "weapon_x_character")
	excelFile.SetCellValue("Character X Trades", "D1", "item_x_character")
	excelFile.SetCellValue("Character X Trades", "E1", "armor_x_character")
	excelFile.SetCellValue("Character X Trades", "F1", "item_owner")
	excelFile.SetCellValue("Character X Trades", "G1", "item_reciever")
	excelFile.SetCellValue("Character X Trades", "H1", "quantity")
	excelFile.SetCellValue("Character X Trades", "I1", "item_name")
	excelFile.SetCellValue("Character X Trades", "J1", "item_type")
}

func insertCharacterTradeRow(excelFile *excelize.File, characterTrade domain.CharacterTrade, i int) {
	excelFile.SetCellValue("Character X Trades", "A"+strconv.Itoa(i+2), characterTrade.CharacterTrade_Id)
	excelFile.SetCellValue("Character X Trades", "B"+strconv.Itoa(i+2), characterTrade.TradeEvent_Id)
	if characterTrade.WeaponXCharacter != nil {
		excelFile.SetCellValue("Character X Trades", "C"+strconv.Itoa(i+2), *characterTrade.WeaponXCharacter)
	}
	if characterTrade.ItemXCharacter != nil {
		excelFile.SetCellValue("Character X Trades", "D"+strconv.Itoa(i+2), *characterTrade.ItemXCharacter)
	}
	if characterTrade.ArmorXCharacter != nil {
		excelFile.SetCellValue("Character X Trades", "E"+strconv.Itoa(i+2), *characterTrade.ArmorXCharacter)
	}
	excelFile.SetCellValue("Character X Trades", "F"+strconv.Itoa(i+2), characterTrade.ItemOwner)
	excelFile.SetCellValue("Character X Trades", "G"+strconv.Itoa(i+2), characterTrade.ItemReciever)
	if characterTrade.Quantity != nil {
		excelFile.SetCellValue("Character X Trades", "H"+strconv.Itoa(i+2), *characterTrade.Quantity)
	}
	excelFile.SetCellValue("Character X Trades", "I"+strconv.Itoa(i+2), characterTrade.ItemName)
	excelFile.SetCellValue("Character X Trades", "J"+strconv.Itoa(i+2), characterTrade.ItemType)
}

func generateTradeEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Trade Events", "A1", "trade_event_id")
	excelFile.SetCellValue("Trade Events", "B1", "session_id")
	excelFile.SetCellValue("Trade Events", "C1", "sender")
	excelFile.SetCellValue("Trade Events", "D1", "receiver")
	excelFile.SetCellValue("Trade Events", "E1", "description")
	excelFile.SetCellValue("Trade Events", "F1", "timestamp")
}

func insertTradeEventRow(excelFile *excelize.File, tradeEvent domain.TradeEvent, i int) {
	excelFile.SetCellValue("Trade Events", "A"+strconv.Itoa(i+2), tradeEvent.TradeEvent_Id)
	excelFile.SetCellValue("Trade Events", "B"+strconv.Itoa(i+2), tradeEvent.Session_Id)
	excelFile.SetCellValue("Trade Events", "C"+strconv.Itoa(i+2), tradeEvent.Sender)
	excelFile.SetCellValue("Trade Events", "D"+strconv.Itoa(i+2), tradeEvent.Receiver)
	excelFile.SetCellValue("Trade Events", "E"+strconv.Itoa(i+2), tradeEvent.Description)
	if tradeEvent.Timestamp != nil {
		excelFile.SetCellValue("Trade Events", "F"+strconv.Itoa(i+2), *tradeEvent.Timestamp)
	}
}

func generateAttackEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Attack Event", "A1", "attack_event_id")
	excelFile.SetCellValue("Attack Event", "B1", "type")
	excelFile.SetCellValue("Attack Event", "C1", "environment")
	excelFile.SetCellValue("Attack Event", "D1", "session_id")
	excelFile.SetCellValue("Attack Event", "E1", "event_protagonist_id")
	excelFile.SetCellValue("Attack Event", "F1", "user_id")
	excelFile.SetCellValue("Attack Event", "G1", "campaign_id")
	excelFile.SetCellValue("Attack Event", "H1", "image_url")
	excelFile.SetCellValue("Attack Event", "I1", "name")
	excelFile.SetCellValue("Attack Event", "J1", "race")
	excelFile.SetCellValue("Attack Event", "K1", "class")
	excelFile.SetCellValue("Attack Event", "L1", "level")
	excelFile.SetCellValue("Attack Event", "M1", "hit_points")
	excelFile.SetCellValue("Attack Event", "N1", "event_resolution")
	excelFile.SetCellValue("Attack Event", "O1", "weapon_id")
	excelFile.SetCellValue("Attack Event", "P1", "spell_id")
	excelFile.SetCellValue("Attack Event", "Q1", "dmg_type")
	excelFile.SetCellValue("Attack Event", "R1", "description")
	excelFile.SetCellValue("Attack Event", "S1", "timestamp")
}

func insertAttackEventRow(excelFile *excelize.File, attackEvent dto.ResponseEventDto, i int) {
	excelFile.SetCellValue("Attack Event", "A"+strconv.Itoa(i+2), attackEvent.AttackEventId)
	excelFile.SetCellValue("Attack Event", "B"+strconv.Itoa(i+2), attackEvent.Type)
	excelFile.SetCellValue("Attack Event", "C"+strconv.Itoa(i+2), attackEvent.Environment)
	excelFile.SetCellValue("Attack Event", "D"+strconv.Itoa(i+2), attackEvent.Session.SessionId)
	excelFile.SetCellValue("Attack Event", "E"+strconv.Itoa(i+2), attackEvent.EventProtagonist.CharacterId)
	if attackEvent.EventProtagonist.UserId != nil {
	excelFile.SetCellValue("Attack Event", "F"+strconv.Itoa(i+2), *attackEvent.EventProtagonist.UserId)
	}
	if attackEvent.EventProtagonist.CampaignID != nil {
	excelFile.SetCellValue("Attack Event", "G"+strconv.Itoa(i+2), *attackEvent.EventProtagonist.CampaignID)
	}
	excelFile.SetCellValue("Attack Event", "H"+strconv.Itoa(i+2), attackEvent.EventProtagonist.ImageUrl)
	excelFile.SetCellValue("Attack Event", "I"+strconv.Itoa(i+2), attackEvent.EventProtagonist.Name)
	excelFile.SetCellValue("Attack Event", "J"+strconv.Itoa(i+2), attackEvent.EventProtagonist.Race)
	excelFile.SetCellValue("Attack Event", "K"+strconv.Itoa(i+2), attackEvent.EventProtagonist.Class)
	excelFile.SetCellValue("Attack Event", "L"+strconv.Itoa(i+2), attackEvent.EventProtagonist.Level)
	excelFile.SetCellValue("Attack Event", "M"+strconv.Itoa(i+2), attackEvent.EventProtagonist.HitPoints)
	excelFile.SetCellValue("Attack Event", "N"+strconv.Itoa(i+2), attackEvent.EventResolution)
	if attackEvent.Weapon != nil {
		excelFile.SetCellValue("Attack Event", "O"+strconv.Itoa(i+2), *attackEvent.Weapon)
	}
	if attackEvent.Spell != nil {
		excelFile.SetCellValue("Attack Event", "P"+strconv.Itoa(i+2), *attackEvent.Spell)
	}
	if attackEvent.DmgType != nil {
		excelFile.SetCellValue("Attack Event", "Q"+strconv.Itoa(i+2), *attackEvent.DmgType)
	}
	if attackEvent.Description != nil {
		excelFile.SetCellValue("Attack Event", "R"+strconv.Itoa(i+2), *attackEvent.Description)
	}
	if attackEvent.TimeStamp != nil {
		excelFile.SetCellValue("Attack Event", "S"+strconv.Itoa(i+2), *attackEvent.TimeStamp)
	}
}

func generateAffectedByAttackEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Affected By Attack Event", "A1", "character_id")
	excelFile.SetCellValue("Affected By Attack Event", "B1", "user_id")
	excelFile.SetCellValue("Affected By Attack Event", "C1", "campaign_id")
	excelFile.SetCellValue("Affected By Attack Event", "D1", "image_url")
	excelFile.SetCellValue("Affected By Attack Event", "E1", "name")
	excelFile.SetCellValue("Affected By Attack Event", "F1", "race")
	excelFile.SetCellValue("Affected By Attack Event", "G1", "class")
	excelFile.SetCellValue("Affected By Attack Event", "H1", "level")
	excelFile.SetCellValue("Affected By Attack Event", "I1", "hit_points")
}

func insertAffectedByAttackEventRow(excelFile *excelize.File, affectedByAttackEvent dto.CharacterCardDto, i int) {
	excelFile.SetCellValue("Affected By Attack Event", "A"+strconv.Itoa(i+2), affectedByAttackEvent.CharacterId)	
	if affectedByAttackEvent.CampaignID != nil {
		excelFile.SetCellValue("Affected By Attack Event", "B"+strconv.Itoa(i+2), *affectedByAttackEvent.CampaignID)
	}
	if affectedByAttackEvent.UserId != nil {
		excelFile.SetCellValue("Affected By Attack Event", "C"+strconv.Itoa(i+2), *affectedByAttackEvent.UserId)
	}
	excelFile.SetCellValue("Affected By Attack Event", "D"+strconv.Itoa(i+2), affectedByAttackEvent.ImageUrl)
	excelFile.SetCellValue("Affected By Attack Event", "E"+strconv.Itoa(i+2), affectedByAttackEvent.Name)
	excelFile.SetCellValue("Affected By Attack Event", "F"+strconv.Itoa(i+2), affectedByAttackEvent.Race)
	excelFile.SetCellValue("Affected By Attack Event", "G"+strconv.Itoa(i+2), affectedByAttackEvent.Class)
	excelFile.SetCellValue("Affected By Attack Event", "H"+strconv.Itoa(i+2), affectedByAttackEvent.Level)
	excelFile.SetCellValue("Affected By Attack Event", "I"+strconv.Itoa(i+2), affectedByAttackEvent.HitPoints)
}

func generateDiceEventHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Dice Event", "A1", "dice_event_id")
	excelFile.SetCellValue("Dice Event", "B1", "stat")
	excelFile.SetCellValue("Dice Event", "C1", "difficulty")
	excelFile.SetCellValue("Dice Event", "D1", "dice_rolled")
	excelFile.SetCellValue("Dice Event", "E1", "dice_result")
	excelFile.SetCellValue("Dice Event", "F1", "event_protagonist")
	excelFile.SetCellValue("Dice Event", "G1", "description")
	excelFile.SetCellValue("Dice Event", "H1", "session_id")
	excelFile.SetCellValue("Dice Event", "I1", "time_stamp")
}

func insertDiceEventRow(excelFile *excelize.File, diceEvent domain.DiceEvent, i int) {
	excelFile.SetCellValue("Dice Event", "A"+strconv.Itoa(i+2), diceEvent.DiceEventId)
	excelFile.SetCellValue("Dice Event", "B"+strconv.Itoa(i+2), diceEvent.Stat)
	excelFile.SetCellValue("Dice Event", "C"+strconv.Itoa(i+2), diceEvent.Difficulty)
	excelFile.SetCellValue("Dice Event", "D"+strconv.Itoa(i+2), diceEvent.DiceRolled)
	excelFile.SetCellValue("Dice Event", "E"+strconv.Itoa(i+2), diceEvent.DiceResult)
	excelFile.SetCellValue("Dice Event", "F"+strconv.Itoa(i+2), diceEvent.EventProtagonist)
	excelFile.SetCellValue("Dice Event", "G"+strconv.Itoa(i+2), diceEvent.Description)
	excelFile.SetCellValue("Dice Event", "H"+strconv.Itoa(i+2), diceEvent.SessionId)
	excelFile.SetCellValue("Dice Event", "I"+strconv.Itoa(i+2), diceEvent.TimeStamp)
}

