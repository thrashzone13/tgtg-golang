package api

import "github.com/thrashzone13/tgtg-golang/client"

type AuthByPollingIdReq struct {
	Email            string `json:"email"`
	DeviceType       string `json:"device_type"`
	RequestPollingId string `json:"request_polling_id"`
}

type AuthByPollingIdRes struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	AccessTokenTtlSeconds int    `json:"access_token_ttl_seconds"`
}

func AuthByPollingId(client *client.Client, req AuthByPollingIdReq) (AuthByPollingIdRes, error) {
	var res AuthByPollingIdRes
	err := client.Request("POST", "/auth/v4/authByRequestPollingId", req, &res, nil)
	return res, err
}
