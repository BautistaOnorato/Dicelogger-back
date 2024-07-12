package report

import (
	"bytes"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/proyecto-dnd/backend/internal/domain"
	"github.com/proyecto-dnd/backend/internal/dto"
)

func (r *ReportGenerator) GenerateCharactersByCampaignReport(id int) (*bytes.Buffer, error) {
	excelFile := excelize.NewFile()
	characterSheetIndex := excelFile.NewSheet("Sheet1")
	excelFile.SetSheetName("Sheet1", "Character Data")
	characterXItemSheetIndex := excelFile.NewSheet("Character X Items")
	characterXWeaponSheetIndex := excelFile.NewSheet("Character X Weapons")
	characterXArmorSheetIndex := excelFile.NewSheet("Character X Armor")
	skillSheetIndex := excelFile.NewSheet("Character X Skills")
	featureSheetIndex := excelFile.NewSheet("Character X Features")
	spellSheetIndex := excelFile.NewSheet("Character X Spells")
	proficiencySheetIndex := excelFile.NewSheet("Character X Proficiencies")

	miniCharactersDto, err := r.characterDataService.GetByCampaignId(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	itemIndex := 0
	weaponIndex := 0
	armorIndex := 0
	skillIndex := 0 
	featureIndex := 0
	spellIndex := 0
	proficiencyIndex := 0
	generateCharacterDataHeaders(excelFile)
	excelFile.SetActiveSheet(characterSheetIndex)
	generateItemXCharacterDataHeaders(excelFile)
	excelFile.SetActiveSheet(characterXItemSheetIndex)
	generateWeaponXCharacterDataHeaders(excelFile)
	excelFile.SetActiveSheet(characterXWeaponSheetIndex)
	generateArmorXCharacterDataHeaders(excelFile)
	excelFile.SetActiveSheet(characterXArmorSheetIndex)
	generateSkillsHeaders(excelFile)
	excelFile.SetActiveSheet(skillSheetIndex)
	generateFeaturesHeaders(excelFile)
	excelFile.SetActiveSheet(featureSheetIndex)
	generateSpellsHeaders(excelFile)
	excelFile.SetActiveSheet(spellSheetIndex)
	generateProficienciesHeaders(excelFile)
	excelFile.SetActiveSheet(proficiencySheetIndex)
	for index, character := range miniCharactersDto {
		fullCharacter, err := r.characterDataService.GetById(character.CharacterId)
		if err != nil {
			return &bytes.Buffer{}, err
		}
		insertCharacterDataRow(excelFile, &fullCharacter ,index)
		for _, item := range fullCharacter.Items {
			insertItemXCharacterDataRow(excelFile, item, itemIndex)
			excelFile.SetActiveSheet(characterXItemSheetIndex)
			itemIndex++
		}
		for _, weapon := range fullCharacter.Weapons {
			insertWeaponXCharacterDataRow(excelFile, weapon, weaponIndex)
			excelFile.SetActiveSheet(characterXWeaponSheetIndex)
			weaponIndex++
		}
		for _, armor := range fullCharacter.Armor {
			insertArmorXCharacterDataRow(excelFile, armor, armorIndex)
			excelFile.SetActiveSheet(characterXArmorSheetIndex)
			armorIndex++
		}
		for _, skill := range fullCharacter.Skills {
			insertSkillsRow(excelFile, skill, character.CharacterId, skillIndex)
			excelFile.SetActiveSheet(skillSheetIndex)
			skillIndex++
		}
		for _, feature := range fullCharacter.Features {
			insertFeaturesRow(excelFile, feature, character.CharacterId, featureIndex)
			excelFile.SetActiveSheet(featureSheetIndex)
			featureIndex++
		}
		for _, spell := range fullCharacter.Spells {
			insertSpellsRow(excelFile, spell, character.CharacterId, spellIndex)
			excelFile.SetActiveSheet(spellSheetIndex)
			spellIndex++
		}
		for _, proficiency := range fullCharacter.Proficiencies {
			insertProficienciesRow(excelFile, proficiency, character.CharacterId, proficiencyIndex)
			excelFile.SetActiveSheet(proficiencySheetIndex)
			proficiencyIndex++
		}
		excelFile.SetActiveSheet(characterSheetIndex)
	}


	excelBytes, err := excelFile.WriteToBuffer()
	if err != nil { // Should change to return buffer
		return &bytes.Buffer{}, err
	}
	return excelBytes, nil
}

func generateCharacterDataHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character Data", "A1", "character_id")
	excelFile.SetCellValue("Character Data", "B1", "user_id")
	excelFile.SetCellValue("Character Data", "C1", "campaign_id")
	excelFile.SetCellValue("Character Data", "D1", "race_id")
	excelFile.SetCellValue("Character Data", "E1", "name")
	excelFile.SetCellValue("Character Data", "F1", "description")
	excelFile.SetCellValue("Character Data", "G1", "speed")
	excelFile.SetCellValue("Character Data", "H1", "str")
	excelFile.SetCellValue("Character Data", "I1", "dex")
	excelFile.SetCellValue("Character Data", "J1", "int")
	excelFile.SetCellValue("Character Data", "K1", "con")
	excelFile.SetCellValue("Character Data", "L1", "wiz")
	excelFile.SetCellValue("Character Data", "M1", "cha")
	excelFile.SetCellValue("Character Data", "N1", "class_id")
	excelFile.SetCellValue("Character Data", "O1", "name")
	excelFile.SetCellValue("Character Data", "P1", "description")
	excelFile.SetCellValue("Character Data", "Q1", "proficiency_bonus")
	excelFile.SetCellValue("Character Data", "R1", "hit_dice")
	excelFile.SetCellValue("Character Data", "S1", "armor_proficiencies")
	excelFile.SetCellValue("Character Data", "T1", "weapon_proficiencies")
	excelFile.SetCellValue("Character Data", "U1", "tool_proficiencies")
	excelFile.SetCellValue("Character Data", "V1", "spellcasting_ability")
	excelFile.SetCellValue("Character Data", "W1", "background_id")
	excelFile.SetCellValue("Character Data", "X1", "name")
	excelFile.SetCellValue("Character Data", "Y1", "languages")
	excelFile.SetCellValue("Character Data", "Z1", "personality_traits")
	excelFile.SetCellValue("Character Data", "AA1", "ideals")
	excelFile.SetCellValue("Character Data", "AB1", "bond")
	excelFile.SetCellValue("Character Data", "AC1", "flaws")
	excelFile.SetCellValue("Character Data", "AD1", "trait")
	excelFile.SetCellValue("Character Data", "AE1", "tool_proficiencies")
	excelFile.SetCellValue("Character Data", "AF1", "name")
	excelFile.SetCellValue("Character Data", "AG1", "story")
	excelFile.SetCellValue("Character Data", "AH1", "alignment")
	excelFile.SetCellValue("Character Data", "AI1", "age")
	excelFile.SetCellValue("Character Data", "AJ1", "hair")
	excelFile.SetCellValue("Character Data", "AK1", "eyes")
	excelFile.SetCellValue("Character Data", "AL1", "skin")
	excelFile.SetCellValue("Character Data", "AM1", "height")
	excelFile.SetCellValue("Character Data", "AN1", "weight")
	excelFile.SetCellValue("Character Data", "AO1", "img")
	excelFile.SetCellValue("Character Data", "AP1", "str")
	excelFile.SetCellValue("Character Data", "AQ1", "dex")
	excelFile.SetCellValue("Character Data", "AR1", "int")
	excelFile.SetCellValue("Character Data", "AS1", "con")
	excelFile.SetCellValue("Character Data", "AT1", "wiz")
	excelFile.SetCellValue("Character Data", "AU1", "cha")
	excelFile.SetCellValue("Character Data", "AV1", "hitpoints")
	excelFile.SetCellValue("Character Data", "AW1", "hitdice")
	excelFile.SetCellValue("Character Data", "AX1", "speed")
	excelFile.SetCellValue("Character Data", "AY1", "armor_class")
	excelFile.SetCellValue("Character Data", "AZ1", "level")
	excelFile.SetCellValue("Character Data", "BA1", "exp")
}

