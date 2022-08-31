package models


type EventActionNotifyRequest struct {
	Subject    string
	Model     string
	Problem   string
	Link      string
	GmtCreate string
}

type EventActionNotifyResponse struct {
	Success bool
}
