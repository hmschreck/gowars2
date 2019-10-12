package api

type Detail struct {
	Type string `json:"type"`
	Weight_Class string `json:"weight_class"`
	Defense int `json:"defense"`
	Infusion_Slots []struct{
		Flags []string `json:"flags"`
		Item_Id int `json:"item_id"`
	}
	Infix_Upgrade []struct{
		Id int `json:"id"`
		Attributes []struct {
			Attribute []string `json:"attribute"`
			Modifier int `json:"modifier"`
		}
		Buff struct{
			Skill_ID int `json:"skill_id"`
			Description string `json:"description"`
		}
	} `json:"infix_upgrade"`
	Suffix_Item_Id int `json:"suffix_item_id"`
	Secondary_Suffix_Item_Id string `json:"secondary_suffix_item_id"`
	Stat_Choices []int `json:"stat_choices"`
	Size int `json:"size"`
	No_Sell_Or_Sort bool `json:"no_sell_or_sort"`
	Description string `json:"description"`
	Duration_Ms int `json:"duration_ms"`
	Unlock_Type []string `json:"unlock_type"`
	Color_ID int `json:"color_id"`
	Recipe_ID int `json:"recipe_id"`
	Extra_Recipe_IDs []int `json:"extra_recipe_ids"`
	Apply_Count int `json:"apply_count"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Skins []int `json:"skins"`
	Vendor_IDs []int `json:"vendor_ids"`
	Minipet_ID int `json:"minipet_id"`
	Charges []int `json:"charges"`
	Bonuses []string `json:"bonuses"`
	Damage_Type string `json:"damage_type"`
	Min_Power int `json:"min_power"`
	Max_Power int `json:"max_power"`
}
type Item struct {
	Id           int      `json:"id"`
	Chat_Link    string   `json:"chat_link"`
	Name         string   `json:"name"`
	Icon         string   `json:"icon"`
	Description  string   `json:"description"`
	Type         string   `json:"type"`
	Rarity       string   `json:"rarity"`
	Level        int      `json:"level"`
	Vendor_Value int      `json:"vendor_value"`
	Default_Skin int      `json:"default_skin"`
	Flags        []string `json:"flags"`
	Game_Types   []string `json:"game_types"`
	Restrictions []string `json:"restrictions"`
	Details Detail `json:"details"`
}
