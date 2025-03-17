package model

type Message struct {
	Id             int    `json:"id"`
	ConversationId string `json:"conversationId"`
	Type           string `json:"type"`
	Content        string `json:"content"`
	IsEnd          bool   `json:"is_end"`
	Timestamp      string `json:"timestamp"`
}
