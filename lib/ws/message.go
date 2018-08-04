package ws

// Message is an abstraction to defines the data that will be transported by the socket.
type Message struct {
	// Defines which clients should receive this message
	To string
	// Defines which clients id sent this message
	From string
	// Defines how this message data should be handled
	Action string
	// Content defines data of the message
	Content interface{}
}

// NewMessage returns a pointer to a message
func NewMessage(to, from, action string, content interface{}) *Message {
	return &Message{
		Action:  action,
		To:      to,
		From:    from,
		Content: content,
	}
}
