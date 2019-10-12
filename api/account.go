package api

import (
	"fmt"
	"time"
)

type Account struct {
	Id string `json:"id"`
	Age int `json:"age"`
	Name string `json:"name"`
	World int `json:"world"`
	Guild []Guild `json:"guilds"`
	GuildLeader []Guild `json:"guild_leader"`
	Created time.Time `json:"created"`
	Access []string `json:"access"`
	Commander bool `json:"commander"`
	FractalLevel int `json:"fractal_level"`
	DailyAP int `json:"daily_ap"`
	MonthlyAP int `json:"monthly_ap"`
	WVWRank int `json:"wvw_rank"`
	LastModified time.Time `json:"last_modified"`
}

func (api *GW2API) Account() (output Account, err error){
	json_body, err := api.Get(API_ENDPOINT + "account")
	if err != nil {
		return
	}
	err = json_body.ToJSON(&output)
	return
}

type Account_Recipes struct {
	Recipes []Recipe
}

func (api *GW2API) Account_Recipes() (output Account_Recipes, err error){
	output.Recipes = []Recipe{}
	json_body, err := api.Get(API_ENDPOINT + "account/recipes")
	if err != nil {
		return
	}
	fmt.Println(json_body.String())
	recipes := make([]int, 0)
	err = json_body.ToJSON(&recipes)
	for _, value := range recipes {
		recipe := Recipe{Id: value}
		output.Recipes = append(output.Recipes, recipe)
	}
	return
}