func insertCharacterDataRow(excelFile *excelize.File, characterData *dto.FullCharacterData, index int) {
	excelFile.SetCellValue("Character Data", "A"+strconv.Itoa(index+2), characterData.Character_Id)
	if characterData.User_Id != nil {
		excelFile.SetCellValue("Character Data", "B"+strconv.Itoa(index+2), *characterData.User_Id)
	} 
	excelFile.SetCellValue("Character Data", "C"+strconv.Itoa(index+2), characterData.Campaign_Id)
	excelFile.SetCellValue("Character Data", "D"+strconv.Itoa(index+2), characterData.Race.RaceID)
	excelFile.SetCellValue("Character Data", "E"+strconv.Itoa(index+2), characterData.Race.Name)
	excelFile.SetCellValue("Character Data", "F"+strconv.Itoa(index+2), characterData.Race.Description)
	excelFile.SetCellValue("Character Data", "G"+strconv.Itoa(index+2), characterData.Race.Speed)
	excelFile.SetCellValue("Character Data", "H"+strconv.Itoa(index+2), characterData.Race.Str)
	excelFile.SetCellValue("Character Data", "I"+strconv.Itoa(index+2), characterData.Race.Dex)
	excelFile.SetCellValue("Character Data", "J"+strconv.Itoa(index+2), characterData.Race.Int)
	excelFile.SetCellValue("Character Data", "K"+strconv.Itoa(index+2), characterData.Race.Con)
	excelFile.SetCellValue("Character Data", "L"+strconv.Itoa(index+2), characterData.Race.Wiz)
	excelFile.SetCellValue("Character Data", "M"+strconv.Itoa(index+2), characterData.Race.Cha)
	excelFile.SetCellValue("Character Data", "N"+strconv.Itoa(index+2), characterData.Class.ClassId)
	excelFile.SetCellValue("Character Data", "O"+strconv.Itoa(index+2), characterData.Class.Name)
	excelFile.SetCellValue("Character Data", "P"+strconv.Itoa(index+2), characterData.Class.Description)
	excelFile.SetCellValue("Character Data", "Q"+strconv.Itoa(index+2), characterData.Class.ProficiencyBonus)
	excelFile.SetCellValue("Character Data", "R"+strconv.Itoa(index+2), characterData.Class.HitDice)
	excelFile.SetCellValue("Character Data", "S"+strconv.Itoa(index+2), characterData.Class.ArmorProficiencies)
	excelFile.SetCellValue("Character Data", "T"+strconv.Itoa(index+2), characterData.Class.WeaponProficiencies)
	excelFile.SetCellValue("Character Data", "U"+strconv.Itoa(index+2), characterData.Class.ToolProficiencies)
	excelFile.SetCellValue("Character Data", "V"+strconv.Itoa(index+2), characterData.Class.SpellcastingAbility)
	excelFile.SetCellValue("Character Data", "W"+strconv.Itoa(index+2), characterData.Background.BackgroundID)
	excelFile.SetCellValue("Character Data", "X"+strconv.Itoa(index+2), characterData.Background.Name)
	excelFile.SetCellValue("Character Data", "Y"+strconv.Itoa(index+2), characterData.Background.Languages)
	excelFile.SetCellValue("Character Data", "Z"+strconv.Itoa(index+2), characterData.Background.PersonalityTraits)
	excelFile.SetCellValue("Character Data", "AA"+strconv.Itoa(index+2), characterData.Background.Ideals)
	excelFile.SetCellValue("Character Data", "AB"+strconv.Itoa(index+2), characterData.Background.Bond)
	excelFile.SetCellValue("Character Data", "AC"+strconv.Itoa(index+2), characterData.Background.Flaws)
	excelFile.SetCellValue("Character Data", "AD"+strconv.Itoa(index+2), characterData.Background.Trait)
	excelFile.SetCellValue("Character Data", "AE"+strconv.Itoa(index+2), characterData.Background.ToolProficiencies)
	excelFile.SetCellValue("Character Data", "AF"+strconv.Itoa(index+2), characterData.Name)
	excelFile.SetCellValue("Character Data", "AG"+strconv.Itoa(index+2), characterData.Story)
	excelFile.SetCellValue("Character Data", "AH"+strconv.Itoa(index+2), characterData.Alignment)
	excelFile.SetCellValue("Character Data", "AI"+strconv.Itoa(index+2), characterData.Age)
	excelFile.SetCellValue("Character Data", "AJ"+strconv.Itoa(index+2), characterData.Hair)
	excelFile.SetCellValue("Character Data", "AK"+strconv.Itoa(index+2), characterData.Eyes)
	excelFile.SetCellValue("Character Data", "AL"+strconv.Itoa(index+2), characterData.Skin)
	excelFile.SetCellValue("Character Data", "AM"+strconv.Itoa(index+2), characterData.Height)
	excelFile.SetCellValue("Character Data", "AN"+strconv.Itoa(index+2), characterData.Weight)
	excelFile.SetCellValue("Character Data", "AO"+strconv.Itoa(index+2), characterData.ImgUrl)
	excelFile.SetCellValue("Character Data", "AP"+strconv.Itoa(index+2), characterData.Str)
	excelFile.SetCellValue("Character Data", "AQ"+strconv.Itoa(index+2), characterData.Dex)
	excelFile.SetCellValue("Character Data", "AR"+strconv.Itoa(index+2), characterData.Int)
	excelFile.SetCellValue("Character Data", "AS"+strconv.Itoa(index+2), characterData.Con)
	excelFile.SetCellValue("Character Data", "AT"+strconv.Itoa(index+2), characterData.Wiz)
	excelFile.SetCellValue("Character Data", "AU"+strconv.Itoa(index+2), characterData.Cha)
	excelFile.SetCellValue("Character Data", "AV"+strconv.Itoa(index+2), characterData.Hitpoints)
	excelFile.SetCellValue("Character Data", "AW"+strconv.Itoa(index+2), characterData.HitDice)
	excelFile.SetCellValue("Character Data", "AX"+strconv.Itoa(index+2), characterData.Speed)
	excelFile.SetCellValue("Character Data", "AY"+strconv.Itoa(index+2), characterData.Armor_Class)
	excelFile.SetCellValue("Character Data", "AZ"+strconv.Itoa(index+2), characterData.Level)
	excelFile.SetCellValue("Character Data", "BA"+strconv.Itoa(index+2), characterData.Exp)
}

func generateItemXCharacterDataHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Items", "A1", "character_item_id")
	excelFile.SetCellValue("Character X Items", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Items", "C1", "item_id")
	excelFile.SetCellValue("Character X Items", "D1", "name")
	excelFile.SetCellValue("Character X Items", "E1", "weight")
	excelFile.SetCellValue("Character X Items", "F1", "price")
	excelFile.SetCellValue("Character X Items", "G1", "description")
	excelFile.SetCellValue("Character X Items", "H1", "campaign_id")
	excelFile.SetCellValue("Character X Items", "I1", "quantity")
}

func insertItemXCharacterDataRow(excelFile *excelize.File, characterItemData domain.ItemXCharacterData, index int) {
	excelFile.SetCellValue("Character X Items", "A"+strconv.Itoa(index+2), characterItemData.Character_Item_Id)
	excelFile.SetCellValue("Character X Items", "B"+strconv.Itoa(index+2), characterItemData.CharacterData_Id)
	excelFile.SetCellValue("Character X Items", "C"+strconv.Itoa(index+2), characterItemData.Item.Item_Id)
	excelFile.SetCellValue("Character X Items", "D"+strconv.Itoa(index+2), characterItemData.Item.Name)
	excelFile.SetCellValue("Character X Items", "E"+strconv.Itoa(index+2), characterItemData.Item.Weight)
	excelFile.SetCellValue("Character X Items", "F"+strconv.Itoa(index+2), characterItemData.Item.Price)
	excelFile.SetCellValue("Character X Items", "G"+strconv.Itoa(index+2), characterItemData.Item.Description)
	if characterItemData.Item.Campaign_Id != nil{
		excelFile.SetCellValue("Character X Items", "H"+strconv.Itoa(index+2), *characterItemData.Item.Campaign_Id)
	}
	excelFile.SetCellValue("Character X Items", "I"+strconv.Itoa(index+2), characterItemData.Quantity)
}

func generateWeaponXCharacterDataHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Weapons", "A1", "character_weapon_id")
	excelFile.SetCellValue("Character X Weapons", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Weapons", "C1", "weapon_id")
	excelFile.SetCellValue("Character X Weapons", "D1", "weapon_type")
	excelFile.SetCellValue("Character X Weapons", "E1", "name")
	excelFile.SetCellValue("Character X Weapons", "F1", "weight")
	excelFile.SetCellValue("Character X Weapons", "G1", "price")
	excelFile.SetCellValue("Character X Weapons", "H1", "category")
	excelFile.SetCellValue("Character X Weapons", "I1", "reach")
	excelFile.SetCellValue("Character X Weapons", "J1", "description")
	excelFile.SetCellValue("Character X Weapons", "K1", "damage")
	excelFile.SetCellValue("Character X Weapons", "L1", "versatile_damage")
	excelFile.SetCellValue("Character X Weapons", "M1", "ammunition")
	excelFile.SetCellValue("Character X Weapons", "N1", "damage_type")
	excelFile.SetCellValue("Character X Weapons", "O1", "campaign_id")
	excelFile.SetCellValue("Character X Weapons", "P1", "equipped")
}

func insertWeaponXCharacterDataRow(excelFile *excelize.File, weaponXCharacterData domain.WeaponXCharacterData, index int) {
	excelFile.SetCellValue("Character X Weapons", "A"+strconv.Itoa(index+2), weaponXCharacterData.Character_Weapon_Id)
	excelFile.SetCellValue("Character X Weapons", "B"+strconv.Itoa(index+2), weaponXCharacterData.CharacterData_Id)
	excelFile.SetCellValue("Character X Weapons", "C"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Weapon_Id)
	excelFile.SetCellValue("Character X Weapons", "D"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Weapon_Type)
	excelFile.SetCellValue("Character X Weapons", "E"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Name)
	excelFile.SetCellValue("Character X Weapons", "F"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Weight)
	excelFile.SetCellValue("Character X Weapons", "G"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Price)
	excelFile.SetCellValue("Character X Weapons", "H"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Category)
	excelFile.SetCellValue("Character X Weapons", "I"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Reach)
	excelFile.SetCellValue("Character X Weapons", "J"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Description)
	excelFile.SetCellValue("Character X Weapons", "K"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Damage)
	excelFile.SetCellValue("Character X Weapons", "L"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Versatile_Damage)
	excelFile.SetCellValue("Character X Weapons", "M"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Ammunition)
	excelFile.SetCellValue("Character X Weapons", "N"+strconv.Itoa(index+2), weaponXCharacterData.Weapon.Damage_Type)
	if weaponXCharacterData.Weapon.Campaign_Id != nil{
		excelFile.SetCellValue("Character X Weapons", "O"+strconv.Itoa(index+2), *weaponXCharacterData.Weapon.Campaign_Id)		
	}
	excelFile.SetCellValue("Character X Weapons", "P"+strconv.Itoa(index+2), weaponXCharacterData.Equipped)
}

func generateArmorXCharacterDataHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Armor", "A1", "character_armor_id")
	excelFile.SetCellValue("Character X Armor", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Armor", "C1", "armor_id")
	excelFile.SetCellValue("Character X Armor", "D1", "material")
	excelFile.SetCellValue("Character X Armor", "E1", "name")
	excelFile.SetCellValue("Character X Armor", "F1", "weight")
	excelFile.SetCellValue("Character X Armor", "G1", "price")
	excelFile.SetCellValue("Character X Armor", "H1", "category")
	excelFile.SetCellValue("Character X Armor", "I1", "protection_type")
	excelFile.SetCellValue("Character X Armor", "J1", "description")
	excelFile.SetCellValue("Character X Armor", "K1", "penalty")
	excelFile.SetCellValue("Character X Armor", "L1", "strength")
	excelFile.SetCellValue("Character X Armor", "M1", "armor_class")
	excelFile.SetCellValue("Character X Armor", "N1", "dex_bonus")
	excelFile.SetCellValue("Character X Armor", "O1", "campaign_id")
	excelFile.SetCellValue("Character X Armor", "P1", "equipped")
}

func insertArmorXCharacterDataRow(excelFile *excelize.File, armorXCharacterData domain.ArmorXCharacterData, index int) {
	excelFile.SetCellValue("Character X Armor", "A"+strconv.Itoa(index+2), armorXCharacterData.ArmorXCharacterData_Id)
	excelFile.SetCellValue("Character X Armor", "B"+strconv.Itoa(index+2), armorXCharacterData.CharacterData_Id)
	excelFile.SetCellValue("Character X Armor", "C"+strconv.Itoa(index+2), armorXCharacterData.Armor.ArmorId)
	excelFile.SetCellValue("Character X Armor", "D"+strconv.Itoa(index+2), armorXCharacterData.Armor.Material)
	excelFile.SetCellValue("Character X Armor", "E"+strconv.Itoa(index+2), armorXCharacterData.Armor.Name)
	excelFile.SetCellValue("Character X Armor", "F"+strconv.Itoa(index+2), armorXCharacterData.Armor.Weight)
	excelFile.SetCellValue("Character X Armor", "G"+strconv.Itoa(index+2), armorXCharacterData.Armor.Price)
	excelFile.SetCellValue("Character X Armor", "H"+strconv.Itoa(index+2), armorXCharacterData.Armor.Category)
	excelFile.SetCellValue("Character X Armor", "I"+strconv.Itoa(index+2), armorXCharacterData.Armor.ProtectionType)
	excelFile.SetCellValue("Character X Armor", "J"+strconv.Itoa(index+2), armorXCharacterData.Armor.Description)
	excelFile.SetCellValue("Character X Armor", "K"+strconv.Itoa(index+2), armorXCharacterData.Armor.Penalty)
	excelFile.SetCellValue("Character X Armor", "L"+strconv.Itoa(index+2), armorXCharacterData.Armor.Strength)
	excelFile.SetCellValue("Character X Armor", "M"+strconv.Itoa(index+2), armorXCharacterData.Armor.ArmorClass)
	excelFile.SetCellValue("Character X Armor", "N"+strconv.Itoa(index+2), armorXCharacterData.Armor.DexBonus)
	if armorXCharacterData.Armor.CampaignId != nil{
		excelFile.SetCellValue("Character X Armor", "O"+strconv.Itoa(index+2), *armorXCharacterData.Armor.CampaignId)
		
	}
	excelFile.SetCellValue("Character X Armor", "P"+strconv.Itoa(index+2), armorXCharacterData.Equipped)
}

func generateSkillsHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Skills", "A1", "skill_id")
	excelFile.SetCellValue("Character X Skills", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Skills", "C1", "name")
	excelFile.SetCellValue("Character X Skills", "D1", "stat")
}

func insertSkillsRow(excelFile *excelize.File, skill domain.Skill, characterid, index int) {
	excelFile.SetCellValue("Character X Skills", "A"+strconv.Itoa(index+2), skill.SkillId)
	excelFile.SetCellValue("Character X Skills", "B"+strconv.Itoa(index+2), characterid)
	excelFile.SetCellValue("Character X Skills", "C"+strconv.Itoa(index+2), skill.Name)
	excelFile.SetCellValue("Character X Skills", "D"+strconv.Itoa(index+2), skill.Stat)
}

func generateFeaturesHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Features", "A1", "feature_id")
	excelFile.SetCellValue("Character X Features", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Features", "C1", "name")
	excelFile.SetCellValue("Character X Features", "D1", "description")
}

func insertFeaturesRow(excelFile *excelize.File, feature domain.Feature, characterid, index int) {
	excelFile.SetCellValue("Character X Features", "A"+strconv.Itoa(index+2), feature.FeatureId)
	excelFile.SetCellValue("Character X Features", "B"+strconv.Itoa(index+2), characterid)
	excelFile.SetCellValue("Character X Features", "C"+strconv.Itoa(index+2), feature.Name)
	excelFile.SetCellValue("Character X Features", "D"+strconv.Itoa(index+2), feature.Description)
}

func generateSpellsHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Spells", "A1", "spell_id")
	excelFile.SetCellValue("Character X Spells", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Spells", "C1", "name")
	excelFile.SetCellValue("Character X Spells", "D1", "description")
	excelFile.SetCellValue("Character X Spells", "E1", "range")
	excelFile.SetCellValue("Character X Spells", "F1", "ritual")
	excelFile.SetCellValue("Character X Spells", "G1", "duration")
	excelFile.SetCellValue("Character X Spells", "H1", "concentration")
	excelFile.SetCellValue("Character X Spells", "I1", "casting_time")
	excelFile.SetCellValue("Character X Spells", "J1", "level")
	excelFile.SetCellValue("Character X Spells", "K1", "damage_type")
	excelFile.SetCellValue("Character X Spells", "L1", "difficulty_class")
	excelFile.SetCellValue("Character X Spells", "M1", "aoe")
	excelFile.SetCellValue("Character X Spells", "N1", "school")
}

func insertSpellsRow(excelFile *excelize.File, spell domain.Spell, characterid, index int) {
	excelFile.SetCellValue("Character X Spells", "A"+strconv.Itoa(index+2), spell.SpellId)
	excelFile.SetCellValue("Character X Spells", "B"+strconv.Itoa(index+2), characterid)
	excelFile.SetCellValue("Character X Spells", "C"+strconv.Itoa(index+2), spell.Name)
	excelFile.SetCellValue("Character X Spells", "D"+strconv.Itoa(index+2), spell.Description)
	excelFile.SetCellValue("Character X Spells", "E"+strconv.Itoa(index+2), spell.Range)
	excelFile.SetCellValue("Character X Spells", "F"+strconv.Itoa(index+2), spell.Ritual)
	excelFile.SetCellValue("Character X Spells", "G"+strconv.Itoa(index+2), spell.Duration)
	excelFile.SetCellValue("Character X Spells", "H"+strconv.Itoa(index+2), spell.Concentration)
	excelFile.SetCellValue("Character X Spells", "I"+strconv.Itoa(index+2), spell.CastingTime)
	excelFile.SetCellValue("Character X Spells", "J"+strconv.Itoa(index+2), spell.Level)
	excelFile.SetCellValue("Character X Spells", "K"+strconv.Itoa(index+2), spell.DamageType)
	excelFile.SetCellValue("Character X Spells", "L"+strconv.Itoa(index+2), spell.DifficultyClass)
	excelFile.SetCellValue("Character X Spells", "M"+strconv.Itoa(index+2), spell.Aoe)
	excelFile.SetCellValue("Character X Spells", "N"+strconv.Itoa(index+2), spell.School)
}

func generateProficienciesHeaders(excelFile *excelize.File) {
	excelFile.SetCellValue("Character X Proficiencies", "A1", "proficiency_id")
	excelFile.SetCellValue("Character X Proficiencies", "B1", "character_data_id")
	excelFile.SetCellValue("Character X Proficiencies", "C1", "name")
	excelFile.SetCellValue("Character X Proficiencies", "D1", "type")
}

func insertProficienciesRow(excelFile *excelize.File, proficiency domain.Proficiency, characterid, index int) {
	excelFile.SetCellValue("Character X Proficiencies", "A"+strconv.Itoa(index+2), proficiency.ProficiencyId)
	excelFile.SetCellValue("Character X Proficiencies", "B"+strconv.Itoa(index+2), characterid)
	excelFile.SetCellValue("Character X Proficiencies", "C"+strconv.Itoa(index+2), proficiency.Name)
	excelFile.SetCellValue("Character X Proficiencies", "D"+strconv.Itoa(index+2), proficiency.Type)
}