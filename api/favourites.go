package api

import "github.com/thrashzone13/tgtg-golang/client"

type FavouritesResponse struct {
	Items []Item `json:"items"`
}

type Item struct {
	ItemsAvailable int      `json:"items_available"`
	ItemData       ItemData `json:"item"`
}

type ItemData struct {
	ItemPrice ItemPrice `json:"item_price"`
}

type ItemPrice struct {
	Code       string `json:"code"`
	MinorUnits int    `json:"minor_units"`
	Decimals   int    `json:"decimals"`
}

func Favourites(client *client.Client, accessToken string) (FavouritesResponse, error) {
	var res FavouritesResponse
	err := client.Request("POST", "/item/v8/", nil, &res, &accessToken)
	return res, err
}
