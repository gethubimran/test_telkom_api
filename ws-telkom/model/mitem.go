package model

type Item struct{
	Id int `json:"id"`
	Urgency int `json:"urgency"`
	Due_interval int `json:"due_interval"`
	Due_unit string `json:"due_unit"`
	Description string `json:"description"`
}