package domain

type MessageResponse struct {
	Message string `json:"message"`
}

type Status string

const (
	StatusActive   Status = "ACTIVE"
	StatusDisabled Status = "DISABLED"
	StatusDeleted  Status = "DELETED"
)

var AllStatus = []Status{
	StatusActive,
	StatusDisabled,
	StatusDeleted,
}
