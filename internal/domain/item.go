package domain

type Item struct{
	Item_Id int `json:"item_id"`
	Name string `json:"name"`
    Weight int `json:"weight"`
	Price int `json:"price"`
	Description string `json:"description"`
	Campaign_Id *int `json:"campaign_id"`
}