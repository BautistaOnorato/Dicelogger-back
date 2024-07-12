package domain

type Feature struct {
	FeatureId   int    `json:"feature_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
