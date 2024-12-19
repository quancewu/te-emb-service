package models

import "time"

// MessageType represents different types of messages
type MessageType string

const (
	TypeNotification MessageType = "notification"
	TypeAlert        MessageType = "alert"
	TypeMetric       MessageType = "metric"
	TypeData         MessageType = "data"
)

// Message represents the structure of messages we'll publish
type Message struct {
	ID        string                 `json:"id"`
	Type      MessageType            `json:"type"`
	Channel   string                 `json:"channel"`
	Content   string                 `json:"content"`
	Priority  int                    `json:"priority"`
	Metadata  map[string]interface{} `json:"metadata"`
	Timestamp time.Time              `json:"timestamp"`
}

// MessageFilter defines criteria for filtering messages
type MessageFilter struct {
	Types            []MessageType
	MinPriority      int
	Channels         []string
	MetadataContains map[string]interface{}
}
