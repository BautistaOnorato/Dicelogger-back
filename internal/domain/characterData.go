package domain

type CharacterData struct {
	Character_Id int        `json:"character_id"`
	User_Id      *string    `json:"user_id"`
	Campaign_Id  int        `json:"campaign_id"`
	Race         Race       `json:"race"`
	Class        Class      `json:"class"`
	Background   Background `json:"background"`
	Name         string     `json:"name"`
	Story        string     `json:"story"`
	Alignment    string     `json:"alignment"`
	Age          int        `json:"age"`
	Hair         string     `json:"hair"`
	Eyes         string     `json:"eyes"`
	Skin         string     `json:"skin"`
	Height       int        `json:"height"`
	Weight       int        `json:"weight"`
	ImgUrl       string     `json:"img"`
	Str          int        `json:"str"`
	Dex          int        `json:"dex"`
	Int          int        `json:"int"`
	Con          int        `json:"con"`
	Wiz          int        `json:"wiz"`
	Cha          int        `json:"cha"`
	Hitpoints    int        `json:"hitpoints"`
	HitDice      string     `json:"hit_dice"`
	Speed        int        `json:"speed"`
	Armor_Class  int        `json:"armor_class"`
	Level        int        `json:"level"`
	Exp          int        `json:"exp"`
}
