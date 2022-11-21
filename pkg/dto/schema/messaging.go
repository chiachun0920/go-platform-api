package schema

type MessagingRequest struct {
	Text string `json:"text"`
	To   string `json:"to"`
}
