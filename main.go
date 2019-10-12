package main

import (
	"fmt"
	"github.com/hmschreck/gowars2/api"
)

func main() {
	api_obj := api.NewGW2API("5EF20C55-3323-3C4D-84CC-3D783B8A88F4F55E85D3-1183-4104-9DBB-4DE81CEABBB0")
	account, _ :=api_obj.Account()
	fmt.Println(account)
	account_recipes, _ := api_obj.Account_Recipes()
	fmt.Println(account_recipes)
}