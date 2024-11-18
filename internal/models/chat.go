package models

type Chat struct {
	ID           int    `json:"id"`
	ApplicationToken string `json:"application_token"`
	Number       int    `json:"number"`
}
