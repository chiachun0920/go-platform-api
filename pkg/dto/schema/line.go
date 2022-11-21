package schema

type Message struct {
	Type string `json:"type"`
	Id   string `json:"id"`
	Text string `json:"text"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

type MessageEvent struct {
	Type        string  `json:"type"`
	MessageData Message `json:"message"`
	SourceData  Source  `json:"source"`
	Timestamp   int64   `json:"timestamp"`
}

type LineWebhookRequest struct {
	Destination string         `json:"destination"`
	Events      []MessageEvent `json:"events"`
}
