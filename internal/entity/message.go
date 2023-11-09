package entity

import "time"

type Message struct {
	Id        string
	Username  string
	Recipient string
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Message   string
}

type RequestFetchMessage struct {
	Username  string
	Recipient string
	StartDate string
	EndDate   string
}
