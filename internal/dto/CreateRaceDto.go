package dto

type CreateRaceDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Speed       int    `json:"speed"`
	Str         int    `json:"str"`
	Dex         int    `json:"dex"`
	Int        int    `json:"inte"`
	Con         int    `json:"con"`
	Wiz         int    `json:"wiz"`
	Cha         int    `json:"cha"`
}
