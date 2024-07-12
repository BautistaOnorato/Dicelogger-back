package domain

type Background struct {
	BackgroundID      int    `json:"background_id"`
	Name              string `json:"name"`
	Languages         string `json:"languages"`
	PersonalityTraits string `json:"personality_traits"`
	Ideals            string `json:"ideals"`
	Bond              string `json:"bond"`
	Flaws             string `json:"flaws"`
	Trait             string `json:"trait"`
	ToolProficiencies string `json:"tool_proficiencies"`
}
