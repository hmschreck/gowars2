package api

import (
	"github.com/imroc/req"
)

const API_ENDPOINT = "https://api.guildwars2.com/v2/"

type GW2API struct {
	Key string
}

type GW2APIReturn interface {}

func NewGW2API(key string) (output *GW2API) {
	output = &GW2API{Key: key}
	return
}

func (api *GW2API) Get(endpoint string) (json_body *req.Resp, err error){
	header := req.Header{
		"Authorization": "Bearer " + api.Key,
	}
	json_body, err = req.Get(endpoint, header)
	if err != nil {
		return
	}
	return
}





