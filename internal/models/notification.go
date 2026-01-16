package models

type Notification struct {
	Recipient Recipient
	Subject   string
	Body      string
}
