package controllers

import (
	"context"
	"encoding/json"
	"log"
	"strings"
	"sync"

	"te-redis-service/initalizers"
	"te-redis-service/models"

	"github.com/redis/go-redis/v9"
)

// Subscriber handles message subscription with filtering
type Subscriber struct {
	client   *redis.Client
	handlers map[string]MessageHandler
	filters  map[string]models.MessageFilter
	mu       sync.RWMutex
}

// MessageHandler represents a function that handles messages for a specific channel
type MessageHandler func(message models.Message) error

// NewSubscriber creates a new Subscriber instance
func NewSubscriber() (*Subscriber, error) {

	return &Subscriber{
		client:   initalizers.Redis,
		handlers: make(map[string]MessageHandler),
		filters:  make(map[string]models.MessageFilter),
	}, nil
}

// RegisterHandler adds a message handler for a specific channel with optional filtering
func (s *Subscriber) RegisterHandler(channel string, handler MessageHandler, filter models.MessageFilter) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[channel] = handler
	s.filters[channel] = filter
}

// messagePassesFilter checks if a message passes the specified filter
func messagePassesFilter(message models.Message, filter models.MessageFilter) bool {
	// Check message type
	if len(filter.Types) > 0 {
		typeMatch := false
		for _, t := range filter.Types {
			if message.Type == t {
				typeMatch = true
				break
			}
		}
		if !typeMatch {
			return false
		}
	}

	// Check priority
	if message.Priority < filter.MinPriority {
		return false
	}

	// Check metadata
	for key, value := range filter.MetadataContains {
		// if msgValue, exists := message.Metadata[key]; !exists || !strings.Contains(msgValue, value) {
		// 	return false
		// }
		if msgValue, exists := message.Metadata[key]; !exists {
			return false
		} else {
			if str, ok := value.(string); ok {
				if msg, msgstr := msgValue.(string); msgstr {
					if !strings.Contains(msg, str) {
						return false
					} else {
						return true
					}
				}
			}
		}
	}

	return true
}

// Subscribe listens for messages on multiple channels with filtering
func (s *Subscriber) Subscribe(ctx context.Context, channels ...string) error {
	pubsub := s.client.Subscribe(ctx, channels...)
	defer pubsub.Close()

	ch := pubsub.Channel()

	for {
		select {
		case msg := <-ch:
			if msg == nil {
				return nil
			}

			var message models.Message
			if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
				log.Printf("Failed to unmarshal message: %v", err)
				continue
			}

			s.mu.RLock()
			handler, hasHandler := s.handlers[msg.Channel]
			filter, hasFilter := s.filters[msg.Channel]
			s.mu.RUnlock()

			if hasHandler && (!hasFilter || messagePassesFilter(message, filter)) {
				if err := handler(message); err != nil {
					log.Printf("Handler error for channel %s: %v", msg.Channel, err)
				}
			}

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
