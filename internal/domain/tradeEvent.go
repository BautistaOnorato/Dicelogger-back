package domain

import "time"

type TradeEvent struct {
	TradeEvent_Id int      `json:"trade_event_id"`
	Session_Id    int      `json:"session_id"`
	Sender        int      `json:"sender"`
	Receiver      int      `json:"receiver"`
	Description   string   `json:"description"`
	Timestamp     *time.Time `json:"timestamp"`
	TradingItems []CharacterTrade `json:"trading_items"`
}