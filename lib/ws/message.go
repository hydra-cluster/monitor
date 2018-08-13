package ws

// Message is an abstraction to defines the data that will be transported by the socket.
type Message struct {
	// Defines message response status
	Status string `json:"status"`
	// Defines which clients should receive this message
	To string `json:"to"`
	// Defines which clients id sent this message
	From string `json:"from"`
	// Defines how this message data should be handled
	Action string `json:"action"`
	// Content defines data of the message
	Content interface{} `json:"content"`
}

// NewMessage returns a pointer to a message
func NewMessage(to, from, action, status string, content interface{}) *Message {
	return &Message{
		Action:  action,
		To:      to,
		From:    from,
		Status:  status,
		Content: content,
	}
}
