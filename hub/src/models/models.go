package models

// Subscriber represents a subscriber to a topic
type Subscriber struct {
	ID          string
	CallbackURL string
	Topic       string
	Secret      string
}

type User struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}
