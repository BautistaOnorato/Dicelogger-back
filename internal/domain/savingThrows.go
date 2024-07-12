package domain

type SavingThrow struct {
	SavingThrowId int  `json:"saving_thow_id"`
	ClassId       int  `json:"class_id"`
	Str           bool `json:"str"`
	Dex           bool `json:"dex"`
	Int           bool `json:"int"`
	Con           bool `json:"con"`
	Wiz           bool `json:"wiz"`
	Cha           bool `json:"cha"`
}
