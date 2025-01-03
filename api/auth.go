package api

import "github.com/thrashzone13/tgtg-golang/client"

type AuthByEmailReq struct {
	Email      string `json:"email"`
	DeviceType string `json:"device_type"`
}

type AuthByEmailRes struct {
	State     string `json:"state"`
	PollingId string `json:"polling_id"`
}

func AuthByEmail(client *client.Client, req AuthByEmailReq) (AuthByEmailRes, error) {
	var res AuthByEmailRes
	err := client.Request("POST", "/auth/v4/authByEmail", req, &res, nil)
	return res, err
}
