package models

type EventActionNotifyRequest struct {
	Subject    string
	Model      string
	ActorAlias string
	Problem    string
	Link       string
	GmtCreate  string
}

type EventActionNotifyResponse struct {
	Success bool
}
