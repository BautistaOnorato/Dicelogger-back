package dto

type SavingThrowDto struct {
	ClassId int  `json:"class_id"`
	Str     bool `json:"str"`
	Dex     bool `json:"dex"`
	Int     bool `json:"int"`
	Con     bool `json:"con"`
	Wiz     bool `json:"wiz"`
	Cha     bool `json:"cha"`
}
