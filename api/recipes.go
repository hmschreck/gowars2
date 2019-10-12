package api

import (
	"fmt"
	"strconv"
)

type Recipe struct {
	Id int `json:"id"`
	Type string `json:"type"`
	Output_Item_Id int `json:"output_item_id"`
	Output_Item_Count int `json:"output_item_count"`
	Time_To_Craft int `json:"time_to_craft_ms"` // time to craft in seconds
	Disciplines []string `json:"disciplines"`
	Min_Rating int `json:"min_rating"`
	Flags []string `json:"flags"`
	Ingredients []Ingredient `json:"ingredients"`
	Chat_Link string `json:"chat_link"`
}

type Ingredient struct {
	Item_ID int `json:"item_id"`
	Count int `json:"count"`
}

type Recipes struct {
	Recipes []Recipe
}

func (api *GW2API) Recipes() (output Recipes, err error){
	json_body, err := api.Get(API_ENDPOINT + "recipes")
	if err != nil {
		return
	}
	err = json_body.ToJSON(&output)
	return
}

func (api *GW2API) Get_Recipe_From_ID(recipe_id int) (output Recipe, err error){
	json_body, err := api.Get(fmt.Sprintf("%s%s/%s", API_ENDPOINT, "recipes", strconv.Itoa(recipe_id)))
	if err != nil {
		return
	}
	err = json_body.ToJSON(&output)
	return
}
